package user

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
	"google.golang.org/appengine/datastore"
)

// userEmail Entity for email uniqueness constraints
type userEmail struct {
	HashedEmail string `datastore:"-"`
	Email       string
	UserID      int64
	CreatedAt   time.Time
}

func newUserEmail(u *user.User) *userEmail {
	now := time.Now()
	return &userEmail{
		HashedEmail: hashEmail(u.Email),
		Email:       u.Email,
		UserID:      u.ID,
		CreatedAt:   now,
	}
}

func userEmailKey(ctx context.Context, email string) *datastore.Key {
	return datastore.NewKey(ctx, "UserEmail", email, 0, nil)
}

// hashEmail is base64 of sha256(email)
func hashEmail(email string) string {
	h := sha256.New()
	h.Write([]byte(email + "UserEmail"))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
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
