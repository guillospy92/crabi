package httpdto

import "github.com/guillospy92/crabi/internal/core/domain"

// UserDTORequest dto request http
type UserDTORequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// ConvertUserEntity convert UserDTORequest to UserEntity
func (u *UserDTORequest) ConvertUserEntity() domain.UserEntity {
	return domain.UserEntity{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}
}

// UserDTOResponse dto response http
type UserDTOResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// UserEntityToUserDTOResponse convert user entity to UserDTOResponse
func UserEntityToUserDTOResponse(user *domain.UserEntity) *UserDTOResponse {
	return &UserDTOResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
