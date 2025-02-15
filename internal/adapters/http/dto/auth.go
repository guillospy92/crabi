package httpdto

import (
	"time"

	"github.com/guillospy92/crabi/internal/core/domain"
)

// UserAuthRequestDTO dto request auth user validation
type UserAuthRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TokenResponseDTO dto response token
type TokenResponseDTO struct {
	ExpiresIn   *time.Time `json:"expires_in"`
	AccessToken string     `json:"access_token"`
}

// UserAuthResponseDTO dto response auth
type UserAuthResponseDTO struct {
	User  *UserDTOResponse `json:"user"`
	Token TokenResponseDTO `json:"token"`
}

// UserAuthEntityToUserAuthDTOResponse convert userAuth entity to UserAuthResponseDTO
func UserAuthEntityToUserAuthDTOResponse(userAuth *domain.UserAuth) *UserAuthResponseDTO {
	return &UserAuthResponseDTO{
		Token: TokenResponseDTO(userAuth.Token),
		User:  UserEntityToUserDTOResponse(userAuth.User),
	}
}
