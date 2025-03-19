package api

import (
	"context"

	coincli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	"github.com/NpoolPlatform/fox-proxy/pkg/crud/regcoininfo"
	crud "github.com/NpoolPlatform/fox-proxy/pkg/crud/transaction"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/fox-proxy/pkg/router"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	"github.com/NpoolPlatform/message/npool/foxproxy"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTransaction(ctx context.Context, in *foxproxy.CreateTransactionRequest) (out *foxproxy.CreateTransactionResponse, err error) {
	out = &foxproxy.CreateTransactionResponse{}

	// args check
	if in.GetName() == "" {
		logger.Sugar().Errorf("CreateTransaction Name: %v empty", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name empty")
	}

	// query coininfo
	coinExist, err := coincli.GetCoinOnly(ctx, &coinpb.Conds{
		Name: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: in.GetName(),
		},
	})
	if err != nil {
		logger.Sugar().Errorf("check coin info %v error %v", in.GetName(), err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	if coinExist == nil {
		logger.Sugar().Errorf("check coin info %v not exist", in.GetName())
		return out, status.Errorf(codes.NotFound, "coin %v not found", in.GetName())
	}

	if in.GetTransactionID() == "" {
		logger.Sugar().Errorf("CreateTransaction TransactionID: %v invalid", in.GetTransactionID())
		return out, status.Error(codes.InvalidArgument, "TransactionID Invalid")
	}

	if in.GetFrom() == "" {
		logger.Sugar().Errorf("CreateTransaction From: %v invalid", in.GetFrom())
		return out, status.Error(codes.InvalidArgument, "From Invalid")
	}

	if in.GetTo() == "" {
		logger.Sugar().Errorf("CreateTransaction To: %v invalid", in.GetTo())
		return out, status.Error(codes.InvalidArgument, "To Invalid")
	}

	if in.GetAmount() <= 0 {
		logger.Sugar().Errorf("CreateTransaction Amount: %v invalid", in.GetAmount())
		return out, status.Error(codes.InvalidArgument, "Amount Invalid")
	}

	exist, err := crud.GetTransactionExist(ctx, crud.GetTransactionExistParam{TransactionID: in.GetTransactionID()})
	if err != nil {
		logger.Sugar().Errorf("CreateTransaction cal GetTransactionExist error: %v", err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	if exist {
		logger.Sugar().Errorf("CreateTransaction TransactionID: %v already exist", in.GetTransactionID())
		return out, status.Errorf(codes.AlreadyExists, "TransactionID: %v already exist", in.GetTransactionID())
	}

	coinInfo, err := regcoininfo.GetRegCoinInfo(ctx, in.Name)
	if err != nil {
		logger.Sugar().Error(err)
		return out, status.Error(codes.InvalidArgument, err.Error())
	}

	txStateSteps, err := router.GetTxStateRouter().GetTxStateSteps(&coinInfo.ChainType, &coinInfo.CoinType)
	if err != nil {
		logger.Sugar().Error(err)
		return out, status.Error(codes.InvalidArgument, err.Error())
	}

	nextState, err := txStateSteps.GetNextStep(nil)
	if err != nil {
		logger.Sugar().Error(err)
		return out, status.Error(codes.InvalidArgument, err.Error())
	}

	// store to db
	if err := crud.CreateTransaction(ctx, &crud.CreateTransactionParam{
		CoinType:      coinInfo.CoinType,
		ChainType:     coinInfo.ChainType,
		State:         nextState.TxState,
		ClientType:    nextState.ClientType,
		TransactionID: in.GetTransactionID(),
		Name:          in.GetName(),
		From:          in.GetFrom(),
		To:            in.GetTo(),
		Value:         in.GetAmount(),
		Memo:          in.GetMemo(),
	}); err != nil {
		logger.Sugar().Errorf("CreateTransaction save to db error: %v,Transaction:%v", err, in)
		return out, status.Error(codes.Internal, "internal server error")
	}

	return out, nil
}

func (s *Server) GetTransaction(ctx context.Context, in *foxproxy.GetTransactionRequest) (out *foxproxy.GetTransactionResponse, err error) {
	if in.GetTransactionID() == "" {
		logger.Sugar().Errorf("GetTransaction TransactionID empty")
		return &foxproxy.GetTransactionResponse{}, status.Error(codes.InvalidArgument, "TransactionID empty")
	}

	info, err := crud.GetTransaction(ctx, in.GetTransactionID())
	if ent.IsNotFound(err) {
		logger.Sugar().Errorf("GetTransaction TransactionID: %v not found", in.GetTransactionID())
		return &foxproxy.GetTransactionResponse{}, status.Errorf(codes.NotFound, "TransactionID: %v not found", in.GetTransactionID())
	}

	if err != nil {
		logger.Sugar().Errorf("GetTransaction call GetTransaction error: %v", err)
		return &foxproxy.GetTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &foxproxy.GetTransactionResponse{
		Info: info,
	}, nil
}

// GetTransactions ..
func (s *Server) GetTransactions(ctx context.Context, in *foxproxy.GetTransactionsRequest) (out *foxproxy.GetTransactionsResponse, err error) {
	transInfos, total, err := crud.GetTransactions(ctx, in)
	if ent.IsNotFound(err) {
		logger.Sugar().Info("GetTransactions no wait transaction")
		return &foxproxy.GetTransactionsResponse{}, nil
	}

	if err != nil {
		logger.Sugar().Errorf("GetTransactions call GetTransactions error: %v", err)
		return &foxproxy.GetTransactionsResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &foxproxy.GetTransactionsResponse{
		Infos: transInfos,
		Total: total,
	}, nil
}

func (s *Server) UpdateTransaction(ctx context.Context, in *foxproxy.UpdateTransactionRequest) (out *foxproxy.UpdateTransactionResponse, err error) {
	if in.GetTransactionID() == "" {
		logger.Sugar().Info("UpdateTransaction TransactionID empty")
		return &foxproxy.UpdateTransactionResponse{},
			status.Error(codes.InvalidArgument, "TransactionID empty")
	}

	if in.GetState() == foxproxy.TransactionState_TransactionStateUnKnow {
		logger.Sugar().Errorw(
			"GetTransactions no wait transaction",
			"TransactionID", in.GetTransactionID(),
		)
		return &foxproxy.UpdateTransactionResponse{},
			status.Error(codes.InvalidArgument, "TransactionState empty")
	}

	exist, err := crud.GetTransactionExist(ctx, crud.GetTransactionExistParam{
		TransactionID: in.GetTransactionID(),
	})
	if err != nil {
		logger.Sugar().Errorw(
			"GetTransactionExist",
			"TransactionID", in.GetTransactionID(),
			"Error", err,
		)
		return &foxproxy.UpdateTransactionResponse{},
			status.Error(codes.Internal, "internal server error")
	}

	if !exist {
		return &foxproxy.UpdateTransactionResponse{},
			status.Errorf(codes.NotFound, "TransactionID: %v and State: %v not found",
				in.GetTransactionID(),
				in.GetState(),
			)
	}

	err = crud.UpdateTransaction(ctx, &crud.UpdateTransactionParams{
		TransactionID: in.GetTransactionID(),
		State:         in.GetState(),
		Payload:       in.GetPayload(),
		Cid:           in.GetCID(),
		ExitCode:      in.GetExitCode(),
	})
	if err != nil {
		logger.Sugar().Errorf("UpdateTransaction call error: %v", err)
		return &foxproxy.UpdateTransactionResponse{},
			status.Error(codes.Internal, "internal server error")
	}

	return &foxproxy.UpdateTransactionResponse{}, nil
}
