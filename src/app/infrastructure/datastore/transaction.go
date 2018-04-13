package datastore

import (
	"context"

	appContext "github.com/hamakn/go_ddd_webapp/src/app/infrastructure/context"
	netContext "golang.org/x/net/context" // datastore lib doesn't support modern context yet, so use old one
	"google.golang.org/appengine/datastore"
)

// RunInTransaction is wraptter for datastore.RunInTransaction
// with resolving nested transaction and set attempts: 1
func RunInTransaction(ctx context.Context, f func(context.Context) error, xg bool) error {
	if !appContext.IsInTransaction(ctx) {
		// mark ctx as in txn
		ctx = appContext.WithTransaction(ctx)

		return datastore.RunInTransaction(ctx, func(tctx netContext.Context) error {
			// cast: old context => modern context
			mctx := context.Context(tctx) // for gsc
			return f(mctx)
		}, &datastore.TransactionOptions{
			XG:       xg,
			Attempts: 1,
		})
	}

	return f(ctx)
}
