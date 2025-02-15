package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	httpcontext "github.com/guillospy92/crabi/internal/adapters/http/context"
	httpdto "github.com/guillospy92/crabi/internal/adapters/http/dto"
	httpresponses "github.com/guillospy92/crabi/internal/adapters/http/responses"
	"github.com/guillospy92/crabi/internal/app/usecases"
	"github.com/guillospy92/logger"
)

// GetUserInfoHandler handler get user info api
type GetUserInfoHandler struct {
	getUserInfoUseCase usecases.UserGetInfoUseCaseInterface
}

// Handler expose api crete user
func (g *GetUserInfoHandler) Handler(c *fiber.Ctx) error {
	ctx := httpcontext.GetContextApplication(c)

	user := httpcontext.GetContextUser(c)

	response, err := g.getUserInfoUseCase.GetUserInfo(ctx, user.Email)
	if err != nil {
		logger.Log().LogAttrs(ctx, slog.LevelError, "CreateUserHandler error validate save user", slog.Any("err", err))
		return httpresponses.EvaluateError(c, err)
	}

	return httpresponses.HandleSuccessWithData(c, "get user info success", httpdto.UserEntityToUserDTOResponse(response))
}

// NewGetUserInfoHandler new instance of GetUserInfoHandler
func NewGetUserInfoHandler(getUserInfoUseCase usecases.UserGetInfoUseCaseInterface) *GetUserInfoHandler {
	return &GetUserInfoHandler{
		getUserInfoUseCase: getUserInfoUseCase,
	}
}
