package registercoin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	coincli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

func RegisterCoin(ctx context.Context, in *foxproxy.RegisterCoinInfo) (bool, error) {
	if in == nil {
		return false, wlog.Errorf("invalid info for register info")
	}

	if exist, err := haveCoin(ctx, in.Name); exist {
		return true, nil
	} else if err != nil {
		return false, wlog.Unwrap(err)
	}

	chainType := in.ChainType.String()
	_, err := coincli.CreateCoin(ctx, &coinpb.CoinReq{
		Name:                &in.Name,
		Unit:                &in.Unit,
		ENV:                 &in.ENV,
		ChainType:           &chainType,
		ChainNativeUnit:     &in.ChainNativeUnit,
		ChainAtomicUnit:     &in.ChainAtomicUnit,
		ChainUnitExp:        &in.ChainUnitExp,
		GasType:             &in.GasType,
		ChainID:             &in.ChainID,
		ChainNickname:       &in.ChainNickname,
		ChainNativeCoinName: &in.ChainNativeCoinName,
	})
	if err != nil {
		return false, wlog.Unwrap(err)
	}
	return false, nil
}

func haveCoin(ctx context.Context, name string) (bool, error) {
	_, total, err := coincli.GetCoins(ctx, &coinpb.Conds{
		Name: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: name,
		},
	}, 0, 1)
	if err != nil {
		return false, err
	} else if total > 0 {
		return true, nil
	}
	return false, nil
}
