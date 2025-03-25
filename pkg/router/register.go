package router

import "github.com/NpoolPlatform/message/npool/foxproxy"

func init() {
	mtRouter := GetMsgTypeRouter()
	// default get balance steps
	mtRouter.RegisterRouter(MsgTypeSteps{
		{foxproxy.MsgType_MsgTypeGetBalance, foxproxy.ClientType_ClientTypePlugin},
	}, foxproxy.MsgType_MsgTypeGetBalance, nil, nil)

	mtRouter.RegisterRouter(MsgTypeSteps{
		{foxproxy.MsgType_MsgTypeGetViewKey, foxproxy.ClientType_ClientTypeSign},
		{foxproxy.MsgType_MsgTypeGetBalance, foxproxy.ClientType_ClientTypePlugin},
	}, foxproxy.MsgType_MsgTypeGetBalance, foxproxy.ChainType_Aleo.Enum(), nil)

	mtRouter.RegisterRouter(MsgTypeSteps{
		{foxproxy.MsgType_MsgTypeCreateWallet, foxproxy.ClientType_ClientTypeSign},
	}, foxproxy.MsgType_MsgTypeCreateWallet, nil, nil)

	mtRouter.RegisterRouter(MsgTypeSteps{
		{foxproxy.MsgType_MsgTypeGetEstimateGas, foxproxy.ClientType_ClientTypePlugin},
	}, foxproxy.MsgType_MsgTypeGetEstimateGas, nil, nil)

	txsRouter := GetTxStateRouter()
	err := txsRouter.RegisterRouter(TxStateSteps{
		{foxproxy.TransactionState_TransactionStatePrepare, foxproxy.ClientType_ClientTypePlugin},
		{foxproxy.TransactionState_TransactionStateSign, foxproxy.ClientType_ClientTypeSign},
		{foxproxy.TransactionState_TransactionStateBroadcast, foxproxy.ClientType_ClientTypePlugin},
		{foxproxy.TransactionState_TransactionStateSync, foxproxy.ClientType_ClientTypePlugin},
	}, nil, nil)
	if err != nil {
		panic(err)
	}

	err = txsRouter.RegisterRouter(TxStateSteps{
		{foxproxy.TransactionState_TransactionStateGetViewKey, foxproxy.ClientType_ClientTypeSign},
		{foxproxy.TransactionState_TransactionStatePrepare, foxproxy.ClientType_ClientTypePlugin},
		{foxproxy.TransactionState_TransactionStateSign, foxproxy.ClientType_ClientTypeSign},
		{foxproxy.TransactionState_TransactionStateBroadcast, foxproxy.ClientType_ClientTypePlugin},
		{foxproxy.TransactionState_TransactionStateSync, foxproxy.ClientType_ClientTypePlugin},
	}, foxproxy.ChainType_Aleo.Enum(), nil)
	if err != nil {
		panic(err)
	}
}
