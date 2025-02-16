package handlers_test

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/guillospy92/crabi/bootstrap"
	"github.com/guillospy92/crabi/internal/adapters/http/dto"
	"github.com/guillospy92/crabi/internal/adapters/http/handlers"
	"github.com/guillospy92/crabi/internal/app/usecases/mocks"
	"github.com/guillospy92/crabi/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthUserHandler_Handler(t *testing.T) {
	t.Parallel()
	type fields struct {
		userAuthUseCase    *usecasesmock.UserAuthUseCaseInterfaceMock
		userAuthRequestDTO httpdto.UserAuthRequestDTO
	}

	tests := []struct {
		mock               func(f fields)
		fields             fields
		name               string
		body               []byte
		responseStatusCode int
	}{
		{
			name: "TestAuthUserHandler_Handler success",
			fields: fields{
				userAuthUseCase: usecasesmock.NewUserAuthUseCaseInterfaceMock(t),
				userAuthRequestDTO: httpdto.UserAuthRequestDTO{
					Email:    "test@test.com",
					Password: "test_test",
				},
			},
			mock: func(f fields) {
				f.userAuthUseCase.On("Login", mock.Anything, f.userAuthRequestDTO.Email, f.userAuthRequestDTO.Password).
					Return(&domain.UserAuth{
						Token: domain.Token{},
						User:  &domain.UserEntity{},
					}, nil)
			},
			body:               []byte(`{"email": "test@test.com", "password": "test_test"}`),
			responseStatusCode: http.StatusOK,
		},
		{
			name: "TestAuthUserHandler_Handler unexpected error",
			fields: fields{
				userAuthUseCase: usecasesmock.NewUserAuthUseCaseInterfaceMock(t),
				userAuthRequestDTO: httpdto.UserAuthRequestDTO{
					Email:    "test@test.com",
					Password: "test_test",
				},
			},
			mock: func(f fields) {
				f.userAuthUseCase.On("Login", mock.Anything, f.userAuthRequestDTO.Email, f.userAuthRequestDTO.Password).
					Return(nil, errors.New("test error"))
			},
			body:               []byte(`{"email": "test@test.com", "password": "test_test"}`),
			responseStatusCode: http.StatusInternalServerError,
		},
		{
			name: "TestAuthUserHandler_Handler required params request",
			fields: fields{
				userAuthUseCase: usecasesmock.NewUserAuthUseCaseInterfaceMock(t),
			},
			mock:               func(fields) {},
			body:               []byte(`{"email": "test@test.com"`),
			responseStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.mock(tt.fields)
			authUserHandler := handlers.NewAuthUserHandler(tt.fields.userAuthUseCase)
			configTest := bootstrap.ConfigTest{Route: "/login", Handler: authUserHandler.Handler, Method: http.MethodPost}
			app := fiber.New()
			app.Post(configTest.Route, configTest.Handler)

			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(tt.body))
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
