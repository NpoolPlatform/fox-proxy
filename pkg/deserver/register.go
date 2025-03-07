package deserver

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/fox-proxy/pkg/crud/regcoininfo"
	"github.com/NpoolPlatform/fox-proxy/pkg/crud/transaction"
	"github.com/NpoolPlatform/fox-proxy/pkg/registercoin"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

func init() {
	mgr := GetDEHandlerMGR()
	mgr.RegisterDEHandler(
		foxproxy.MsgType_MsgTypeRegisterCoin,
		&[]*foxproxy.RegisterCoinInfo{},
		func(ctx context.Context, req interface{}) (interface{}, error) {
			infos, ok := req.(*[]*foxproxy.RegisterCoinInfo)
			if !ok {
				return nil, fmt.Errorf("cannot transfer payload to req")
			}

			for _, info := range *infos {
				_, err := registercoin.RegisterCoin(ctx, info)
				if err != nil {
					return nil, err
				}
			}
			return nil, nil
		},
	)

	mgr.RegisterDEHandler(
		foxproxy.MsgType_MsgTypeAssginPluginTxs,
		&[]*foxproxy.CoinInfo{},
		func(ctx context.Context, req interface{}) (interface{}, error) {
			infos, ok := req.(*[]*foxproxy.CoinInfo)
			if !ok {
				return nil, fmt.Errorf("cannot transfer payload to req")
			}
			names := []string{}
			for _, info := range *infos {
				names = append(names, info.Name)
			}
			return transaction.AssginTxs(ctx, foxproxy.ClientType_ClientTypePlugin, names)
		},
	)

	mgr.RegisterDEHandler(
		foxproxy.MsgType_MsgTypeAssginSignTxs,
		&[]*foxproxy.CoinInfo{},
		func(ctx context.Context, req interface{}) (interface{}, error) {
			infos, ok := req.(*[]*foxproxy.CoinInfo)
			if !ok {
				return nil, fmt.Errorf("cannot transfer payload to req")
			}

			_names := []string{}
			for _, info := range *infos {
				_names = append(_names, info.Name)
			}
			regInfos, err := regcoininfo.GetRegCoinInfos(ctx, _names)
			if err != nil {
				return nil, err
			}
			names := []string{}
			for _, info := range regInfos {
				names = append(names, info.Name)
			}

			return transaction.AssginTxs(ctx, foxproxy.ClientType_ClientTypeSign, names)
		},
	)

	mgr.RegisterDEHandler(
		foxproxy.MsgType_MsgTypeSubmitTx,
		&foxproxy.SubmitTransaction{},
		func(ctx context.Context, req interface{}) (interface{}, error) {
			info, ok := req.(*foxproxy.SubmitTransaction)
			if !ok {
				return nil, fmt.Errorf("cannot transfer payload to req")
			}

			return nil, transaction.SubmitTx(ctx, info)
		},
	)
}
