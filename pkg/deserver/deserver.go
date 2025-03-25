package deserver

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

type DEServer struct {
	foxproxy.FoxProxyStream_DEStreamServer
	*foxproxy.ClientInfo
	ctx          context.Context
	cancel       context.CancelFunc
	onCloseFuncs []func(conn *DEServer)
	recvHandlers []func(data *foxproxy.DataElement)
	closeOnce    sync.Once
}

func RegisterDEServer(stream foxproxy.FoxProxyStream_DEStreamServer) (*DEServer, error) {
	select {
	case <-time.NewTicker(time.Second * 3).C:
		return nil, wlog.Errorf("timeout for register connection")
	default:
		data, err := stream.Recv()
		if err != nil {
			return nil, wlog.WrapError(err)
		}

		statusMsg := ""

		connInfo := &foxproxy.ClientInfo{}
		err = json.Unmarshal(data.Payload, connInfo)
		if err != nil {
			statusMsg = err.Error()
		}

		if err == nil && data.ConnectID != connInfo.ID {
			statusMsg = "connectid dont match"
		}

		err = stream.Send(&foxproxy.DataElement{
			ConnectID: data.ConnectID,
			MsgID:     data.MsgID,
			ErrMsg:    &statusMsg,
		})
		if err != nil {
			return nil, wlog.WrapError(err)
		}

		if statusMsg != "" {
			return nil, wlog.Errorf(statusMsg)
		}

		ctx, cancel := context.WithCancel(stream.Context())
		return &DEServer{
			FoxProxyStream_DEStreamServer: stream,
			ClientInfo:                    connInfo,
			ctx:                           ctx,
			cancel:                        cancel,
		}, nil
	}
}

func (conn *DEServer) WatchClose(onClose func(conn *DEServer)) {
	conn.onCloseFuncs = append(conn.onCloseFuncs, onClose)
}

func (conn *DEServer) WatchRecv(onRecv func(data *foxproxy.DataElement)) {
	conn.recvHandlers = append(conn.recvHandlers, onRecv)
}

func (conn *DEServer) Close() {
	conn.closeOnce.Do(func() {
		for _, onClose := range conn.onCloseFuncs {
			onClose(conn)
		}
		conn.cancel()
		logger.Sugar().Warnf(
			"connection is closed, siderType: %v, ID: %v, Position: %v",
			conn.ClientType,
			conn.Position,
			conn.Position)
	})
}

func (conn *DEServer) OnRecv() {
	go func() {
		defer conn.Close()
		for {
			data, err := conn.Recv()
			if err != nil {
				logger.Sugar().Error(err)
				return
			}
			for _, recvHandler := range conn.recvHandlers {
				go recvHandler(data)
			}
		}
	}()

	<-conn.ctx.Done()
}
