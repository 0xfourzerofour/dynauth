package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *) error
}