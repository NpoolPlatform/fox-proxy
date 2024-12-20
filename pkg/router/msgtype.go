package router

import (
	"fmt"

	"github.com/NpoolPlatform/message/npool/foxproxy"
)

type MsgTypeSteps []struct {
	MsgType    foxproxy.MsgType
	ClientType foxproxy.ClientType
}
type MsgTypeRouter struct{ treeNode }

func (r *MsgTypeRouter) RegisterRouter(
	val MsgTypeSteps,
	msgType foxproxy.MsgType,
	chainType *foxproxy.ChainType,
	coinType *foxproxy.CoinType,
) error {
	pathList, err := r.getPathList(msgType, chainType, coinType)
	if err != nil {
		return err
	}

	r.registerRouter(val, pathList...)
	return nil
}

func (r *MsgTypeRouter) GetMsgTypeSteps(
	msgType foxproxy.MsgType,
	chainType *foxproxy.ChainType,
	coinType *foxproxy.CoinType,
) (MsgTypeSteps, error) {
	pathList, err := r.getPathList(msgType, chainType, coinType)
	if err != nil {
		return nil, err
	}
	val, _, err := r.getVal(pathList...)
	if err != nil {
		return nil, err
	}
	return val.(MsgTypeSteps), nil
}

func (r *MsgTypeRouter) getPathList(
	msgType foxproxy.MsgType,
	chainType *foxproxy.ChainType,
	coinType *foxproxy.CoinType,
) ([]int, error) {
	if coinType != nil && chainType == nil {
		return nil, fmt.Errorf("chaintype cannot be nil,when cointype is not nil")
	}
	pathList := []int{int(msgType)}
	if chainType != nil {
		pathList = append(pathList, int(*chainType))
	}
	if coinType != nil {
		pathList = append(pathList, int(*coinType))
	}
	return pathList, nil
}
