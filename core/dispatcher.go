package core

import (
	"context"

	"github.com/carapace/core/api/v0/proto"
)

type Dispatcher interface {
	TransactionService(context.Context, *v0.Transaction) (res *v0.Transaction, err error)
}
