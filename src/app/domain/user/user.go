package user

import (
	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

// User Entity
type User struct {
	ID         int64     `json:"id" yaml:"id" datastore:"-" goon:"id"`
	Email      string    `json:"email" yaml:"email" validate:"required,email"`
	ScreenName string    `json:"screen_name" yaml:"screen_name" validate:"required,printascii,min=3,max=16"`
	Age        int       `json:"age" yaml:"age" validate:"required,min=0,max=120"`
	CreatedAt  time.Time `json:"created_at" yaml:"created_at" validate:"required"`
	UpdatedAt  time.Time `json:"updated_at" yaml:"updated_at" validate:"required"`
}

// Validate validates User
func (u *User) Validate() error {
	v := validator.New()
	return v.Struct(u)
}
