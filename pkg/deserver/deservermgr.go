package deserver

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/foxproxy"
	"github.com/google/uuid"
)

//nolint:revive
type DEServerMGR struct {
	coinInfos   map[string]*foxproxy.CoinInfo
	connInfos   map[string]map[foxproxy.ClientType][]*DEServer
	recvChannel sync.Map
	connections []*DEServer
}

var cmgr *DEServerMGR

func GetDEServerMGR() *DEServerMGR {
	if cmgr == nil {
		cmgr = &DEServerMGR{
			coinInfos:   make(map[string]*foxproxy.CoinInfo),
			connInfos:   make(map[string]map[foxproxy.ClientType][]*DEServer),
			recvChannel: sync.Map{},
		}
	}
	return cmgr
}

func (mgr *DEServerMGR) AddDEServer(conn *DEServer) {
	for _, info := range conn.Infos {
		mgr.coinInfos[info.Name] = info
		if _, ok := mgr.connInfos[info.Name]; !ok {
			mgr.connInfos[info.Name] = make(map[foxproxy.ClientType][]*DEServer)
		}
		mgr.connInfos[info.Name][conn.ClientType] = append(mgr.connInfos[info.Name][conn.ClientType], conn)
	}
	mgr.connections = append(mgr.connections, conn)
	conn.WatchRecv(mgr.DealDataElement)
	conn.WatchClose(mgr.deleteConnection)
}

func (mgr *DEServerMGR) GetClientInfos() []*foxproxy.ClientInfo {
	ret := []*foxproxy.ClientInfo{}
	for _, info := range mgr.connections {
		ret = append(ret, info.ClientInfo)
	}
	return ret
}

type MsgInfo struct {
	Payload    []byte
	StatusCode *foxproxy.StatusCode
	StatusMsg  *string
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
	if _, ok := mgr.connInfos[name]; !ok {
		return fmt.Errorf("cannot find any sider,for %v", name)
	}

	conns, ok := mgr.connInfos[name][clientType]
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

	if msg.StatusCode == nil {
		msg.StatusCode = foxproxy.StatusCode_StatusCodeSuccess.Enum()
	}

	return conn.Send(&foxproxy.DataElement{
		ConnectID:  conn.ID,
		MsgID:      *msgID,
		MsgType:    msgType,
		Payload:    msg.Payload,
		StatusCode: *msg.StatusCode,
		StatusMsg:  msg.StatusMsg,
	})
}

func (mgr *DEServerMGR) SendAndRecv(ctx context.Context, name string, clientType foxproxy.ClientType, msgType foxproxy.MsgType, req, resp interface{}) (*foxproxy.StatusCode, error) {
	inPayload, err := json.Marshal(req)
	if err != nil {
		return foxproxy.StatusCode_StatusCodeMarshalErr.Enum(), err
	}

	recvChannel := make(chan MsgInfo)
	defer close(recvChannel)

	err = mgr.SendMsg(name, clientType, msgType, nil, nil, &MsgInfo{Payload: inPayload}, recvChannel)
	if err != nil {
		return foxproxy.StatusCode_StatusCodeFailed.Enum(), err
	}

	var recvMsg MsgInfo
	select {
	case <-ctx.Done():
		return foxproxy.StatusCode_StatusCodeFailed.Enum(), ctx.Err()
	case <-time.NewTimer(time.Second * 3).C:
		return foxproxy.StatusCode_StatusCodeFailed.Enum(), fmt.Errorf("timeout for recv response")
	case recvMsg = <-recvChannel:
	}

	if recvMsg.StatusCode.String() != foxproxy.StatusCode_StatusCodeSuccess.String() {
		if recvMsg.StatusMsg == nil {
			return recvMsg.StatusCode, fmt.Errorf("")
		}
		return recvMsg.StatusCode, fmt.Errorf(*recvMsg.StatusMsg)
	}

	err = json.Unmarshal(recvMsg.Payload, resp)
	if err != nil {
		return foxproxy.StatusCode_StatusCodeUnmarshalErr.Enum(), err
	}

	return foxproxy.StatusCode_StatusCodeSuccess.Enum(), nil
}

func (mgr *DEServerMGR) DealDataElement(data *foxproxy.DataElement) {
	if ch, ok := mgr.recvChannel.LoadAndDelete(data.MsgID); ok {
		select {
		case <-time.NewTimer(time.Second).C:
		case ch.(chan MsgInfo) <- MsgInfo{
			Payload:    data.Payload,
			StatusCode: &data.StatusCode,
			StatusMsg:  data.StatusMsg,
		}:
		}
	}

	handler, err := GetDEHandlerMGR().GetDEHandler(data.MsgType)
	if err != nil {
		logger.Sugar().Error(err)
		return
	}

	err = handler(data)
	if err != nil {
		logger.Sugar().Error(err)
	}
}

func (mgr *DEServerMGR) GetCoinInfo(name string) *foxproxy.CoinInfo {
	return mgr.coinInfos[name]
}
