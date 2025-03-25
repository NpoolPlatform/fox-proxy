package transaction

import (
	"context"

	"github.com/NpoolPlatform/fox-proxy/pkg/db"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

type GetTransactionsParam struct {
	CoinType         foxproxy.CoinType
	TransactionState foxproxy.TransactionState
	ENV              string
}

// GetTransactions ..
func GetTransactions(ctx context.Context, req *foxproxy.GetTransactionsRequest) ([]*foxproxy.Transaction, uint32, error) {
	client, err := db.Client()
	if err != nil {
		return nil, 0, err
	}

	eqConds := setQueryParms(req)

	stm := client.
		Transaction.
		Query().
		Where(
			eqConds...,
		).
		Order(ent.Asc(transaction.FieldCreatedAt))
	total, err := stm.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	txs, err := stm.
		Offset(int(req.Offset)).
		Limit(int(req.Limit)).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return EntsToProtos(txs), uint32(total), nil
}

func setQueryParms(req *foxproxy.GetTransactionsRequest) []predicate.Transaction {
	ret := []predicate.Transaction{}
	if req.Name != nil {
		ret = append(ret, transaction.NameEQ(*req.Name))
	}
	if req.From != nil {
		ret = append(ret, transaction.FromEQ(*req.From))
	}
	if req.To != nil {
		ret = append(ret, transaction.ToEQ(*req.To))
	}
	if req.CID != nil {
		ret = append(ret, transaction.CidEQ(*req.CID))
	}
	if req.Memo != nil {
		ret = append(ret, transaction.MemoEQ(*req.Memo))
	}
	if req.Amount != nil {
		ret = append(ret, transaction.AmountEQ(price.VisualPriceToDBPrice(*req.Amount)))
	}
	if req.CoinType != nil {
		ret = append(ret, transaction.CoinTypeEQ(int32(*req.CoinType.Enum())))
	}
	if req.ChainType != nil {
		ret = append(ret, transaction.ChainTypeEQ(int32(*req.ChainType.Enum())))
	}
	if req.State != nil {
		ret = append(ret, transaction.StateEQ(int32(*req.State.Enum())))
	}
	return ret
}
