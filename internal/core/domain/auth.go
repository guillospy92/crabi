package domain

import "time"

// Token entity token domain logic
type Token struct {
	ExpiresIn   *time.Time
	AccessToken string
}

// UserAuth entity user auth domain logic
type UserAuth struct {
	User  *UserEntity
	Token Token
}
