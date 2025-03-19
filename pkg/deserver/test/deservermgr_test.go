package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/NpoolPlatform/fox-proxy/api/stream"
	"github.com/NpoolPlatform/fox-proxy/pkg/deserver"
	testinit "github.com/NpoolPlatform/fox-proxy/pkg/test-init"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/message/npool/foxproxy"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcPort = 50023

var (
	clientType      = foxproxy.ClientType_ClientTypePlugin
	clientPosition  = "ssss"
	clientCoinInfos = []*foxproxy.CoinInfo{
		{Name: "ssss", ChainType: foxproxy.ChainType_Aleo, CoinType: foxproxy.CoinType_CoinTypealeo},
	}
)

// for test
func MockOnServer(ctx context.Context, grpcPort int) {
	fmt.Println("start to mock server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	defer server.Stop()

	stream.Register(server)

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()
	<-ctx.Done()
	fmt.Println("end to mock server")
}

// for test
func MockClient(ctx context.Context, grpcPort int) foxproxy.FoxProxyStream_DEStreamClient {
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("localhost:%d", grpcPort), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := foxproxy.NewFoxProxyStreamClient(conn)
	clientConn, err := client.DEStream(ctx)
	if err != nil {
		panic(err)
	}

	err = RegisterDEClient(clientConn, clientType, clientPosition, clientCoinInfos)
	if err != nil {
		panic(err)
	}

	return clientConn
}

// for test
func RegisterDEClient(
	client foxproxy.FoxProxyStream_DEStreamClient,
	clientType foxproxy.ClientType,
	position string,
	infos []*foxproxy.CoinInfo,
) error {
	select {
	case <-time.NewTicker(time.Second * 3).C:
		return wlog.Errorf("timeout for register connection")
	default:
		if len(infos) == 0 {
			return wlog.Errorf("have no infos")
		}

		msgID := uuid.NewString()
		connID := uuid.NewString()
		connInfo := &foxproxy.ClientInfo{
			ClientType: clientType,
			ID:         connID,
			Infos:      infos,
			Position:   position,
		}
		payload, err := json.Marshal(connInfo)
		if err != nil {
			return wlog.WrapError(err)
		}

		err = client.Send(&foxproxy.DataElement{
			ConnectID: connID,
			MsgID:     msgID,
			Payload:   payload,
		})
		if err != nil {
			return wlog.WrapError(err)
		}

		data, err := client.Recv()
		if err != nil {
			return wlog.WrapError(err)
		}

		if data.ErrMsg != nil && *data.ErrMsg != "" {
			return wlog.Errorf("failed to register to proxy, err: %v", *data.ErrMsg)
		}

		return nil
	}
}

func TestDEServerMGR(t *testing.T) {
	err := testinit.Init()
	if !assert.Nil(t, err) {
		return
	}
	err = logger.Init(logger.DebugLevel, "./a.log")
	if !assert.Nil(t, err) {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go MockOnServer(ctx, grpcPort)
	time.Sleep(time.Second * 3)
	cli := MockClient(ctx, grpcPort)
	time.Sleep(time.Second * 3)
	mgr := deserver.GetDEServerMGR()

	infos := mgr.GetClientInfos()
	if !assert.NotEqual(t, 0, len(infos)) {
		return
	}
	clientInfo := infos[0]

	msgInfo := deserver.MsgInfo{
		Payload: []byte("payload"),
	}
	err = mgr.SendMsg(clientCoinInfos[0].Name, clientType, foxproxy.MsgType_MsgTypeDefault, nil, nil, &msgInfo)
	if !assert.Nil(t, err) {
		return
	}

	dataEle, err := cli.Recv()
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, dataEle.ConnectID, clientInfo.ID)
	assert.Equal(t, dataEle.Payload, msgInfo.Payload)

	payload := []byte{0, 1, 2, 3}
	msgID := uuid.NewString()
	err = cli.Send(&foxproxy.DataElement{
		ConnectID: infos[0].ID,
		MsgID:     msgID,
		MsgType:   foxproxy.MsgType_MsgTypeEcho,
		Payload:   payload,
	})
	if !assert.Nil(t, err) {
		return
	}

	dataEle, err = cli.Recv()
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, dataEle.ConnectID, clientInfo.ID)
	assert.Equal(t, dataEle.Payload, payload)
	assert.Equal(t, dataEle.MsgID, msgID)

	mgr.CloseAll()
	infos = mgr.GetClientInfos()
	assert.Equal(t, 0, len(infos))

	dataEle, err = cli.Recv()
	assert.NotNil(t, err)
	assert.Nil(t, dataEle)
}
