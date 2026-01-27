package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api/auth"
	"github.com/frostnzx/go-ecommerce-api/internal/core/utils"
)

func GetAuthMiddlewareFunc(tokenMaker *utils.JWTMaker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the authorization header
			// verify the token
			claims, err := verifyClaimsFromAuthHeader(r, tokenMaker)
			if err != nil {
				http.Error(w, fmt.Sprintf("error verifying token: %v", err), http.StatusUnauthorized)
				return
			}

			// pass the payload/claims down the context
			ctx := auth.SetClaimsInContext(r.Context(), claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetAdminMiddlewareFunc(tokenMaker *utils.JWTMaker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the authorization header
			// verify the token
			claims, err := verifyClaimsFromAuthHeader(r, tokenMaker)
			if err != nil {
				http.Error(w, fmt.Sprintf("error verifying token: %v", err), http.StatusUnauthorized)
				return
			}

			if !claims.IsAdmin {
				http.Error(w, "user is not an admin", http.StatusForbidden)
				return
			}

			// pass the payload/claims down the context
			ctx := auth.SetClaimsInContext(r.Context(), claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func verifyClaimsFromAuthHeader(r *http.Request, tokenMaker *utils.JWTMaker) (*utils.UserClaims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	fields := strings.Fields(authHeader)
	if len(fields) != 2 || fields[0] != "Bearer" {
		return nil, fmt.Errorf("invalid authorization header")
	}

	tokenStr := fields[1]
	claims, err := tokenMaker.VerifyToken(tokenStr)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	return claims, nil
}
