package transaction

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/fox-proxy/pkg/db"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/fox-proxy/pkg/router"
	"github.com/NpoolPlatform/fox-proxy/pkg/utils"
	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/message/npool/foxproxy"
)

var assignLock sync.Mutex

// PullTransactions ..
func pullTxs(ctx context.Context, cli *ent.Tx, clientType foxproxy.ClientType, names []string) ([]*ent.Transaction, error) {
	querySql := fmt.Sprintf(
		"select * from transactions WHERE name IN (\"%v\") AND state NOT IN (%v) group by name,`from`;",
		strings.Join(names, "\",\""),
		strings.Join(utils.I32ToSliceString([]int32{
			int32(foxproxy.TransactionState_TransactionStateUnKnow),
			int32(foxproxy.TransactionState_TransactionStateDone),
			int32(foxproxy.TransactionState_TransactionStateFail),
			int32(foxproxy.TransactionState_TransactionStateExceedTime),
		}), ","),
	)

	ret := []*ent.Transaction{}
	_ret := []*ent.Transaction{}
	rows, err := cli.Transaction.QueryContext(ctx, querySql)
	if err != nil {
		return nil, err
	}

	err = sql.ScanSlice(rows, &_ret)
	if err != nil {
		return nil, err
	}

	for _, v := range _ret {
		if v.ClientType == int32(clientType) && v.LockTime == 0 {
			ret = append(ret, v)
		}
	}
	return ret, nil
}

func lockTxs(ctx context.Context, cli *ent.Tx, txs []*ent.Transaction) error {
	ids := []uint32{}
	now := uint32(time.Now().Unix())
	for _, v := range txs {
		ids = append(ids, v.ID)
		v.LockTime = now
	}

	_, err := cli.Transaction.
		Update().
		Where(
			transaction.IDIn(ids...),
			transaction.LockTimeEQ(0),
		).
		SetLockTime(now).
		Save(ctx)
	return err
}

func AssignTxs(ctx context.Context, clientType foxproxy.ClientType, names []string) ([]*foxproxy.Transaction, error) {
	assignLock.Lock()
	defer assignLock.Unlock()

	var ret []*foxproxy.Transaction
	err := db.WithTx(ctx, func(ctx context.Context, cli *ent.Tx) error {
		txs, err := pullTxs(ctx, cli, clientType, names)
		if err != nil {
			return err
		}

		err = lockTxs(ctx, cli, txs)
		if err != nil {
			return err
		}
		ret = EntsToProtos(txs)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// if exitcode < 0, state will be fail
// if exitcode = 0, state will be next state
// else exitcode > 0, state will not be changed
func SubmitTx(ctx context.Context, t *foxproxy.SubmitTransaction) error {
	tx, err := getTransaction(ctx, t.TransactionID)
	if err != nil {
		return wlog.WrapError(err)
	}
	currentStateStep := &router.TxStateStep{
		TxState:    foxproxy.TransactionState(tx.State),
		ClientType: foxproxy.ClientType(tx.ClientType),
	}

	if t.ExitCode == 0 {
		txStateSteps, err := router.
			GetTxStateRouter().
			GetTxStateSteps(
				(foxproxy.ChainType)(tx.ChainType).Enum(),
				(foxproxy.CoinType)(tx.CoinType).Enum(),
			)
		if err != nil {
			return wlog.WrapError(err)
		}
		currentStateStep, err = txStateSteps.GetNextStep((foxproxy.TransactionState)(tx.State).Enum())
		if err != nil {
			return wlog.WrapError(err)
		}
	}

	if t.ExitCode < 0 {
		currentStateStep = &router.TxStateStep{
			TxState:    foxproxy.TransactionState_TransactionStateFail,
			ClientType: foxproxy.ClientType_ClientTypeDefault,
		}
	}

	client, err := db.Client()
	if err != nil {
		return wlog.WrapError(err)
	}
	stmt := client.
		Transaction.
		Update().
		Where(
			transaction.TransactionIDEQ(t.TransactionID),
			transaction.StateEQ(int32(t.State)),
			transaction.LockTimeEQ(t.LockTime),
		).
		SetClientType(int32(currentStateStep.ClientType)).
		SetState(int32(currentStateStep.TxState)).
		SetLockTime(0).
		SetExitCode(t.ExitCode)

	if t.CID != nil {
		stmt.SetCid(*t.CID)
	}
	if t.Payload != nil {
		stmt.SetPayload(t.Payload)
	}

	return stmt.Exec(ctx)
}
