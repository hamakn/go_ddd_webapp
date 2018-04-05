package user

import "time"

// Factory is interface of user factory
type Factory interface {
	Create(email string, screenName string, age int) *User
}

type factory struct {
}

// NewFactory returns Factory
func NewFactory() Factory {
	return &factory{}
}

func (f *factory) Create(email string, screenName string, age int) *User {
	now := time.Now()
	return &User{
		Email:      email,
		ScreenName: screenName,
		Age:        age,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
