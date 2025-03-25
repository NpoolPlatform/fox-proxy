//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/foxproxy"
)

func (s *Server) GetBalance(ctx context.Context, in *foxproxy.GetBalanceRequest) (resp *foxproxy.GetBalanceResponse, err error) {
	resp = &foxproxy.GetBalanceResponse{}
	err = s.Forwarder(ctx, in.Name, foxproxy.MsgType_MsgTypeGetBalance, in, resp)
	return
}
