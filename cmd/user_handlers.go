package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"regexp"
	"time"
)

func (s *server) validateRegisterBody(email string, password string) error {
	if email == "" {
		return errors.New("email could not be empty")
	}
	if password == "" {
		return errors.New("password could not be empty")
	}

	// Validate email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
}

func (s *server) registerUser(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var b Body
	if err := s.DecodeBody(r, &b); err != nil {
		s.LogError("registerUser", err)
		s.JSON(w, map[string]string{"error": "invalid body"}, 401)
		return
	}

	err := s.validateRegisterBody(b.Email, b.Password)

	if err != nil {
		s.LogError("registerUser", err)
		s.JSON(w, map[string]string{"error": "validation failed", "message": err.Error()}, 500)
		return
	}

	namespace := fmt.Sprintf("user-%s", random_string_from_charset(6))

	_, err = s.kubernetes_create_namespace(namespace)
	if err != nil {
		s.LogError("registerUser", err)
		s.JSON(w, map[string]string{"error": "k8s namespace failed"}, 500)
		return
	}

	ingress, err := kubernetes_new_ingress(s.kconfig, namespace)
	if err != nil {
		s.LogError("registerUser", err)
		_ = s.kubernetes_delete_namespace(namespace)
		s.JSON(w, map[string]string{"error": "k8s ingress failed"}, 500)
		return
	}

	tx, err := s.db.Begin()
	if err != nil {
		s.LogError("registerUser", err)
		_ = s.kubernetes_delete_ingress(namespace, ingress.Name)
		_ = s.kubernetes_delete_namespace(namespace)
		s.JSON(w, map[string]string{"error": "db transaction failed"}, 500)
		return
	}

	_, err = tx.Exec(`INSERT INTO users (email, password, namespace_id)
                      VALUES ($1, $2, $3)`, b.Email, b.Password, namespace)
	if err != nil {
		s.LogError("registerUser", err)
		tx.Rollback()
		_ = s.kubernetes_delete_ingress(namespace, ingress.Name)
		_ = s.kubernetes_delete_namespace(namespace)
		s.JSON(w, map[string]string{"error": "db insert failed"}, 500)
		return
	}

	if err := tx.Commit(); err != nil {
		s.LogError("registerUser", err)
		_ = s.kubernetes_delete_ingress(namespace, ingress.Name)
		_ = s.kubernetes_delete_namespace(namespace)
		s.JSON(w, map[string]string{"error": "db commit failed"}, 500)
		return
	}

	s.JSON(w, map[string]string{"status": "ok", "namespace_id": namespace}, 200)
}

func (s *server) LogUser(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var b Body
	if err := s.DecodeBody(r, &b); err != nil {
		s.JSON(w, map[string]string{"error": "invalid body"}, 401)
		return
	}

	err := s.validateRegisterBody(b.Email, b.Password)

	if err != nil {
		s.LogError("registerUser", err)
		s.JSON(w, map[string]string{"error": "validation failed", "message": err.Error()}, 400)
	}

	row := s.db.QueryRow(`
        SELECT password, namespace_id 
        FROM users 
        WHERE email=$1
    `, b.Email)

	fmt.Println(b.Email)
	fmt.Println(b.Password)

	var hashedPass string
	var namespaceID string

	err = row.Scan(&hashedPass, &namespaceID)
	if err != nil {
		s.LogError("logUser", err)
		s.JSON(w, map[string]string{"error": "invalid credentials"}, 401)
		return
	}

	if hashedPass != b.Password {
		s.JSON(w, map[string]string{"error": "invalid credentials"}, 401)
		return
	}

	claims := AuthClaims{
		NamespaceId: namespaceID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "your-app",
			Subject:   b.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(hmacKey)
	if err != nil {
		s.LogError("logUser", err)
		s.JSON(w, map[string]string{"error": "token generation failed"}, 500)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_id",
		Value:    signed,
		HttpOnly: false,
		Secure:   false,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	s.JSON(w, map[string]string{
		"status":       "ok",
		"namespace_id": namespaceID,
	}, 200)
}

func (s *server) checkAuth(w http.ResponseWriter, r *http.Request) {

	namespace_name, ok := r.Context().Value("namespace_id").(string)

	if !ok {
		s.JSON(w, map[string]string{"error": "forbidden"}, 401)
		return
	}

	if namespace_name == "" {
		s.JSON(w, map[string]string{"error": "forbidden"}, 401)
		return
	}

	s.JSON(w, map[string]string{
		"status":       "ok",
		"namespace_id": namespace_name,
	}, 200)
}

func (s *server) handleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
	})

	s.JSON(w, map[string]string{
		"message": "Successfully logged out",
	}, 200)
}
