// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/regcoininfo"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/schema"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/transaction"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	regcoininfoMixin := schema.RegCoinInfo{}.Mixin()
	regcoininfoMixinFields0 := regcoininfoMixin[0].Fields()
	_ = regcoininfoMixinFields0
	regcoininfoFields := schema.RegCoinInfo{}.Fields()
	_ = regcoininfoFields
	// regcoininfoDescEntID is the schema descriptor for ent_id field.
	regcoininfoDescEntID := regcoininfoMixinFields0[1].Descriptor()
	// regcoininfo.DefaultEntID holds the default value on creation for the ent_id field.
	regcoininfo.DefaultEntID = regcoininfoDescEntID.Default.(func() uuid.UUID)
	// regcoininfoDescChainType is the schema descriptor for chain_type field.
	regcoininfoDescChainType := regcoininfoFields[0].Descriptor()
	// regcoininfo.DefaultChainType holds the default value on creation for the chain_type field.
	regcoininfo.DefaultChainType = regcoininfoDescChainType.Default.(int32)
	// regcoininfoDescCoinType is the schema descriptor for coin_type field.
	regcoininfoDescCoinType := regcoininfoFields[1].Descriptor()
	// regcoininfo.DefaultCoinType holds the default value on creation for the coin_type field.
	regcoininfo.DefaultCoinType = regcoininfoDescCoinType.Default.(int32)
	// regcoininfoDescTempName is the schema descriptor for temp_name field.
	regcoininfoDescTempName := regcoininfoFields[2].Descriptor()
	// regcoininfo.DefaultTempName holds the default value on creation for the temp_name field.
	regcoininfo.DefaultTempName = regcoininfoDescTempName.Default.(string)
	// regcoininfoDescEnv is the schema descriptor for env field.
	regcoininfoDescEnv := regcoininfoFields[4].Descriptor()
	// regcoininfo.DefaultEnv holds the default value on creation for the env field.
	regcoininfo.DefaultEnv = regcoininfoDescEnv.Default.(string)
	// regcoininfoDescCreatedAt is the schema descriptor for created_at field.
	regcoininfoDescCreatedAt := regcoininfoFields[5].Descriptor()
	// regcoininfo.DefaultCreatedAt holds the default value on creation for the created_at field.
	regcoininfo.DefaultCreatedAt = regcoininfoDescCreatedAt.Default.(func() uint32)
	// regcoininfoDescUpdatedAt is the schema descriptor for updated_at field.
	regcoininfoDescUpdatedAt := regcoininfoFields[6].Descriptor()
	// regcoininfo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	regcoininfo.DefaultUpdatedAt = regcoininfoDescUpdatedAt.Default.(func() uint32)
	// regcoininfo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	regcoininfo.UpdateDefaultUpdatedAt = regcoininfoDescUpdatedAt.UpdateDefault.(func() uint32)
	// regcoininfoDescDeletedAt is the schema descriptor for deleted_at field.
	regcoininfoDescDeletedAt := regcoininfoFields[7].Descriptor()
	// regcoininfo.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	regcoininfo.DefaultDeletedAt = regcoininfoDescDeletedAt.Default.(func() uint32)
	transactionMixin := schema.Transaction{}.Mixin()
	transactionMixinFields0 := transactionMixin[0].Fields()
	_ = transactionMixinFields0
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescEntID is the schema descriptor for ent_id field.
	transactionDescEntID := transactionMixinFields0[1].Descriptor()
	// transaction.DefaultEntID holds the default value on creation for the ent_id field.
	transaction.DefaultEntID = transactionDescEntID.Default.(func() uuid.UUID)
	// transactionDescCoinType is the schema descriptor for coin_type field.
	transactionDescCoinType := transactionFields[0].Descriptor()
	// transaction.DefaultCoinType holds the default value on creation for the coin_type field.
	transaction.DefaultCoinType = transactionDescCoinType.Default.(int32)
	// transactionDescNonce is the schema descriptor for nonce field.
	transactionDescNonce := transactionFields[1].Descriptor()
	// transaction.DefaultNonce holds the default value on creation for the nonce field.
	transaction.DefaultNonce = transactionDescNonce.Default.(uint64)
	// transactionDescTransactionType is the schema descriptor for transaction_type field.
	transactionDescTransactionType := transactionFields[2].Descriptor()
	// transaction.DefaultTransactionType holds the default value on creation for the transaction_type field.
	transaction.DefaultTransactionType = transactionDescTransactionType.Default.(int8)
	// transactionDescRecentBhash is the schema descriptor for recent_bhash field.
	transactionDescRecentBhash := transactionFields[3].Descriptor()
	// transaction.DefaultRecentBhash holds the default value on creation for the recent_bhash field.
	transaction.DefaultRecentBhash = transactionDescRecentBhash.Default.(string)
	// transactionDescTxData is the schema descriptor for tx_data field.
	transactionDescTxData := transactionFields[4].Descriptor()
	// transaction.DefaultTxData holds the default value on creation for the tx_data field.
	transaction.DefaultTxData = transactionDescTxData.Default.([]byte)
	// transactionDescCid is the schema descriptor for cid field.
	transactionDescCid := transactionFields[6].Descriptor()
	// transaction.DefaultCid holds the default value on creation for the cid field.
	transaction.DefaultCid = transactionDescCid.Default.(string)
	// transactionDescExitCode is the schema descriptor for exit_code field.
	transactionDescExitCode := transactionFields[7].Descriptor()
	// transaction.DefaultExitCode holds the default value on creation for the exit_code field.
	transaction.DefaultExitCode = transactionDescExitCode.Default.(int64)
	// transactionDescName is the schema descriptor for name field.
	transactionDescName := transactionFields[8].Descriptor()
	// transaction.DefaultName holds the default value on creation for the name field.
	transaction.DefaultName = transactionDescName.Default.(string)
	// transactionDescFrom is the schema descriptor for from field.
	transactionDescFrom := transactionFields[9].Descriptor()
	// transaction.DefaultFrom holds the default value on creation for the from field.
	transaction.DefaultFrom = transactionDescFrom.Default.(string)
	// transactionDescTo is the schema descriptor for to field.
	transactionDescTo := transactionFields[10].Descriptor()
	// transaction.DefaultTo holds the default value on creation for the to field.
	transaction.DefaultTo = transactionDescTo.Default.(string)
	// transactionDescMemo is the schema descriptor for memo field.
	transactionDescMemo := transactionFields[11].Descriptor()
	// transaction.DefaultMemo holds the default value on creation for the memo field.
	transaction.DefaultMemo = transactionDescMemo.Default.(string)
	// transactionDescAmount is the schema descriptor for amount field.
	transactionDescAmount := transactionFields[12].Descriptor()
	// transaction.DefaultAmount holds the default value on creation for the amount field.
	transaction.DefaultAmount = transactionDescAmount.Default.(uint64)
	// transactionDescPayload is the schema descriptor for payload field.
	transactionDescPayload := transactionFields[13].Descriptor()
	// transaction.DefaultPayload holds the default value on creation for the payload field.
	transaction.DefaultPayload = transactionDescPayload.Default.([]byte)
	// transaction.PayloadValidator is a validator for the "payload" field. It is called by the builders before save.
	transaction.PayloadValidator = transactionDescPayload.Validators[0].(func([]byte) error)
	// transactionDescState is the schema descriptor for state field.
	transactionDescState := transactionFields[14].Descriptor()
	// transaction.DefaultState holds the default value on creation for the state field.
	transaction.DefaultState = transactionDescState.Default.(uint8)
	// transactionDescCreatedAt is the schema descriptor for created_at field.
	transactionDescCreatedAt := transactionFields[15].Descriptor()
	// transaction.DefaultCreatedAt holds the default value on creation for the created_at field.
	transaction.DefaultCreatedAt = transactionDescCreatedAt.Default.(func() uint32)
	// transactionDescUpdatedAt is the schema descriptor for updated_at field.
	transactionDescUpdatedAt := transactionFields[16].Descriptor()
	// transaction.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	transaction.DefaultUpdatedAt = transactionDescUpdatedAt.Default.(func() uint32)
	// transaction.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	transaction.UpdateDefaultUpdatedAt = transactionDescUpdatedAt.UpdateDefault.(func() uint32)
	// transactionDescDeletedAt is the schema descriptor for deleted_at field.
	transactionDescDeletedAt := transactionFields[17].Descriptor()
	// transaction.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	transaction.DefaultDeletedAt = transactionDescDeletedAt.Default.(func() uint32)
}
