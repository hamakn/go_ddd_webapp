package user

import "context"

// Repository is interface of user repository
type Repository interface {
	GetAll() ([]*User, error)
	GetByID(id int64) (*User, error)
	CreateFixture() ([]*User, error)
}

// NewRepository returns repository
// DI from infrastructure
var NewRepository func(context.Context) Repository
