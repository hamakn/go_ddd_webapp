package user

import "context"

// Repository is interface of user repository
type Repository interface {
	CreateFixture() ([]*User, error)
}

// NewRepository returns repository
// DI from infrastructure
var NewRepository func(context.Context) Repository
