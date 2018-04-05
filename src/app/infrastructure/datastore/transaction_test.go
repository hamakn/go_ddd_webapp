package datastore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

type Item struct {
	Message string
}

func TestTransactionNested(t *testing.T) {
	// setup ctx
	ctx, done, err := aetest.NewContext()
	defer done()
	require.Nil(t, err)

	id := int64(42)
	message := "hey"
	key := datastore.NewKey(ctx, "Item", "", id, nil)

	err = RunInTransaction(ctx, func(tctx context.Context) error {
		return RunInTransaction(tctx, func(ttctx context.Context) error {
			i := Item{message}
			_, err := datastore.Put(ttctx, key, &i)
			if err != nil {
				return err
			}

			return nil
		}, true)
	}, true)

	require.Nil(t, err)

	i := Item{}
	datastore.Get(ctx, key, &i)
	require.Equal(t, message, i.Message)
}
