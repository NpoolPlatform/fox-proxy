package stream

import (
	"github.com/NpoolPlatform/fox-proxy/pkg/deserver"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

func (s *Server) DEStream(stream foxproxy.FoxProxyStream_DEStreamServer) error {
	conn, err := deserver.RegisterDEServer(stream)
	if err != nil {
		logger.Sugar().Error(err)
		return err
	}
	defer conn.Close()

	err = deserver.GetDEServerMGR().AddDEServer(conn)
	if err != nil {
		logger.Sugar().Error(err)
		return err
	}

	conn.OnRecv()
	return nil
}
