package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type AuthClaims struct {
	NamespaceId string `json:"namespace_id"`
	jwt.RegisteredClaims
}

func (a AuthClaims) Validate() error {
	if a.NamespaceId == "" {
		return errors.New("auth_id invalid")
	}

	return nil
}

const AUTH_COOKIE_NAME = "auth_id"

var hmacKey = []byte("a-very-long-random-secret-key-at-least-32-bytes")

func keyFunction(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return hmacKey, nil
}

func (s *server) MiddlewareExtractCookie(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(AUTH_COOKIE_NAME)
		if err == http.ErrNoCookie {
			s.JSON(w, map[string]string{"error": "auth_id invalid", "message": "you have to log back in"}, 401)
			return
		}
		if err != nil {
			s.JSON(w, map[string]string{"error": "auth_id invalid", "message": "you have to log back in"}, 401)
			return
		}

		var claims AuthClaims

		token, err := jwt.ParseWithClaims(cookie.Value, &claims, keyFunction)

		if err != nil || !token.Valid {
			s.JSON(w, map[string]string{"error": "auth_id invalid", "message": "you have to log back in"}, 401)
			return
		}

		s.LogMsg(fmt.Sprintf("successfully decoded token: %v", claims.NamespaceId))

		ctx := context.WithValue(r.Context(), "namespace_id", claims.NamespaceId)
		r = r.WithContext(ctx)

		next(w, r)
	}
}

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlewareLoggin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		g := color.New(color.FgRed, color.Bold)
		whiteBackground := g.Add(color.BgWhite)
		whiteBackground.Printf("%s   %s   %s   at   %s\n", r.Method, r.URL.Path, r.Host, time.Now().Format("2006-01-02 15:04:05"))
		next.ServeHTTP(w, r)
	})
}
