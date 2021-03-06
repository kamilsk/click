package middleware

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.octolab.org/ecosystem/click/internal/domain"
)

const (
	// AuthHeader defines authorization header.
	AuthHeader = "authorization"
	// AuthScheme defines authorization scheme.
	AuthScheme = "bearer"
)

// TokenInjector TODO issue#131
func TokenInjector(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, AuthScheme)
	if err != nil {
		return nil, err
	}
	tokenID := domain.ID(token)
	if !tokenID.IsValid() {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %s", token)
	}
	return context.WithValue(ctx, tokenKey{}, tokenID), nil
}

// TokenExtractor TODO issue#131
func TokenExtractor(ctx context.Context) (domain.ID, error) {
	tokenID, found := ctx.Value(tokenKey{}).(domain.ID)
	if !found {
		return tokenID, status.Error(codes.Unauthenticated, "auth token not found")
	}
	return tokenID, nil
}
