package auth

import (
	"context"
)

type AuthRepository interface {
	GetUser(ctx context.Context) error
}