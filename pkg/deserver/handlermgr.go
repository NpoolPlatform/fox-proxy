package deserver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/fox-proxy/pkg/utils"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

type DEHandlerFunc func(ctx context.Context, data *foxproxy.DataElement) *MsgInfo

type DEHandlerMGR struct {
	deHandlers map[foxproxy.MsgType]DEHandlerFunc
}

var hmgr *DEHandlerMGR

func GetDEHandlerMGR() *DEHandlerMGR {
	if hmgr == nil {
		hmgr = &DEHandlerMGR{
			deHandlers: make(map[foxproxy.MsgType]DEHandlerFunc),
		}
		hmgr.registerDefaultDEHandler()
	}
	return hmgr
}

func (mgr *DEHandlerMGR) registerDefaultDEHandler() {
	mgr.deHandlers[foxproxy.MsgType_MsgTypeEcho] = func(ctx context.Context, data *foxproxy.DataElement) *MsgInfo {
		return &MsgInfo{
			Payload:  data.Payload,
			ErrMsg:   data.ErrMsg,
			CoinInfo: data.CoinInfo,
		}
	}
	mgr.deHandlers[foxproxy.MsgType_MsgTypeResponse] = func(ctx context.Context, data *foxproxy.DataElement) *MsgInfo {
		return nil
	}
}

func (mgr *DEHandlerMGR) RegisterDEHandler(
	msgType foxproxy.MsgType,
	in interface{},
	handler func(ctx context.Context, req interface{}) (interface{}, error),
) {
	deHandler := func(ctx context.Context, data *foxproxy.DataElement) *MsgInfo {
		outPayload, err := func(data *foxproxy.DataElement, in interface{}) ([]byte, error) {
			inData := utils.Copy(in)
			err := json.Unmarshal(data.Payload, inData)
			if err != nil {
				return nil, err
			}

			out, err := handler(ctx, inData)
			if err != nil {
				return nil, err
			}

			outPayload, err := json.Marshal(out)
			if err != nil {
				return nil, err
			}
			return outPayload, err
		}(data, in)

		statusMsg := ""
		if err != nil {
			statusMsg = err.Error()
		}

		return &MsgInfo{
			Payload:  outPayload,
			ErrMsg:   &statusMsg,
			CoinInfo: data.CoinInfo,
		}
	}
	mgr.deHandlers[msgType] = deHandler
}

func (mgr *DEHandlerMGR) GetDEHandler(msgType foxproxy.MsgType) (DEHandlerFunc, error) {
	h, ok := mgr.deHandlers[msgType]
	if !ok {
		return nil, fmt.Errorf("have no handler for %v", msgType)
	}
	return h, nil
}
