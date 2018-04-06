package user

import (
	"context"
	"strings"
	"time"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
	"google.golang.org/appengine/datastore"
)

// userEmail Entity for email uniqueness constraints
type userEmail struct {
	Email     string
	UserID    int64
	CreatedAt time.Time
}

func newUserEmail(u *user.User) *userEmail {
	now := time.Now()
	return &userEmail{
		Email:     u.Email,
		UserID:    u.ID,
		CreatedAt: now,
	}
}

func userEmailKey(ctx context.Context, email string) *datastore.Key {
	return datastore.NewKey(ctx, "UserEmail", userEmailKeyString(email), 0, nil)
}

// userEmailKeyString is downcased email
func userEmailKeyString(email string) string {
	return strings.ToLower(email)
}

func canTakeUserEmail(ctx context.Context, email string) bool {
	k := userEmailKey(ctx, email)
	err := datastore.Get(ctx, k, &userEmail{})
	if err != nil && err.Error() == "datastore: no such entity" {
		return true
	}
	return false
}

func takeUserEmail(ctx context.Context, u *userEmail) error {
	k := userEmailKey(ctx, u.Email)
	_, err := datastore.Put(ctx, k, u)
	return err
}
