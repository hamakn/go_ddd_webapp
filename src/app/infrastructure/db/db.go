package db

import (
	"context"

	"github.com/mjibson/goon"
	"google.golang.org/appengine/datastore"
)

// PutMulti puts src objects to db
func PutMulti(ctx context.Context, src interface{}) ([]*datastore.Key, error) {
	g := goon.FromContext(ctx)
	return g.PutMulti(src)
}
