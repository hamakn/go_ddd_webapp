package user

import "context"

// Repository is interface of user repository
type Repository interface {
	GetAll(context.Context) ([]*User, error)
	GetByID(context.Context, int64) (*User, error)
	Create(context.Context, *User) error
	Update(context.Context, *User) error
	Delete(context.Context, *User) error
	CreateFixture(context.Context) ([]*User, error)
}

// NewRepository returns Repository
// DI from infrastructure
var NewRepository func() Repository
