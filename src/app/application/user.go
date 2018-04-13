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
	// ErrCreateUser is error on CreateUser
	ErrCreateUser = errors.New("app-application-user: CreateUser failed")
	// ErrUpdateUser is error on CreateUser
	ErrUpdateUser = errors.New("app-application-user: UpdateUser failed")
	// ErrDeleteUser is error on DeleteUser
	ErrDeleteUser = errors.New("app-application-user: DeleteUser failed")
)

// GetUsers returns users
func GetUsers(ctx context.Context) ([]*user.User, error) {
	users, err := user.NewRepository().GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, ErrGetUsers.Error())
	}

	return users, nil
}

// GetUserByID returns user specified by id
func GetUserByID(ctx context.Context, id int64) (*user.User, error) {
	u, err := user.NewRepository().GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, ErrGetUserByID.Error())
	}

	return u, nil
}

// CreateUser creates user from request
func CreateUser(ctx context.Context, req user.CreateUserValue) (*user.User, error) {
	u := user.NewFactory().Create(*req.Email, *req.ScreenName, *req.Age)

	err := user.NewRepository().Create(ctx, u)
	if err != nil {
		return nil, errors.Wrap(err, ErrCreateUser.Error())
	}

	return u, nil
}

// UpdateUser updates user from request
func UpdateUser(ctx context.Context, id int64, req user.UpdateUserValue) (*user.User, error) {
	r := user.NewRepository()
	u, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, ErrUpdateUser.Error())
	}

	hasUpdate := req.UpdateUser(u)
	if !hasUpdate {
		return nil, user.ErrNothingToUpdate
	}

	err = r.Update(ctx, u)
	if err != nil {
		return nil, errors.Wrap(err, ErrUpdateUser.Error())
	}

	return u, nil
}

// DeleteUser deletes user specified by id
func DeleteUser(ctx context.Context, id int64) error {
	r := user.NewRepository()
	u, err := r.GetByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, ErrDeleteUser.Error())
	}

	err = r.Delete(ctx, u)
	if err != nil {
		return errors.Wrap(err, ErrUpdateUser.Error())
	}

	return nil
}
