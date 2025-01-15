package router

import (
	"testing"

	"github.com/NpoolPlatform/message/npool/foxproxy"
	"github.com/stretchr/testify/assert"
)

func TestMsgTypeRouter(t *testing.T) {
	mtR := &MsgTypeRouter{}
	steps1 := MsgTypeSteps{
		{MsgType: foxproxy.MsgType_MsgTypeDefault, ClientType: foxproxy.ClientType_ClientTypePlugin},
		{MsgType: foxproxy.MsgType_MsgTypeEcho, ClientType: foxproxy.ClientType_ClientTypeSign},
		{MsgType: foxproxy.MsgType_MsgTypeSubmitTx, ClientType: foxproxy.ClientType_ClientTypePlugin},
	}
	mtR.RegisterRouter(
		steps1,
		foxproxy.MsgType_MsgTypeGetBalance,
		nil,
		nil,
	)
	ret, err := mtR.GetMsgTypeSteps(foxproxy.MsgType_MsgTypeGetBalance, nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, steps1, ret)

	ret, err = mtR.GetMsgTypeSteps(foxproxy.MsgType_MsgTypeGetBalance, foxproxy.ChainType_Aleo.Enum(), nil)
	assert.Nil(t, err)
	assert.Equal(t, steps1, ret)

	ret, err = mtR.GetMsgTypeSteps(foxproxy.MsgType_MsgTypeGetBalance, foxproxy.ChainType_Aleo.Enum(), foxproxy.CoinType_CoinTypealeo.Enum())
	assert.Nil(t, err)
	assert.Equal(t, steps1, ret)

	ret, err = mtR.GetMsgTypeSteps(foxproxy.MsgType_MsgTypeGetBalance, nil, foxproxy.CoinType_CoinTypealeo.Enum())
	assert.NotNil(t, err)
	assert.Nil(t, ret)
	ret, err = mtR.GetMsgTypeSteps(foxproxy.MsgType_MsgTypeEcho, nil, nil)
	assert.NotNil(t, err)
	assert.Nil(t, ret)

	steps2 := MsgTypeSteps{
		{MsgType: foxproxy.MsgType_MsgTypeEcho, ClientType: foxproxy.ClientType_ClientTypeSign},
		{MsgType: foxproxy.MsgType_MsgTypeSubmitTx, ClientType: foxproxy.ClientType_ClientTypePlugin},
	}
	mtR.RegisterRouter(
		steps2,
		foxproxy.MsgType_MsgTypeEcho,
		foxproxy.ChainType_Aleo.Enum(),
		nil,
	)

	ret, err = mtR.GetMsgTypeSteps(foxproxy.MsgType_MsgTypeEcho, nil, nil)
	assert.NotNil(t, err)
	assert.Nil(t, ret)

	ret, err = mtR.GetMsgTypeSteps(foxproxy.MsgType_MsgTypeEcho, foxproxy.ChainType_Aleo.Enum(), nil)
	assert.Nil(t, err)
	assert.Equal(t, steps2, ret)

	ret, err = mtR.GetMsgTypeSteps(foxproxy.MsgType_MsgTypeEcho, foxproxy.ChainType_Aleo.Enum(), foxproxy.CoinType_CoinTypealeo.Enum())
	assert.Nil(t, err)
	assert.Equal(t, steps2, ret)

	ret, err = mtR.GetMsgTypeSteps(foxproxy.MsgType_MsgTypeEcho, foxproxy.ChainType_Filecoin.Enum(), nil)
	assert.NotNil(t, err)
	assert.Nil(t, ret)

	ret, err = mtR.GetMsgTypeSteps(foxproxy.MsgType_MsgTypeEcho, foxproxy.ChainType_Filecoin.Enum(), foxproxy.CoinType_CoinTypefilecoin.Enum())
	assert.NotNil(t, err)
	assert.Nil(t, ret)
}
