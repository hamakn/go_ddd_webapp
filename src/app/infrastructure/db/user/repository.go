package user

import (
	"context"
	"time"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
	appDatastore "github.com/hamakn/go_ddd_webapp/src/app/infrastructure/datastore"
	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/db"
	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/fixture"
	"golang.org/x/sync/errgroup"
	"google.golang.org/appengine/datastore"
)

const (
	kind = "User"
)

type repository struct {
}

// NewRepository returns user.Repository
func NewRepository() user.Repository {
	return &repository{}
}

func (r *repository) GetAll(ctx context.Context) ([]*user.User, error) {
	users := []*user.User{}
	q := datastore.NewQuery(kind)

	_, err := db.GetAll(ctx, q, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) GetByID(ctx context.Context, id int64) (*user.User, error) {
	u := &user.User{ID: id}
	err := db.Get(ctx, u)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, user.ErrNoSuchEntity
		}
		return nil, err
	}

	return u, nil
}

func (r *repository) Create(ctx context.Context, u *user.User) error {
	err := u.Validate()
	if err != nil {
		return user.ErrValidationFailed
	}

	return appDatastore.RunInTransaction(ctx, func(tctx context.Context) error {
		var eg errgroup.Group

		eg.Go(func() error {
			// check email uniqueness
			tctx := tctx
			if !canTakeUserEmail(tctx, u.Email) {
				return user.ErrEmailCannotTake
			}
			return nil
		})
		eg.Go(func() error {
			// check screen_name uniquness
			tctx := tctx
			if !canTakeUserScreenName(tctx, u.ScreenName) {
				return user.ErrScreenNameCannotTake
			}
			return nil
		})
		eg.Go(func() error {
			// Put user
			tctx := tctx
			return db.Put(tctx, u)
		})
		if err := eg.Wait(); err != nil {
			return err
		}

		// create userEmail and userScreenName after u.ID assigned
		userEmail := newUserEmail(u)
		userScreenName := newUserScreenName(u)

		eg.Go(func() error {
			// take email
			tctx := tctx
			err = takeUserEmail(tctx, userEmail)
			if err != nil {
				return user.ErrEmailCannotTake
			}
			return nil
		})
		eg.Go(func() error {
			// take screen_name
			tctx := tctx
			err = takeUserScreenName(tctx, userScreenName)
			if err != nil {
				return user.ErrScreenNameCannotTake
			}
			return nil
		})
		return eg.Wait()
	},
		true, // XG
	)
}

func (r *repository) Update(ctx context.Context, u *user.User) error {
	err := u.Validate()
	if err != nil {
		return user.ErrValidationFailed
	}

	return appDatastore.RunInTransaction(ctx, func(tctx context.Context) error {
		// get oldUser to get old email and screen name
		oldUser, err := r.GetByID(tctx, u.ID)
		if err != nil {
			return err
		}

		var eg errgroup.Group

		eg.Go(func() error {
			// userEmail
			tctx := tctx
			if oldUser.Email != u.Email {
				return updateUserEmail(tctx, u, oldUser.Email)
			}
			return nil
		})
		eg.Go(func() error {
			// userScreenName
			tctx := tctx
			if oldUser.ScreenName != u.ScreenName {
				return updateUserScreenName(tctx, u, oldUser.ScreenName)
			}
			return nil
		})
		eg.Go(func() error {
			// Update user
			tctx := tctx

			// TODO: rollback if error occurred
			u.UpdatedAt = time.Now()

			return db.Put(tctx, u)
		})
		return eg.Wait()
	},
		true, // XG
	)
}

func (r *repository) Delete(ctx context.Context, u *user.User) error {
	return appDatastore.RunInTransaction(ctx, func(tctx context.Context) error {
		var eg errgroup.Group

		eg.Go(func() error {
			// userEmail
			tctx := tctx
			return deleteUserEmail(tctx, u.Email, u.ID)
		})
		eg.Go(func() error {
			// userScreenName
			tctx := tctx
			return deleteUserScreenName(tctx, u.ScreenName, u.ID)
		})
		eg.Go(func() error {
			// delete user
			tctx := tctx
			// lock user
			txu, err := r.GetByID(tctx, u.ID)
			if err != nil {
				return err
			}
			return db.Delete(tctx, txu)
		})
		return eg.Wait()
	},
		true, // XG
	)
}

func (r *repository) CreateFixture(ctx context.Context) ([]*user.User, error) {
	users := []*user.User{}

	err := fixture.Load("users", &users)
	if err != nil {
		return nil, err
	}

	// NOTE: run in out of txn by datastore (25 entities) limit
	for _, u := range users {
		r.Create(ctx, u)
	}

	return users, nil
}
