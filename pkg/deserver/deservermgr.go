package deserver

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/NpoolPlatform/fox-proxy/pkg/crud/regcoininfo"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/foxproxy"
	"github.com/google/uuid"
)

//nolint:revive
type DEServerMGR struct {
	connInfos   map[string]map[foxproxy.ClientType][]*DEServer
	recvChannel sync.Map
	connections []*DEServer
}

var cmgr *DEServerMGR

func GetDEServerMGR() *DEServerMGR {
	if cmgr == nil {
		cmgr = &DEServerMGR{
			connInfos:   make(map[string]map[foxproxy.ClientType][]*DEServer),
			recvChannel: sync.Map{},
		}
	}
	return cmgr
}

func (mgr *DEServerMGR) AddDEServer(conn *DEServer) error {
	for _, info := range conn.Infos {
		err := regcoininfo.CreateRegCoinInfo(conn.ctx, info)
		if err != nil {
			return err
		}
		if _, ok := mgr.connInfos[info.Name]; !ok {
			mgr.connInfos[info.Name] = make(map[foxproxy.ClientType][]*DEServer)
		}
		mgr.connInfos[info.Name][conn.ClientType] = append(mgr.connInfos[info.Name][conn.ClientType], conn)
	}
	mgr.connections = append(mgr.connections, conn)
	conn.WatchRecv(mgr.DealDataElement)
	conn.WatchClose(mgr.deleteConnection)
	return nil
}

func (mgr *DEServerMGR) GetClientInfos() []*foxproxy.ClientInfo {
	ret := []*foxproxy.ClientInfo{}
	for _, info := range mgr.connections {
		ret = append(ret, info.ClientInfo)
	}
	return ret
}

type MsgInfo struct {
	Payload  []byte
	ErrMsg   *string
	CoinInfo *foxproxy.CoinInfo
}

// delete conn from connectionMGR
func (mgr *DEServerMGR) deleteConnection(conn *DEServer) {
	for _, info := range conn.Infos {
		if _, ok := mgr.connInfos[info.Name]; !ok {
			continue
		}

		conns, ok := mgr.connInfos[info.Name][conn.ClientType]
		if !ok || len(conns) == 0 {
			continue
		}
		for i := 0; i < len(conns); i++ {
			idx := len(conns) - 1 - i
			if conns[idx].ID == conn.ID {
				conns = append(conns[:idx], conns[idx+1:]...)
			}
		}
		mgr.connInfos[info.Name][conn.ClientType] = conns
	}

	for i := 0; i < len(mgr.connections); i++ {
		idx := len(mgr.connections) - 1 - i
		if mgr.connections[idx].ID == conn.ID {
			mgr.connections = append(mgr.connections[:idx], mgr.connections[idx+1:]...)
		}
	}
}

func (mgr *DEServerMGR) CloseAll() {
	for _, conn := range mgr.connections {
		conn.Close()
	}
}

// if recvChannel is not nil, recv response will send to it
// default value of statusCode is success
func (mgr *DEServerMGR) SendMsg(
	name string,
	clientType foxproxy.ClientType,
	msgType foxproxy.MsgType,
	msgID *string,
	connID *string,
	msg *MsgInfo,
	recvChannel chan MsgInfo,
) error {
	if msg.CoinInfo == nil {
		coinInfo, err := regcoininfo.GetRegCoinInfo(context.Background(), name)
		if err != nil {
			return err
		}
		msg.CoinInfo = coinInfo
	}

	_name := name
	if clientType == foxproxy.ClientType_ClientTypeSign {
		_name = msg.CoinInfo.TempName
	}

	if _, ok := mgr.connInfos[_name]; !ok {
		return fmt.Errorf("cannot find any sider,for %v", name)
	}

	conns, ok := mgr.connInfos[_name][clientType]
	if !ok || len(conns) == 0 {
		return fmt.Errorf("cannot find any sider,for %v-%v", name, clientType)
	}

	var conn *DEServer
	if connID == nil {
		conn = conns[time.Now().Second()%len(conns)]
	} else {
		for _, _conn := range conns {
			if _conn.ID == *connID {
				conn = _conn
				break
			}
		}
		if conn == nil {
			return fmt.Errorf("cannot find any sider,for %v-%v-%v", name, clientType, connID)
		}
	}

	return mgr.sendMsg(msgType, msgID, msg, conn, recvChannel)
}

