//go:build !codeanalysis
// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/fox-proxy/pkg/registercoin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/foxproxy"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RegisterCoin(ctx context.Context, in *foxproxy.RegisterCoinRequest) (*foxproxy.RegisterCoinResponse, error) {
	exist, err := registercoin.RegisterCoin(ctx, in.Info)
	if err != nil {
		logger.Sugar().Error(err)
		return &foxproxy.RegisterCoinResponse{}, status.Error(codes.Internal, "internal server error")
	}

	if exist {
		logger.Sugar().Infof("coin have been regitered, name: %v", in.Info.Name)
	}

	return &foxproxy.RegisterCoinResponse{}, nil
}
