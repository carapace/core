package core

import (
	"context"
	"database/sql"
)

type ContextKey int

const (
	txkey ContextKey = 1
)

func ContextWithTransaction(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txkey, tx)
}

func TXFromContext(ctx context.Context) *sql.Tx {
	tx, ok := ctx.Value(txkey).(*sql.Tx)
	if !ok {
		panic("context values corrupted; TX key does not contain *sql.Tx")
	}
	return tx
}
