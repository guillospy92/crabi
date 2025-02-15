package ports

import (
	"context"

	"github.com/guillospy92/crabi/internal/core/domain"
)

// UserRepositoryInterface operate with user and methods persistent
type UserRepositoryInterface interface {
	Save(ctx context.Context, user domain.UserEntity) error
	FindByEmail(ctx context.Context, email string) (*domain.UserEntity, error)
}
