package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/NpoolPlatform/fox-proxy/pkg/deserver"
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
	testHandler := func(_ context.Context, in *A) (*B, *foxproxy.StatusCode, error) {
		assert.Equal(t, a.Msg, in.Msg)
		return &B{Msg: "cccc" + in.Msg, num: 122}, foxproxy.StatusCode_StatusCodeSuccess.Enum(), fmt.Errorf("")
	}

	mgr := deserver.GetDEHandlerMGR()
	mgr.RegisterDEHandler(
		foxproxy.MsgType_MsgTypeUpdateTx,
		new(A),
		func(ctx context.Context, req interface{}) (interface{}, *foxproxy.StatusCode, error) {
			return testHandler(ctx, req.(*A))
		})
	payload, err := json.Marshal(a)
	assert.Nil(t, err)

	h, err := mgr.GetDEHandler(foxproxy.MsgType_MsgTypeUpdateTx)
	assert.Nil(t, err)

	err = h(&foxproxy.DataElement{Payload: payload})
	assert.NotNil(t, err)
}
