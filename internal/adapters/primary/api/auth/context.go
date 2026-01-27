package auth

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/core/utils"
)

type authKey struct{}

// GetClaimsFromContext retrieves UserClaims from the request context
func GetClaimsFromContext(ctx context.Context) (*utils.UserClaims, bool) {
	claims, ok := ctx.Value(authKey{}).(*utils.UserClaims)
	return claims, ok
}

// SetClaimsInContext stores UserClaims in the request context
func SetClaimsInContext(ctx context.Context, claims *utils.UserClaims) context.Context {
	return context.WithValue(ctx, authKey{}, claims)
}
