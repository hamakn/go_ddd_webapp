package db

import (
	"context"

	"github.com/mjibson/goon"
	"google.golang.org/appengine/datastore"
)

// GetAll get objects by query
func GetAll(ctx context.Context, q *datastore.Query, dst interface{}) ([]*datastore.Key, error) {
	g := goon.FromContext(ctx)
	return g.GetAll(q, dst)
}

// PutMulti puts src objects to db
func PutMulti(ctx context.Context, src interface{}) ([]*datastore.Key, error) {
	g := goon.FromContext(ctx)
	return g.PutMulti(src)
}
