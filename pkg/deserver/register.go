package deserver

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/fox-proxy/pkg/registercoin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

func init() {
	mgr := GetDEHandlerMGR()
	mgr.RegisterDEHandler(
		foxproxy.MsgType_MsgTypeRegisterCoin,
		&foxproxy.RegisterCoinInfo{},
		func(ctx context.Context, req interface{}) (interface{}, *foxproxy.StatusCode, error) {
			info, ok := req.(*foxproxy.RegisterCoinInfo)
			if !ok {
				return nil, foxproxy.StatusCode_StatusCodeUnmarshalErr.Enum(), fmt.Errorf("cannot transfer payload to req")
			}
			logger.Sugar().Error(info)
			_, err := registercoin.RegisterCoin(ctx, req.(*foxproxy.RegisterCoinInfo))
			if err != nil {
				return nil, foxproxy.StatusCode_StatusCodeFailed.Enum(), err
			}
			return nil, nil, nil
		},
	)
}
