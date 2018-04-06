package user

import "errors"

var (
	// ErrNoSuchEntity is entity not found error
	ErrNoSuchEntity = errors.New("app-domain-user: No such entity")
	// ErrEmailCannotTake is email cannot take error
	ErrEmailCannotTake = errors.New("app-domain-user: Email cannot take")
	// ErrScreenNameCannotTake is screen_name cannot take error
	ErrScreenNameCannotTake = errors.New("app-domain-user: ScreenName cannot take")
	// ErrNothingToUpdate is nothing to update error
	ErrNothingToUpdate = errors.New("app-domain-user-user: Nothing to update")
)
