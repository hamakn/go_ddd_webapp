package user

import (
	"context"
	"errors"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
	appDatastore "github.com/hamakn/go_ddd_webapp/src/app/infrastructure/datastore"
	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/db"
	"github.com/hamakn/go_ddd_webapp/src/app/infrastructure/fixture"
	"google.golang.org/appengine/datastore"
)

const (
	kind = "User"
)

type repository struct {
	Ctx context.Context
}

// NewRepository returns user.Repository
func NewRepository(ctx context.Context) user.Repository {
	return &repository{Ctx: ctx}
}

func (r *repository) GetAll() ([]*user.User, error) {
	users := []*user.User{}
	q := datastore.NewQuery(kind)

	_, err := db.GetAll(r.Ctx, q, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) GetByID(id int64) (*user.User, error) {
	user := &user.User{ID: id}
	err := db.Get(r.Ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Create(u *user.User) error {
	return appDatastore.RunInTransaction(r.Ctx, func(tctx context.Context) error {
		// check email uniqueness
		if !canTakeUserEmail(tctx, u.Email) {
			return errors.New("app-infra-db-user-repository: Email cannot take")
		}

		// check nickname uniquness
		if !canTakeUserScreenName(tctx, u.ScreenName) {
			return errors.New("app-infra-db-user-repository: ScreenName cannot take")
		}

		err := db.Put(tctx, u)
		if err != nil {
			return err
		}

		// take email
		userEmail := createUserEmail(u)
		err = takeUserEmail(tctx, userEmail)
		if err != nil {
			return err
		}

		// take nickname
		userScreenName := createUserScreenName(u)
		err = takeUserScreenName(tctx, userScreenName)
		if err != nil {
			return err
		}

		return nil
	},
		true, // XG
	)
}

func (r *repository) CreateFixture() ([]*user.User, error) {
	users := []*user.User{}

	err := fixture.Load("users", &users)
	if err != nil {
		return nil, err
	}

	_, err = db.PutMulti(r.Ctx, users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
