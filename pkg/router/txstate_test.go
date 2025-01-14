package router

import (
	"testing"

	"github.com/NpoolPlatform/message/npool/foxproxy"
	"github.com/stretchr/testify/assert"
)

func TestTxStateRouter(t *testing.T) {
	txR := &TxStateRouter{}
	txSteops1 := TxStateSteps{
		{TxState: foxproxy.TransactionState_TransactionStatePrepare, ClientType: foxproxy.ClientType_ClientTypePlugin},
		{TxState: foxproxy.TransactionState_TransactionStateBroadcast, ClientType: foxproxy.ClientType_ClientTypePlugin},
	}
	err := txR.RegisterRouter(txSteops1, foxproxy.ChainType_Aleo.Enum(), foxproxy.CoinType_CoinTypealeo.Enum())
	assert.Nil(t, err)

	ret, err := txR.GetTxStateSteps(foxproxy.ChainType_Binancecoin.Enum(), nil)
	assert.NotNil(t, err)
	assert.Nil(t, ret)

	ret, err = txR.GetTxStateSteps(nil, nil)
	assert.NotNil(t, err)
	assert.Nil(t, ret)

	ret, err = txR.GetTxStateSteps(foxproxy.ChainType_Aleo.Enum(), nil)
	assert.NotNil(t, err)
	assert.Nil(t, ret)

	ret, err = txR.GetTxStateSteps(foxproxy.ChainType_Aleo.Enum(), foxproxy.CoinType_CoinTypebinancecoin.Enum())
	assert.NotNil(t, err)
	assert.Nil(t, ret)

	ret, err = txR.GetTxStateSteps(foxproxy.ChainType_Aleo.Enum(), foxproxy.CoinType_CoinTypealeo.Enum())
	assert.Nil(t, err)
	assert.Equal(t, txSteops1, ret)

	txSteops2 := TxStateSteps{
		{TxState: foxproxy.TransactionState_TransactionStatePrepare, ClientType: foxproxy.ClientType_ClientTypeSign},
		{TxState: foxproxy.TransactionState_TransactionStateSign, ClientType: foxproxy.ClientType_ClientTypePlugin},
	}

	err = txR.RegisterRouter(txSteops2, foxproxy.ChainType_Depinc.Enum(), nil)
	assert.Nil(t, err)

	ret, err = txR.GetTxStateSteps(foxproxy.ChainType_Binancecoin.Enum(), nil)
	assert.NotNil(t, err)
	assert.Nil(t, ret)

	ret, err = txR.GetTxStateSteps(nil, nil)
	assert.NotNil(t, err)
	assert.Nil(t, ret)

	ret, err = txR.GetTxStateSteps(foxproxy.ChainType_Depinc.Enum(), nil)
	assert.Nil(t, err)
	assert.Equal(t, txSteops2, ret)

	ret, err = txR.GetTxStateSteps(foxproxy.ChainType_Depinc.Enum(), foxproxy.CoinType_CoinTypebinancecoin.Enum())
	assert.Nil(t, err)
	assert.Equal(t, txSteops2, ret)

	txSteops3 := TxStateSteps{
		{TxState: foxproxy.TransactionState_TransactionStatePrepare, ClientType: foxproxy.ClientType_ClientTypeSign},
		{TxState: foxproxy.TransactionState_TransactionStateSign, ClientType: foxproxy.ClientType_ClientTypePlugin},
		{TxState: foxproxy.TransactionState_TransactionStateBroadcast, ClientType: foxproxy.ClientType_ClientTypePlugin},
		{TxState: foxproxy.TransactionState_TransactionStateSync, ClientType: foxproxy.ClientType_ClientTypePlugin},
	}

	err = txR.RegisterRouter(txSteops3, nil, nil)
	assert.Nil(t, err)

	ret, err = txR.GetTxStateSteps(foxproxy.ChainType_Binancecoin.Enum(), nil)
	assert.Nil(t, err)
	assert.Equal(t, txSteops3, ret)

	ret, err = txR.GetTxStateSteps(nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, txSteops3, ret)

	ret, err = txR.GetTxStateSteps(foxproxy.ChainType_Binancecoin.Enum(), foxproxy.CoinType_CoinTypebinancecoin.Enum())
	assert.Nil(t, err)
	assert.Equal(t, txSteops3, ret)

	txStateStep, err := ret.GetNextStep(foxproxy.TransactionState_TransactionStateGetViewKey.Enum())
	assert.NotNil(t, err)
	assert.Nil(t, txStateStep)

	txStateStep, err = ret.GetNextStep(foxproxy.TransactionState_TransactionStatePrepare.Enum())
	assert.Nil(t, err)
	assert.Equal(t, *txSteops3[1], *txStateStep)

	txStateStep, err = ret.GetNextStep(foxproxy.TransactionState_TransactionStateSign.Enum())
	assert.Nil(t, err)
	assert.Equal(t, *txSteops3[2], *txStateStep)

	txStateStep, err = ret.GetNextStep(foxproxy.TransactionState_TransactionStateSync.Enum())
	assert.Nil(t, err)
	assert.Nil(t, txStateStep)
}
