package application

import (
	"context"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
	"github.com/pkg/errors"
)

var (
	// ErrGetUsers is error on GetUsers
	ErrGetUsers = errors.New("app-application-user: GetUsers failed")
	// ErrGetUserByID is error on GetUser
	ErrGetUserByID = errors.New("app-application-user: GetUserByID failed")
)

// GetUsers returns users
func GetUsers(ctx context.Context) ([]*user.User, error) {
	users, err := user.NewRepository(ctx).GetAll()
	if err != nil {
		return nil, errors.Wrap(err, ErrGetUsers.Error())
	}

	return users, nil
}

// GetUserByID returns user specified by id
func GetUserByID(ctx context.Context, id int64) (*user.User, error) {
	user, err := user.NewRepository(ctx).GetByID(id)
	if err != nil {
		if err.Error() == "datastore: no such entity" {
			return nil, err
		}
		return nil, errors.Wrap(err, ErrGetUserByID.Error())
	}

	return user, nil
}
