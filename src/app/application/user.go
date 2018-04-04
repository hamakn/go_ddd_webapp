package application

import (
	"context"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
	"github.com/pkg/errors"
)

var (
	// ErrGetUsers is error on GetUsers
	ErrGetUsers = errors.New("app-application-user: GetUsers failed")
)

// GetUsers returns users
func GetUsers(ctx context.Context) ([]*user.User, error) {
	users, err := user.NewRepository(ctx).GetAll()
	if err != nil {
		return nil, errors.Wrap(err, ErrGetUsers.Error())
	}

	return users, nil
}
