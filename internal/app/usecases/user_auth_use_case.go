package usecases

import (
	"context"
	"net/http"
	"time"

	"github.com/guillospy92/crabi/internal/core/domain"
	"github.com/guillospy92/crabi/internal/core/errors"
	"github.com/guillospy92/crabi/internal/core/ports"
	"github.com/guillospy92/crabi/pkg/bcrypt"
	"github.com/guillospy92/crabi/pkg/jwt"
	"github.com/guillospy92/crabi/resources"
)

//go:generate mockery --name=UserAuthUseCaseInterface --structname=UserAuthUseCaseInterfaceMock --filename=user_auth_use_case_mock_mock.go --output=mocks --outpkg=usecasesmock

// UserAuthUseCaseInterface interface that contains the method to log in a user
type UserAuthUseCaseInterface interface {
	Login(ctx context.Context, email, password string) (*domain.UserAuth, error)
}

// UserAuthUseCase implementation of the login  user interface
type UserAuthUseCase struct {
	userRepository ports.UserRepositoryInterface
}

// Login use case handler log in user
func (u UserAuthUseCase) Login(ctx context.Context, email, password string) (*domain.UserAuth, error) {
	userExists, err := u.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if userExists.Email == "" {
		return nil, errorlogic.GenerateErrorSinceMessage(
			errorlogic.ErrorCode(errorlogic.LoginErrorNotFound),
			errorlogic.StatusCode(http.StatusUnauthorized),
			errorlogic.Message("user email not exists"),
		)
	}

	checkPassword := pkgbcrypt.CheckPasswordHash(password, userExists.Password)
	if !checkPassword {
		return nil, errorlogic.GenerateErrorSinceMessage(
			errorlogic.ErrorCode(errorlogic.LoginErrorNotFound),
			errorlogic.StatusCode(http.StatusUnauthorized),
			errorlogic.Message("user password incorrect"),
		)
	}

	expirationDate := time.Now().Add(time.Duration(resources.ConfigurationEnv().JWTExpiredTokenTimeMinute) * time.Minute)
	token, err := pkgjwt.GenerateJWTToken(userExists, expirationDate)

	if err != nil {
		return nil, err
	}

	return &domain.UserAuth{
		User: userExists,
		Token: domain.Token{
			AccessToken: token,
			ExpiresIn:   &expirationDate,
		},
	}, nil
}

// NewUserAuthUseCase new instance of UserAuthUseCase
func NewUserAuthUseCase(userRepository ports.UserRepositoryInterface) *UserAuthUseCase {
	return &UserAuthUseCase{userRepository: userRepository}
}
