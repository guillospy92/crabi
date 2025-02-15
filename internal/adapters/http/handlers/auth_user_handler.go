package handlers

import (
	_ "embed"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
	httpcontext "github.com/guillospy92/crabi/internal/adapters/http/context"
	httpdto "github.com/guillospy92/crabi/internal/adapters/http/dto"
	httpresponses "github.com/guillospy92/crabi/internal/adapters/http/responses"
	"github.com/guillospy92/crabi/internal/app/usecases"
	"github.com/guillospy92/logger"
)

//go:embed schema_auth_required.json
var jsonSchemaUserLogin string

// AuthUserHandler handler log in user api
type AuthUserHandler struct {
	userAuthUseCase usecases.UserAuthUseCaseInterface
}

// Handler expose api check login
func (a *AuthUserHandler) Handler(c *fiber.Ctx) error {
	ctx := httpcontext.GetContextApplication(c)

	var userAuthDTO httpdto.UserAuthRequestDTO
	_ = c.BodyParser(&userAuthDTO)

	if errValidation := ValidateSchemeJSON(userAuthDTO, jsonSchemaUserLogin); errValidation.Err != nil {
		logger.Log().LogAttrs(ctx, slog.LevelInfo, "AuthUserHandler error validate request", slog.Any("err", errValidation.Err))
		return c.Status(http.StatusBadRequest).JSON(httpresponses.ResponseErrorWithAttribute{
			ErrorCode:     httpresponses.ErrorCodeBadFieldRequest,
			StatusCode:    http.StatusBadRequest,
			Message:       "error validate request",
			MessageErrors: errValidation.TransformErrors(),
		})
	}

	response, err := a.userAuthUseCase.Login(ctx, userAuthDTO.Email, userAuthDTO.Password)
	if err != nil {
		logger.Log().LogAttrs(ctx, slog.LevelError, "AuthUserHandler error validate auth user login", slog.Any("err", err))
		return httpresponses.EvaluateError(c, err)
	}

	return httpresponses.HandleSuccessWithData(c, "login success", httpdto.UserAuthEntityToUserAuthDTOResponse(response))
}

// NewAuthUserHandler new instance of AuthUserHandler
func NewAuthUserHandler(userAuthUseCase usecases.UserAuthUseCaseInterface) *AuthUserHandler {
	return &AuthUserHandler{
		userAuthUseCase: userAuthUseCase,
	}
}