// if recvChannel is not nil, recv response will send to it
// default value of statusCode is success
func (mgr *DEServerMGR) SendMsgWithConnID(
	msgType foxproxy.MsgType,
	connID string,
	msgID *string,
	msg *MsgInfo,
	recvChannel chan MsgInfo,
) error {
	var conn *DEServer
	for _, _conn := range mgr.connections {
		if _conn.ID == connID {
			conn = _conn
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("cannot find any sider,for %v", connID)
	}

	return mgr.sendMsg(msgType, msgID, msg, conn, recvChannel)
}

// if recvChannel is not nil, recv response will send to it
// default value of statusCode is success
func (mgr *DEServerMGR) sendMsg(
	msgType foxproxy.MsgType,
	msgID *string,
	msg *MsgInfo,
	conn *DEServer,
	recvChannel chan MsgInfo,
) error {
	if conn == nil {
		return fmt.Errorf("connection is nil")
	}

	if msg == nil {
		msg = &MsgInfo{}
	}

	if msgID == nil {
		_msgID := uuid.NewString()
		msgID = &_msgID
	}
	if recvChannel != nil {
		mgr.recvChannel.Store(*msgID, recvChannel)
	}

	return conn.Send(&foxproxy.DataElement{
		ConnectID: conn.ID,
		MsgID:     *msgID,
		MsgType:   msgType,
		Payload:   msg.Payload,
		ErrMsg:    msg.ErrMsg,
		CoinInfo:  msg.CoinInfo,
	})
}

func (mgr *DEServerMGR) SendAndRecvRaw(ctx context.Context, name string, clientType foxproxy.ClientType, msgType foxproxy.MsgType, reqPayload []byte) ([]byte, error) {
	recvChannel := make(chan MsgInfo)
	defer close(recvChannel)

	err := mgr.SendMsg(name, clientType, msgType, nil, nil, &MsgInfo{Payload: reqPayload}, recvChannel)
	if err != nil {
		return nil, err
	}

	var recvMsg MsgInfo
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.NewTimer(time.Second * 3).C:
		return nil, fmt.Errorf("timeout for recv response")
	case recvMsg = <-recvChannel:
	}

	if recvMsg.ErrMsg != nil && *recvMsg.ErrMsg != "" {
		return nil, fmt.Errorf(*recvMsg.ErrMsg)
	}

	return recvMsg.Payload, nil
}

func (mgr *DEServerMGR) SendAndRecv(ctx context.Context, name string, clientType foxproxy.ClientType, msgType foxproxy.MsgType, req, resp interface{}) error {
	reqPayload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	respPayload, err := mgr.SendAndRecvRaw(ctx, name, clientType, msgType, reqPayload)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respPayload, resp)
	if err != nil {
		return err
	}

	return nil
}

func (mgr *DEServerMGR) DealDataElement(data *foxproxy.DataElement) {
	if ch, ok := mgr.recvChannel.LoadAndDelete(data.MsgID); ok {
		select {
		case <-time.NewTimer(time.Second).C:
		case ch.(chan MsgInfo) <- MsgInfo{
			Payload: data.Payload,
			ErrMsg:  data.ErrMsg,
		}:
		}
	}

	h, err := GetDEHandlerMGR().GetDEHandler(data.MsgType)
	if err != nil {
		logger.Sugar().Error(err)
		return
	}

	resp := h(context.Background(), data)
	if resp == nil {
		return
	}

	err = mgr.SendMsgWithConnID(foxproxy.MsgType_MsgTypeResponse, data.ConnectID, &data.MsgID, resp, nil)
	if err != nil {
		logger.Sugar().Error(err)
		return
	}
}
