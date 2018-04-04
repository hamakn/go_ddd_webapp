package user

import (
	"context"

	"github.com/hamakn/go_ddd_webapp/src/app/domain/user"
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
