package user

import (
	"context"
	"errors"
)

// Repository is interface of user repository
type Repository interface {
	GetAll() ([]*User, error)
	GetByID(id int64) (*User, error)
	Create(u *User) error
	CreateFixture() ([]*User, error)
}

var (
	// ErrNoSuchEntity is entity not found error
	ErrNoSuchEntity = errors.New("app-domain-user-repository: No such entity")
	// ErrEmailCannotTake is email cannot take error
	ErrEmailCannotTake = errors.New("app-domain-user-repository: Email cannot take")
	// ErrScreenNameCannotTake is screen_name cannot take error
	ErrScreenNameCannotTake = errors.New("app-domain-user-repository: ScreenName cannot take")
)

// NewRepository returns Repository
// DI from infrastructure
var NewRepository func(context.Context) Repository
