package usecases

import (
	"context"
	"net/http"

	"github.com/guillospy92/crabi/internal/core/domain"
	"github.com/guillospy92/crabi/internal/core/errors"
	"github.com/guillospy92/crabi/internal/core/ports"
	"github.com/guillospy92/crabi/pkg/bcrypt"
)

//go:generate mockery --name=UserCreateUseCaseInterface --structname=UserCreateUseCaseInterfaceMock --filename=user_Create_use_case_mock.go --output=mocks --outpkg=usecasesmock

// UserCreateUseCaseInterface interface that contains the method to create a user
type UserCreateUseCaseInterface interface {
	CreateUser(ctx context.Context, user domain.UserEntity) error
}

// UserCreateUseCase implementation of the creation a user interface
type UserCreateUseCase struct {
	userBlackList  ports.UserBlackListerInterface
	userRepository ports.UserRepositoryInterface
}

// CreateUser use case handler create user
func (u *UserCreateUseCase) CreateUser(ctx context.Context, user domain.UserEntity) error {
	userExists, err := u.userRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if userExists.Email != "" {
		return errorlogic.GenerateErrorSinceMessage(
			errorlogic.ErrorCode(errorlogic.UserErrorExists),
			errorlogic.StatusCode(http.StatusConflict),
			errorlogic.Message("user exists"),
		)
	}

	verifyClient, err := u.userBlackList.VerifyUserBlackList(user)
	if err != nil {
		return err
	}

	if verifyClient {
		return errorlogic.GenerateErrorSinceMessage(
			errorlogic.ErrorCode(errorlogic.UserErrorBlocked),
			errorlogic.StatusCode(http.StatusForbidden),
			errorlogic.Message("user is blocked"),
		)
	}

	user.Password, err = pkgbcrypt.HashPassword(user.Password)
	if err != nil {
		return err
	}

	return u.userRepository.Save(ctx, user)
}

// NewUserCreateUseCase new instance of UserCreateUseCase
func NewUserCreateUseCase(userBlackList ports.UserBlackListerInterface, userRepository ports.UserRepositoryInterface) *UserCreateUseCase {
	return &UserCreateUseCase{
		userBlackList:  userBlackList,
		userRepository: userRepository,
	}
}
