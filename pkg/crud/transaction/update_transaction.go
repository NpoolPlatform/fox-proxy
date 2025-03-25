package transaction

import (
	"context"

	"github.com/NpoolPlatform/fox-proxy/pkg/db"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

// update nonce/utxo and state
type UpdateTransactionParams struct {
	TransactionID string
	State         foxproxy.TransactionState
	Payload       []byte
	Cid           string
	ExitCode      int64
	LockTime      uint32
}

// UpdateTransaction update transaction info
func UpdateTransaction(ctx context.Context, t *UpdateTransactionParams) error {
	client, err := db.Client()
	if err != nil {
		return err
	}

	stmt := client.
		Transaction.
		Update().
		Where(
			transaction.TransactionIDEQ(t.TransactionID),
		).
		SetPayload(t.Payload).
		SetCid(t.Cid).
		SetState(int32(t.State)).
		SetExitCode(t.ExitCode)

	if t.Cid != "" {
		stmt.
			SetCid(t.Cid)
	}

	return stmt.Exec(ctx)
}
