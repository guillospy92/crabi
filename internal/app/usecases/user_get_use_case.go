package usecases

import (
	"context"

	"github.com/guillospy92/crabi/internal/core/domain"
	"github.com/guillospy92/crabi/internal/core/ports"
)

// UserGetInfoUseCaseInterface interface that contains the method to get user
type UserGetInfoUseCaseInterface interface {
	GetUserInfo(ctx context.Context, email string) (*domain.UserEntity, error)
}

// UserGetUseCaseInfoCase implementation of the user info interface
type UserGetUseCaseInfoCase struct {
	userRepository ports.UserRepositoryInterface
}

// GetUserInfo use case handler get user info
func (u UserGetUseCaseInfoCase) GetUserInfo(ctx context.Context, email string) (*domain.UserEntity, error) {
	return u.userRepository.FindByEmail(ctx, email)
}

// NewUserGetUseCaseInfoCase new instance of UserGetUseCaseInfoCase
func NewUserGetUseCaseInfoCase(userRepository ports.UserRepositoryInterface) *UserGetUseCaseInfoCase {
	return &UserGetUseCaseInfoCase{userRepository: userRepository}
}
