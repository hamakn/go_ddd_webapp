package user

import "time"

// Factory is interface of user factory
type Factory interface {
	Create(email string, nickname string, age int) *User
}

type factory struct {
}

// NewFactory returns Factory
func NewFactory() Factory {
	return &factory{}
}

func (f *factory) Create(email string, nickname string, age int) *User {
	now := time.Now()
	return &User{
		Email:     email,
		NickName:  nickname,
		Age:       age,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
