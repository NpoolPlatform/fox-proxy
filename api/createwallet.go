//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/foxproxy"
)

func (s *Server) CreateWallet(ctx context.Context, in *foxproxy.CreateWalletRequest) (resp *foxproxy.CreateWalletResponse, err error) {
	resp = &foxproxy.CreateWalletResponse{}
	err = s.Forwarder(ctx, in.Name, foxproxy.MsgType_MsgTypeCreateWallet, in, resp)
	return
}
