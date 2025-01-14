package transaction

import (
	"context"

	constant "github.com/NpoolPlatform/fox-proxy/pkg/const"
	"github.com/NpoolPlatform/fox-proxy/pkg/db"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

type GetTransactionsParam struct {
	CoinType         foxproxy.CoinType
	TransactionState foxproxy.TransactionState
}

// GetTransactions ..
func GetTransactions(ctx context.Context, params GetTransactionsParam) ([]*foxproxy.Transaction, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}

	stmt := client.
		Transaction.
		Query().
		Where(
			transaction.StateEQ(int32(params.TransactionState)),
		)

	if params.CoinType != foxproxy.CoinType_CoinTypeUnKnow {
		stmt = stmt.Where(transaction.CoinTypeEQ(int32(params.CoinType)))
	}

	txs, err := stmt.Order(ent.Asc(transaction.FieldCreatedAt)).
		Limit(constant.DefaultPageSize).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return EntsToProtos(txs), nil
}
