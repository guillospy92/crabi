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
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

func TestCreateUserHandler_Handler(t *testing.T) {
	t.Parallel()
	type fields struct {
		createUserUseCase *usecasesmock.UserCreateUseCaseInterfaceMock
		userDTO           httpdto.UserDTORequest
	}

	tests := []struct {
		mock               func(f fields)
		fields             fields
		name               string
		body               []byte
		responseStatusCode int
	}{
		{
			name: "TestCreateUserHandler_Handler success",
			fields: fields{
				createUserUseCase: usecasesmock.NewUserCreateUseCaseInterfaceMock(t),
				userDTO:           httpdto.UserDTORequest{},
			},
			mock: func(f fields) {
				f.createUserUseCase.On("CreateUser", mock.Anything, mock.Anything).
					Return(nil)
			},
			body:               []byte(`{"first_name": "will", "last_name": "will", "email": "will@example.com", "password": "password"}`),
			responseStatusCode: http.StatusOK,
		},
		{
			name: "TestCreateUserHandler_Handler unexpected error",
			fields: fields{
				createUserUseCase: usecasesmock.NewUserCreateUseCaseInterfaceMock(t),
				userDTO:           httpdto.UserDTORequest{},
			},
			mock: func(f fields) {
				f.createUserUseCase.On("CreateUser", mock.Anything, mock.Anything).
					Return(errors.New("unexpected error"))
			},
			body:               []byte(`{"first_name": "will", "last_name": "will", "email": "will@example.com", "password": "password"}`),
			responseStatusCode: http.StatusInternalServerError,
		},
		{
			name: "TestAuthUserHandler_Handler required params request",
			fields: fields{
				createUserUseCase: usecasesmock.NewUserCreateUseCaseInterfaceMock(t),
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
			authUserHandler := handlers.NewCreateUserHandler(tt.fields.createUserUseCase)
			configTest := bootstrap.ConfigTest{Route: "/user", Handler: authUserHandler.Handler, Method: http.MethodPost}
			app := fiber.New()
			app.Post(configTest.Route, configTest.Handler)

			req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(tt.body))
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
