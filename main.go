package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/NpoolPlatform/fox-proxy/pkg/crud/transaction"
	"github.com/NpoolPlatform/fox-proxy/pkg/db"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/foxproxy"
	_ "github.com/go-sql-driver/mysql" // nolint
	"github.com/google/uuid"

	"github.com/NpoolPlatform/fox-proxy/pkg/servicename"
	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	mysqlconst "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	redisconst "github.com/NpoolPlatform/go-service-framework/pkg/redis/const"
)

var (
	names  = []string{"A", "B", "C", "D", "E", "F", "G"}
	froms  = []string{"11", "22", "33", "44", "55", "66", "77", "88", "99"}
	states = []foxproxy.TransactionState{
		foxproxy.TransactionState_TransactionStateUnKnow,
		foxproxy.TransactionState_TransactionStatePrepare,
		foxproxy.TransactionState_TransactionStateGetViewKey,
		foxproxy.TransactionState_TransactionStateSign,
		foxproxy.TransactionState_TransactionStateBroadcast,
		foxproxy.TransactionState_TransactionStateSync,
		foxproxy.TransactionState_TransactionStateDone,
		foxproxy.TransactionState_TransactionStateFail,
	}
)

const serviceName = servicename.ServiceName

func main() {
	logger.Init(logger.DebugLevel, "./a.log")
	description := fmt.Sprintf("my %v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		serviceName, serviceName)
	err := app.Init(
		serviceName,
		description,
		"",
		"",
		"./",
		nil,
		nil,
		mysqlconst.MysqlServiceName,
		redisconst.RedisServiceName,
	)
	fmt.Println(err)
	err = db.Init()
	fmt.Println(err)
	// PrepareTestData()

	// ret, err := transaction.AssginTxs(context.Background(), foxproxy.ClientType_ClientTypePlugin, []string{"A", "B"})
	// fmt.Println(err)
	// fmt.Println(utils.PrettyStruct(ret))
	err = transaction.SubmitTx(context.Background(), &foxproxy.SubmitTransaction{
		TransactionID: "85009d48-880d-40ed-ad75-a056d288b926",
		LockTime:      1736765867,
		State:         foxproxy.TransactionState_TransactionStateSync,
		Payload:       []byte("hello lll"),
		ExitCode:      -1,
	})
	fmt.Println(err)
}

func PrepareTestData() {
	for i := 0; i < 1000000; i++ {
		err := transaction.CreateTransaction(context.Background(),
			&transaction.CreateTransactionParam{
				CoinType:      foxproxy.CoinType_CoinTypealeo,
				ChainType:     foxproxy.ChainType_Aleo,
				ClientType:    foxproxy.ClientType_ClientTypePlugin,
				State:         states[rand.Int()%len(states)],
				TransactionID: uuid.NewString(),
				Name:          names[rand.Int()%len(names)],
				From:          froms[rand.Int()%len(froms)],
				To:            froms[rand.Int()%len(froms)],
				Value:         0,
				Memo:          "ssss",
			})
		fmt.Println(err, i)
	}
}
