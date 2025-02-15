package ports

import "github.com/guillospy92/crabi/internal/core/domain"

// UserBlackListerInterface check if a user is on a blacklist
type UserBlackListerInterface interface {
	VerifyUserBlackList(user domain.UserEntity) (bool, error)
}
