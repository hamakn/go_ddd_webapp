package user

import "time"

// User Entity
type User struct {
	ID        int64     `yaml:"id" datastore:"-" goon:"id"`
	Email     string    `yaml:"email" validate:"required,email"`
	NickName  string    `yaml:"nickname" validate:"required"`
	Age       int       `yaml:"age" validate:"required,min=0,max=120"`
	CreatedAt time.Time `yaml:"created_at" validate:"required"`
	UpdatedAt time.Time `yaml:"updated_at" validate:"required"`
}
