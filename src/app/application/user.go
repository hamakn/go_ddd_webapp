package application

import (
	"context"
	"time"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
)

// GetUsers returns users
func GetUsers(ctx context.Context) ([]*user.User, error) {
	users := []*user.User{
		{
			ID:        42,
			Email:     "foobar@hamakn.test",
			NickName:  "foobar",
			Age:       17,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return users, nil
}
