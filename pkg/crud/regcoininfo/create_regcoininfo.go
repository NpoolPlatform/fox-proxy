package regcoininfo

import (
	"context"

	"github.com/NpoolPlatform/fox-proxy/pkg/db"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

func CreateRegCoinInfo(ctx context.Context, t *foxproxy.CoinInfo) error {
	client, err := db.Client()
	if err != nil {
		return err
	}
	return client.RegCoinInfo.Create().
		SetCoinType(int32(t.CoinType)).
		SetChainType(int32(t.ChainType)).
		SetName(t.Name).
		SetTempName(t.TempName).
		SetEnv(t.ENV).
		OnConflict().UpdateNewValues().Exec(ctx)
}
