package transaction

import (
	"context"

	"github.com/NpoolPlatform/fox-proxy/pkg/db"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

// GetTransaction ..
func GetTransaction(ctx context.Context, transactionID string) (*foxproxy.Transaction, error) {
	tx, err := getTransaction(ctx, transactionID)
	if err != nil {
		return nil, err
	}

	return EntToProto(tx), nil
}

func getTransaction(ctx context.Context, transactionID string) (*ent.Transaction, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}

	tx, err := client.
		Transaction.
		Query().
		Where(
			transaction.TransactionIDEQ(transactionID),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

type GetTransactionExistParam struct {
	TransactionID    string
	TransactionState foxproxy.TransactionState
}

func GetTransactionExist(ctx context.Context, params GetTransactionExistParam) (bool, error) {
	client, err := db.Client()
	if err != nil {
		return false, err
	}

	stm := client.
		Transaction.
		Query().
		Select(transaction.FieldID).
		Where(
			transaction.TransactionIDEQ(params.TransactionID),
		)

	if params.TransactionState != foxproxy.TransactionState_TransactionStateUnKnow {
		stm.Where(transaction.StateEQ(int32(params.TransactionState)))
	}

	return stm.Exist(ctx)
}

func EntToProto(tx *ent.Transaction) *foxproxy.Transaction {
	return &foxproxy.Transaction{
		TransactionID: tx.TransactionID,
		Name:          tx.Name,
		Amount:        price.DBPriceToVisualPrice(tx.Amount),
		Payload:       tx.Payload,
		From:          tx.From,
		To:            tx.To,
		Memo:          tx.Memo,

		ExitCode: tx.ExitCode,
		CID:      tx.Cid,
		State:    foxproxy.TransactionState(tx.State),

		LockTime:  tx.LockTime,
		CreatedAt: tx.CreatedAt,
		UpdatedAt: tx.UpdatedAt,
	}
}

func EntsToProtos(txs []*ent.Transaction) []*foxproxy.Transaction {
	ret := []*foxproxy.Transaction{}
	for _, v := range txs {
		ret = append(ret, EntToProto(v))
	}
	return ret
}
