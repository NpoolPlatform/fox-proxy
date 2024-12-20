package deserver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/fox-proxy/pkg/utils"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

type DEHandlerFunc func(data *foxproxy.DataElement) error

type DEHandler struct {
	MsgType foxproxy.MsgType
	Handler DEHandlerFunc
}

type DEHandlerMGR struct {
	DEHandlers map[foxproxy.MsgType]*DEHandler
}

var hmgr *DEHandlerMGR

func GetDEHandlerMGR() *DEHandlerMGR {
	if hmgr == nil {
		hmgr = &DEHandlerMGR{
			DEHandlers: make(map[foxproxy.MsgType]*DEHandler),
		}
		hmgr.registerDefaultDEHandler()
	}
	return hmgr
}

func (mgr *DEHandlerMGR) registerDefaultDEHandler() {
	echoHandler := DEHandler{
		MsgType: foxproxy.MsgType_MsgTypeEcho,
		Handler: func(data *foxproxy.DataElement) error {
			return GetDEServerMGR().SendMsgWithConnID(
				data.MsgType,
				data.ConnectID,
				&data.MsgID,
				&MsgInfo{
					Payload:    data.Payload,
					StatusCode: &data.StatusCode,
					StatusMsg:  data.StatusMsg,
				},
				nil,
			)
		},
	}
	mgr.DEHandlers[echoHandler.MsgType] = &echoHandler
}

func (mgr *DEHandlerMGR) RegisterDEHandler(
	msgType foxproxy.MsgType,
	in interface{},
	handler func(ctx context.Context, req interface{}) (interface{}, *foxproxy.StatusCode, error),
) {
	deHandler := DEHandler{
		MsgType: msgType,
		Handler: func(data *foxproxy.DataElement) error {
			outPayload, statusCode, err := func(data *foxproxy.DataElement, in interface{}) ([]byte, *foxproxy.StatusCode, error) {
				inData := utils.Copy(in)
				err := json.Unmarshal(data.Payload, inData)
				if err != nil {
					return nil, foxproxy.StatusCode_StatusCodeUnmarshalErr.Enum(), err
				}

				out, statusCode, err := handler(context.Background(), inData)
				if err != nil {
					if statusCode == nil {
						statusCode = foxproxy.StatusCode_StatusCodeFailed.Enum()
					}
					return nil, statusCode, err
				}

				outPayload, err := json.Marshal(out)
				if err != nil {
					return nil, foxproxy.StatusCode_StatusCodeMarshalErr.Enum(), err
				}
				return outPayload, foxproxy.StatusCode_StatusCodeSuccess.Enum(), err
			}(data, in)

			statusMsg := ""
			if err != nil {
				statusMsg = err.Error()
			}

			return GetDEServerMGR().SendMsgWithConnID(
				data.MsgType,
				data.ConnectID,
				&data.MsgID,
				&MsgInfo{
					Payload:    outPayload,
					StatusCode: statusCode,
					StatusMsg:  &statusMsg,
				},
				nil,
			)
		},
	}
	mgr.DEHandlers[deHandler.MsgType] = &deHandler
}

func (mgr *DEHandlerMGR) GetDEHandler(msgType foxproxy.MsgType) (DEHandlerFunc, error) {
	h, ok := mgr.DEHandlers[msgType]
	if !ok {
		return nil, fmt.Errorf("have no handler for %v", msgType)
	}
	return h.Handler, nil
}
