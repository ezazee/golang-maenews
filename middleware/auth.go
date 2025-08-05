package middleware

import (
	"context"
	"maenews/backend/auth"
	"net/http"
	"strings"
)

// JWTMiddleware adalah middleware untuk memverifikasi token JWT
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "email", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
