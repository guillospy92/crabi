package handlers_test

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/guillospy92/crabi/bootstrap"
	httpcontext "github.com/guillospy92/crabi/internal/adapters/http/context"
	"github.com/guillospy92/crabi/internal/adapters/http/handlers"
	usecasesmock "github.com/guillospy92/crabi/internal/app/usecases/mocks"
	"github.com/guillospy92/crabi/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserInfoHandler_Handler(t *testing.T) {
	t.Parallel()
	tests := []struct {
		getUserInfoUseCase *usecasesmock.UserGetInfoUseCaseInterfaceMock
		mock               func(getUserInfoUseCase *usecasesmock.UserGetInfoUseCaseInterfaceMock)
		name               string
		responseStatusCode int
	}{
		{
			name:               "TestGetUserInfoHandler_Handler success",
			getUserInfoUseCase: usecasesmock.NewUserGetInfoUseCaseInterfaceMock(t),
			mock: func(getUserInfoUseCase *usecasesmock.UserGetInfoUseCaseInterfaceMock) {
				getUserInfoUseCase.On("GetUserInfo", mock.Anything, mock.Anything).
					Return(
						&domain.UserEntity{},
						nil,
					)
			},
			responseStatusCode: http.StatusOK,
		},
		{
			name:               "TestGetUserInfoHandler_Handler error",
			getUserInfoUseCase: usecasesmock.NewUserGetInfoUseCaseInterfaceMock(t),
			mock: func(getUserInfoUseCase *usecasesmock.UserGetInfoUseCaseInterfaceMock) {
				getUserInfoUseCase.On("GetUserInfo", mock.Anything, mock.Anything).
					Return(
						nil,
						errors.New("mock error"),
					)
			},
			responseStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.mock(tt.getUserInfoUseCase)
			authUserHandler := handlers.NewGetUserInfoHandler(tt.getUserInfoUseCase)
			configTest := bootstrap.ConfigTest{Route: "/user", Handler: authUserHandler.Handler, Method: http.MethodGet}

			app := fiber.New()
			app.Use(generateUserContextTest())
			app.Get(configTest.Route, configTest.Handler)

			req := httptest.NewRequest(http.MethodGet, "/user", nil)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			if err != nil {
				t.Errorf("error test template user auth handler %v", err)
			}

			defer func(body io.ReadCloser) {
				err := body.Close()
				if err != nil {
					log.Print("error close body")
				}
			}(resp.Body)

			assert.Equal(t, tt.responseStatusCode, resp.StatusCode)
		})
	}
}

func generateUserContextTest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(httpcontext.UserKeyContextKey, domain.UserEntity{})
		return c.Next()
	}
}
