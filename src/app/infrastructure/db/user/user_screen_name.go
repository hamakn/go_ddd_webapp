package user

import (
	"context"
	"time"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
	"google.golang.org/appengine/datastore"
)

// userScreenName Entity for nickname uniqueness constraints
type userScreenName struct {
	ScreenName string
	UserID     int64
	CreatedAt  time.Time
}

func createUserScreenName(u *user.User) *userScreenName {
	now := time.Now()
	return &userScreenName{
		ScreenName: u.ScreenName,
		UserID:     u.ID,
		CreatedAt:  now,
	}
}

func userScreenNameKey(ctx context.Context, nickname string) *datastore.Key {
	return datastore.NewKey(ctx, "UserScreenName", nickname, 0, nil)
}

func canTakeUserScreenName(ctx context.Context, nickname string) bool {
	k := userScreenNameKey(ctx, nickname)
	err := datastore.Get(ctx, k, &userScreenName{})
	if err != nil && err.Error() == "datastore: no such entity" {
		return true
	}
	return false
}

func takeUserScreenName(ctx context.Context, u *userScreenName) error {
	k := userScreenNameKey(ctx, u.ScreenName)
	_, err := datastore.Put(ctx, k, u)
	return err
}
