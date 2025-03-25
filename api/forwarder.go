package api

import (
	"context"
	"encoding/json"

	"github.com/NpoolPlatform/fox-proxy/pkg/crud/regcoininfo"
	"github.com/NpoolPlatform/fox-proxy/pkg/deserver"
	"github.com/NpoolPlatform/fox-proxy/pkg/router"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

func (s *Server) Forwarder(ctx context.Context, name string, msgType foxproxy.MsgType, req, resp interface{}) error {
	mgr := deserver.GetDEServerMGR()
	coinInfo, err := regcoininfo.GetRegCoinInfo(ctx, name)
	if err != nil {
		return err
	}

	steps, err := router.GetMsgTypeRouter().GetMsgTypeSteps(msgType, &coinInfo.ChainType, &coinInfo.CoinType)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	for _, step := range steps {
		respPayload, err := mgr.SendAndRecvRaw(ctx, name, step.ClientType, step.MsgType, payload)
		if err != nil {
			return err
		}
		payload = respPayload
	}

	return json.Unmarshal(payload, resp)
}
