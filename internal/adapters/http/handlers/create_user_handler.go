package handlers

import (
	"context"
	_ "embed"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/guillospy92/crabi/internal/adapters/http/context"
	"github.com/guillospy92/crabi/internal/adapters/http/dto"
	"github.com/guillospy92/crabi/internal/adapters/http/responses"
	"github.com/guillospy92/crabi/internal/app/usecases"
	"github.com/guillospy92/logger"
)

// CreateUserHandler handler create user api
type CreateUserHandler struct {
	createUserUseCase usecases.UserCreateUseCaseInterface
}

//go:embed schema_save_required.json
var jsonSchemaSaveUser string

// Handler expose api crete user
func (h *CreateUserHandler) Handler(c *fiber.Ctx) error {
	ctx := httpcontext.GetContextApplication(c)

	var userDTO httpdto.UserDTORequest
	_ = c.BodyParser(&userDTO)

	if errValidation := ValidateSchemeJSON(userDTO, jsonSchemaSaveUser); errValidation.Err != nil {
		logger.Log().LogAttrs(ctx, slog.LevelInfo, "CreateUserHandler error validate request", slog.Any("err", errValidation.Err))
		return c.Status(http.StatusBadRequest).JSON(httpresponses.ResponseErrorWithAttribute{
			ErrorCode:     httpresponses.ErrorCodeBadFieldRequest,
			StatusCode:    http.StatusBadRequest,
			Message:       "error validate request",
			MessageErrors: errValidation.TransformErrors(),
		})
	}

	err := h.createUserUseCase.CreateUser(context.Background(), userDTO.ConvertUserEntity())
	if err != nil {
		logger.Log().LogAttrs(ctx, slog.LevelError, "CreateUserHandler error validate save user", slog.Any("err", err))
		return httpresponses.EvaluateError(c, err)
	}

	return httpresponses.HandleSuccess(c, "user created")
}

// NewCreateUserHandler new instance of CreateUserHandler
func NewCreateUserHandler(createUserUseCase usecases.UserCreateUseCaseInterface) *CreateUserHandler {
	return &CreateUserHandler{
		createUserUseCase: createUserUseCase,
	}
}
