//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/foxproxy"
)

func (s *Server) GetEstimateGas(ctx context.Context, in *foxproxy.GetEstimateGasRequest) (resp *foxproxy.GetEstimateGasResponse, err error) {
	resp = &foxproxy.GetEstimateGasResponse{}
	err = s.Forwarder(ctx, in.Name, foxproxy.MsgType_MsgTypeGetEstimateGas, in, resp)
	return
}
