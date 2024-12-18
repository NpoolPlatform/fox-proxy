package api

import (
	"context"

	"github.com/NpoolPlatform/fox-proxy/api/stream"
	"github.com/NpoolPlatform/message/npool/foxproxy"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// Server ..
type Server struct {
	foxproxy.UnimplementedFoxProxyServer
}

func Register(server grpc.ServiceRegistrar) {
	foxproxy.RegisterFoxProxyServer(server, &Server{})
	foxproxy.RegisterFoxProxyStreamServer(server, &stream.Server{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return foxproxy.RegisterFoxProxyHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
