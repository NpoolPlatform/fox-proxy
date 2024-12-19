package main

import "github.com/NpoolPlatform/message/npool/foxproxy"

type MsgClientType struct {
	MsgType    foxproxy.MsgType
	ClientType foxproxy.ClientType
}

type derNode struct {
	nextNodes      map[int]*derNode
	msgClientTypes *[]MsgClientType
}

type DERouter struct {
	router map[foxproxy.MsgType]derNode
}

func (r *DERouter) RegisterRouter(
	msgType foxproxy.MsgType,
	chainType foxproxy.ChainType,
	coinType foxproxy.CoinType,
)

func main() {

}
