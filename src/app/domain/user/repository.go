package user

import (
	"context"
)

// Repository is interface of user repository
type Repository interface {
	GetAll() ([]*User, error)
	GetByID(id int64) (*User, error)
	Create(u *User) error
	Update(u *User) error
	Delete(u *User) error
	CreateFixture() ([]*User, error)
}

// NewRepository returns Repository
// DI from infrastructure
var NewRepository func(context.Context) Repository
