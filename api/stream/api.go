package stream

import (
	"github.com/NpoolPlatform/message/npool/foxproxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// Server ..
type Server struct {
	foxproxy.UnimplementedFoxProxyStreamServer
}

func Register(server grpc.ServiceRegistrar) {
	foxproxy.RegisterFoxProxyStreamServer(server, &Server{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
}
