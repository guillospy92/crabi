package adapterapi

import (
	"github.com/guillospy92/clientHttp/gohttp"
	"github.com/guillospy92/crabi/internal/core/domain"
	"github.com/guillospy92/crabi/resources"
)

// UserBlackListAdapter struct implement UserBlackListerInterface
type UserBlackListAdapter struct {
	clientHTTP gohttp.ClientInterface
}

type responseBlackList struct {
	IsBlackListed bool `json:"is_in_blacklist"`
}

// VerifyUserBlackList verify if a user is blocked
func (u *UserBlackListAdapter) VerifyUserBlackList(user domain.UserEntity) (bool, error) {
	resp, err := u.clientHTTP.Post(resources.ConfigurationEnv().URLAPIVerifyUserBlocked, nil, user)
	if err != nil {
		return false, err
	}

	var responseBody responseBlackList

	err = resp.UnMarshal(&responseBody)
	if err != nil {
		return false, err
	}

	return responseBody.IsBlackListed, nil
}

// NewUserBlackListAdapter new instance of UserBlackListAdapter
func NewUserBlackListAdapter(clientHTTP gohttp.ClientInterface) *UserBlackListAdapter {
	return &UserBlackListAdapter{
		clientHTTP: clientHTTP,
	}
}
