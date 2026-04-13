package auth

import (
	"context"
	"net/http"
	"strings"
	"taskflow/internal/pkg/response"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				// http.Error(w, "unauthorized", http.StatusUnauthorized)
				response.Error(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				// http.Error(w, "invalid auth header", http.StatusUnauthorized)
				response.Error(w, http.StatusUnauthorized, "invalid auth header")
				return
			}

			tokenStr := parts[1]

			userID, err := ParseJWT(tokenStr, secret)
			if err != nil {
				// http.Error(w, "invalid token", http.StatusUnauthorized)
				response.Error(w, http.StatusUnauthorized, "invalid token")
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserID(ctx context.Context) string {
	id, _ := ctx.Value(UserIDKey).(string)
	return id
}
