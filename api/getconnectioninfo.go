//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/fox-proxy/pkg/deserver"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

func (s *Server) GetClientInfos(ctx context.Context, in *foxproxy.GetClientInfosRequest) (*foxproxy.GetClientInfosResponse, error) {
	infos := deserver.GetDEServerMGR().GetClientInfos()
	return &foxproxy.GetClientInfosResponse{
		Infos: infos,
		Total: uint32(len(infos)),
	}, nil
}
