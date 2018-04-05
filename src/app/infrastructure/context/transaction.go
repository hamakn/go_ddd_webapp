package context

import (
	"context"
)

type contextKey string

const (
	transactionKey = contextKey("appTransaction")
)

// IsInTransaction returns in txn or not
func IsInTransaction(ctx context.Context) bool {
	value := ctx.Value(transactionKey)

	if value == nil {
		return false
	}

	return value.(bool)
}

// WithTransaction returns ctx marked in txn
func WithTransaction(ctx context.Context) context.Context {
	return context.WithValue(ctx, transactionKey, true)
}
