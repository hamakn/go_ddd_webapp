package user

import (
	"context"
	"time"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
	"google.golang.org/appengine/datastore"
)

// userScreenName Entity for screen_name uniqueness constraints
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

func userScreenNameKey(ctx context.Context, screenName string) *datastore.Key {
	return datastore.NewKey(ctx, "UserScreenName", screenName, 0, nil)
}

func canTakeUserScreenName(ctx context.Context, screenName string) bool {
	k := userScreenNameKey(ctx, screenName)
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
