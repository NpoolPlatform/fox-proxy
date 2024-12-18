package deserver

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/NpoolPlatform/message/npool/foxproxy"
	"github.com/stretchr/testify/assert"
)

type A struct {
	Msg string
}

type B struct {
	Msg string
	num int
}

func TestDEHandlerMGR(t *testing.T) {
	a := &A{Msg: "sssssssssss"}
	test_hanlder := func(_ context.Context, in *A) (*B, *foxproxy.StatusCode, error) {
		assert.Equal(t, a.Msg, in.Msg)
		return &B{Msg: "cccc" + in.Msg, num: 122}, nil, nil
	}

	mgr := GetDEHandlerMGR()
	mgr.RegisterDEHandler(
		foxproxy.MsgType_MsgTypeUpdateTx,
		new(A),
		func(ctx context.Context, req interface{}) (interface{}, *foxproxy.StatusCode, error) {
			return test_hanlder(ctx, req.(*A))
		})
	payload, err := json.Marshal(a)
	assert.Nil(t, err)

	h, err := mgr.GetDEHandler(foxproxy.MsgType_MsgTypeUpdateTx)
	assert.Nil(t, err)

	h(&foxproxy.DataElement{Payload: payload})
}
