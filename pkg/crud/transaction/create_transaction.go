package transaction

import (
	"context"

	"github.com/NpoolPlatform/fox-proxy/pkg/db"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

type CreateTransactionParam struct {
	CoinType      foxproxy.CoinType
	ChainType     foxproxy.ChainType
	ClientType    foxproxy.ClientType
	State         foxproxy.TransactionState
	TransactionID string
	Name          string
	From          string
	To            string
	Value         float64
	Memo          string
}

func CreateTransaction(ctx context.Context, t *CreateTransactionParam) error {
	client, err := db.Client()
	if err != nil {
		return err
	}
	_, err = client.Transaction.Create().
		SetCoinType(int32(t.CoinType)).
		SetChainType(int32(t.ChainType)).
		SetClientType(int32(t.ClientType)).
		SetTransactionID(t.TransactionID).
		SetName(t.Name).
		SetFrom(t.From).
		SetTo(t.To).
		SetMemo(t.Memo).
		SetAmount(price.VisualPriceToDBPrice(t.Value)).
		SetState(int32(t.State)).
		Save(ctx)
	return err
}
