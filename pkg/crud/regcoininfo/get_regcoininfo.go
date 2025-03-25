package regcoininfo

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/fox-proxy/pkg/db"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent"
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

	return EntToProto(ret), nil
}

// GetRegCoinInfo ..
func GetRegCoinInfos(ctx context.Context, names []string) ([]*foxproxy.CoinInfo, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}

	ret, err := client.
		RegCoinInfo.
		Query().
		Where(
			regcoininfo.TempNameIn(names...),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	if ret == nil {
		return nil, fmt.Errorf("cannot get any coininfo for %v", names)
	}

	return EntsToProtos(ret), nil
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

func EntToProto(info *ent.RegCoinInfo) *foxproxy.CoinInfo {
	return &foxproxy.CoinInfo{
		Name:      info.Name,
		TempName:  info.TempName,
		CoinType:  foxproxy.CoinType(info.CoinType),
		ChainType: foxproxy.ChainType(info.ChainType),
		ENV:       info.Env,
	}
}

func EntsToProtos(infos []*ent.RegCoinInfo) []*foxproxy.CoinInfo {
	ret := []*foxproxy.CoinInfo{}
	for _, v := range infos {
		ret = append(ret, EntToProto(v))
	}
	return ret
}
