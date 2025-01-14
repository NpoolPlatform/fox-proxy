package regcoininfo

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/fox-proxy/pkg/db"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/regcoininfo"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

// GetRegCoinInfo ..
func GetRegCoinInfo(ctx context.Context, name string) (*foxproxy.CoinInfo, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}

	ret, err := client.
		RegCoinInfo.
		Query().
		Select(
			regcoininfo.FieldChainType,
			regcoininfo.FieldCoinType,
			regcoininfo.FieldName,
			regcoininfo.FieldTempName,
			regcoininfo.FieldEnv,
		).
		Where(
			regcoininfo.NameEQ(name),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if ret == nil {
		return nil, fmt.Errorf("cannot get any coininfo for %v", name)
	}

	return &foxproxy.CoinInfo{
		Name:      ret.Name,
		TempName:  ret.TempName,
		CoinType:  foxproxy.CoinType(ret.CoinType),
		ChainType: foxproxy.ChainType(ret.ChainType),
		ENV:       ret.Env,
	}, nil
}

type GetRegCoinInfoExistParam struct {
	Name string
}

func GetRegCoinInfoExist(ctx context.Context, params GetRegCoinInfoExistParam) (bool, error) {
	client, err := db.Client()
	if err != nil {
		return false, err
	}

	stm := client.
		RegCoinInfo.
		Query().
		Select(regcoininfo.FieldID).
		Where(
			regcoininfo.NameEQ(params.Name),
		)

	return stm.Exist(ctx)
}
