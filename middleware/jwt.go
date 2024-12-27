package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	
)

var jwtSecret = []byte("your-secret-key")
	func JWTMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract the token from the Authorization header
			authHeader := r.Header.Get("Authorization")
		
			// Check if the Authorization header is present and has the correct format
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
				// remove the "Bearer " prefix from the token
				tokenString := strings.TrimPrefix(authHeader, "Bearer ")
				// Parse the token (verify token)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Invalid token signing method")
				}
				return jwtSecret, nil
			})
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			if !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			claims,ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}
			// setting the userId in the user context ; 

			userID :=uint(claims["userID"].(float64))
			ctx := context.WithValue(r.Context(), "userID", userID)
			
			

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
