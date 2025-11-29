package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func (s *server) registerUser(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var b Body
	if err := s.DecodeBody(r, &b); err != nil {
		s.JSON(w, map[string]string{"error": "invalid body"}, 401)
		return
	}

	namespace := fmt.Sprintf("user-%s", random_string_from_charset(6))

	_, err := s.kubernetes_create_namespace(namespace)
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

	row := s.db.QueryRow(`
        SELECT password, namespace_id 
        FROM users 
        WHERE email=$1
    `, b.Email)

	var hashedPass string
	var namespaceID string

	err := row.Scan(&hashedPass, &namespaceID)
	if err != nil {
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
		s.JSON(w, map[string]string{"error": "token generation failed"}, 500)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_id",
		Value:    signed,
		HttpOnly: true,
		Secure:   false, 
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	s.JSON(w, map[string]string{
		"status":       "ok",
		"namespace_id": namespaceID,
	}, 200)
}
