package middleware

import (
	"context"
	"net/http"
	"strings"
	"todo/utils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserCtxKey = contextKey("user")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			next.ServeHTTP(w, r)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := utils.ValidateToken(tokenStr)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		if token != nil && token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			userID := uint(userIDFloat)

			ctx := context.WithValue(r.Context(), UserCtxKey, userID)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}