package router

import (
	"fmt"

	"github.com/NpoolPlatform/message/npool/foxproxy"
)

type TxStateStep struct {
	TxState    foxproxy.TransactionState
	ClientType foxproxy.ClientType
}

type TxStateSteps []*TxStateStep

type TxStateRouter struct{ treeNode }

func (r *TxStateRouter) RegisterRouter(
	val TxStateSteps,
	chainType *foxproxy.ChainType,
	coinType *foxproxy.CoinType,
) error {
	pathList, err := r.getPathList(chainType, coinType)
	if err != nil {
		return err
	}

	r.registerRouter(val, pathList...)
	return nil
}

func (r *TxStateRouter) GetTxStateSteps(
	chainType *foxproxy.ChainType,
	coinType *foxproxy.CoinType,
) (TxStateSteps, error) {
	pathList, err := r.getPathList(chainType, coinType)
	if err != nil {
		return nil, err
	}
	val, _, err := r.getVal(pathList...)
	if err != nil {
		return nil, err
	}
	return val.(TxStateSteps), nil
}

func (r *TxStateRouter) getPathList(
	chainType *foxproxy.ChainType,
	coinType *foxproxy.CoinType,
) ([]int, error) {
	if coinType != nil && chainType == nil {
		return nil, fmt.Errorf("chaintype cannot be nil,when cointype is not nil")
	}
	pathList := []int{}
	if chainType != nil {
		pathList = append(pathList, int(*chainType))
	}
	if coinType != nil {
		pathList = append(pathList, int(*coinType))
	}
	return pathList, nil
}

// get next txStateStep
// will return nil, if txState is last state
func (tss TxStateSteps) GetNextStep(txState *foxproxy.TransactionState) (*TxStateStep, error) {
	if len(tss) == 0 {
		return nil, fmt.Errorf("cannot find any tx state step")
	}
	idx := 0
	for ; idx < len(tss); idx++ {
		if tss[idx].TxState == *txState {
			if idx == len(tss)-1 {
				return nil, nil
			}
			return tss[idx+1], nil
		}
	}
	return nil, fmt.Errorf("cannot find next state")
}
