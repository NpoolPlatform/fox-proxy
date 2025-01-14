package deserver

import (
	"context"
	"fmt"

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
}
