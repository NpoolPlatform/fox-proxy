// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/regcoininfo"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/transaction"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   regcoininfo.Table,
			Columns: regcoininfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: regcoininfo.FieldID,
			},
		},
		Type: "RegCoinInfo",
		Fields: map[string]*sqlgraph.FieldSpec{
			regcoininfo.FieldEntID:     {Type: field.TypeUUID, Column: regcoininfo.FieldEntID},
			regcoininfo.FieldChainType: {Type: field.TypeInt32, Column: regcoininfo.FieldChainType},
			regcoininfo.FieldCoinType:  {Type: field.TypeInt32, Column: regcoininfo.FieldCoinType},
			regcoininfo.FieldTempName:  {Type: field.TypeString, Column: regcoininfo.FieldTempName},
			regcoininfo.FieldName:      {Type: field.TypeString, Column: regcoininfo.FieldName},
			regcoininfo.FieldEnv:       {Type: field.TypeString, Column: regcoininfo.FieldEnv},
			regcoininfo.FieldCreatedAt: {Type: field.TypeUint32, Column: regcoininfo.FieldCreatedAt},
			regcoininfo.FieldUpdatedAt: {Type: field.TypeUint32, Column: regcoininfo.FieldUpdatedAt},
			regcoininfo.FieldDeletedAt: {Type: field.TypeUint32, Column: regcoininfo.FieldDeletedAt},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   transaction.Table,
			Columns: transaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: transaction.FieldID,
			},
		},
		Type: "Transaction",
		Fields: map[string]*sqlgraph.FieldSpec{
			transaction.FieldEntID:         {Type: field.TypeUUID, Column: transaction.FieldEntID},
			transaction.FieldCoinType:      {Type: field.TypeInt32, Column: transaction.FieldCoinType},
			transaction.FieldChainType:     {Type: field.TypeInt32, Column: transaction.FieldChainType},
			transaction.FieldClientType:    {Type: field.TypeInt32, Column: transaction.FieldClientType},
			transaction.FieldTransactionID: {Type: field.TypeString, Column: transaction.FieldTransactionID},
			transaction.FieldCid:           {Type: field.TypeString, Column: transaction.FieldCid},
			transaction.FieldExitCode:      {Type: field.TypeInt64, Column: transaction.FieldExitCode},
			transaction.FieldName:          {Type: field.TypeString, Column: transaction.FieldName},
			transaction.FieldFrom:          {Type: field.TypeString, Column: transaction.FieldFrom},
			transaction.FieldTo:            {Type: field.TypeString, Column: transaction.FieldTo},
			transaction.FieldMemo:          {Type: field.TypeString, Column: transaction.FieldMemo},
			transaction.FieldAmount:        {Type: field.TypeUint64, Column: transaction.FieldAmount},
			transaction.FieldPayload:       {Type: field.TypeBytes, Column: transaction.FieldPayload},
			transaction.FieldState:         {Type: field.TypeInt32, Column: transaction.FieldState},
			transaction.FieldLockTime:      {Type: field.TypeUint32, Column: transaction.FieldLockTime},
			transaction.FieldCreatedAt:     {Type: field.TypeUint32, Column: transaction.FieldCreatedAt},
			transaction.FieldUpdatedAt:     {Type: field.TypeUint32, Column: transaction.FieldUpdatedAt},
			transaction.FieldDeletedAt:     {Type: field.TypeUint32, Column: transaction.FieldDeletedAt},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (rciq *RegCoinInfoQuery) addPredicate(pred func(s *sql.Selector)) {
	rciq.predicates = append(rciq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the RegCoinInfoQuery builder.
func (rciq *RegCoinInfoQuery) Filter() *RegCoinInfoFilter {
	return &RegCoinInfoFilter{config: rciq.config, predicateAdder: rciq}
}

// addPredicate implements the predicateAdder interface.
func (m *RegCoinInfoMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the RegCoinInfoMutation builder.
func (m *RegCoinInfoMutation) Filter() *RegCoinInfoFilter {
	return &RegCoinInfoFilter{config: m.config, predicateAdder: m}
}

// RegCoinInfoFilter provides a generic filtering capability at runtime for RegCoinInfoQuery.
type RegCoinInfoFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *RegCoinInfoFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *RegCoinInfoFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(regcoininfo.FieldID))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *RegCoinInfoFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(regcoininfo.FieldEntID))
}

// WhereChainType applies the entql int32 predicate on the chain_type field.
func (f *RegCoinInfoFilter) WhereChainType(p entql.Int32P) {
	f.Where(p.Field(regcoininfo.FieldChainType))
}

// WhereCoinType applies the entql int32 predicate on the coin_type field.
func (f *RegCoinInfoFilter) WhereCoinType(p entql.Int32P) {
	f.Where(p.Field(regcoininfo.FieldCoinType))
}

// WhereTempName applies the entql string predicate on the temp_name field.
func (f *RegCoinInfoFilter) WhereTempName(p entql.StringP) {
	f.Where(p.Field(regcoininfo.FieldTempName))
}

// WhereName applies the entql string predicate on the name field.
func (f *RegCoinInfoFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(regcoininfo.FieldName))
}

// WhereEnv applies the entql string predicate on the env field.
func (f *RegCoinInfoFilter) WhereEnv(p entql.StringP) {
	f.Where(p.Field(regcoininfo.FieldEnv))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *RegCoinInfoFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(regcoininfo.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *RegCoinInfoFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(regcoininfo.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *RegCoinInfoFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(regcoininfo.FieldDeletedAt))
}

// addPredicate implements the predicateAdder interface.
func (tq *TransactionQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TransactionQuery builder.
func (tq *TransactionQuery) Filter() *TransactionFilter {
	return &TransactionFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TransactionMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TransactionMutation builder.
func (m *TransactionMutation) Filter() *TransactionFilter {
	return &TransactionFilter{config: m.config, predicateAdder: m}
}

// TransactionFilter provides a generic filtering capability at runtime for TransactionQuery.
type TransactionFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TransactionFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *TransactionFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(transaction.FieldID))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *TransactionFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(transaction.FieldEntID))
}

// WhereCoinType applies the entql int32 predicate on the coin_type field.
func (f *TransactionFilter) WhereCoinType(p entql.Int32P) {
	f.Where(p.Field(transaction.FieldCoinType))
}

// WhereChainType applies the entql int32 predicate on the chain_type field.
func (f *TransactionFilter) WhereChainType(p entql.Int32P) {
	f.Where(p.Field(transaction.FieldChainType))
}

// WhereClientType applies the entql int32 predicate on the client_type field.
func (f *TransactionFilter) WhereClientType(p entql.Int32P) {
	f.Where(p.Field(transaction.FieldClientType))
}

// WhereTransactionID applies the entql string predicate on the transaction_id field.
func (f *TransactionFilter) WhereTransactionID(p entql.StringP) {
	f.Where(p.Field(transaction.FieldTransactionID))
}

// WhereCid applies the entql string predicate on the cid field.
func (f *TransactionFilter) WhereCid(p entql.StringP) {
	f.Where(p.Field(transaction.FieldCid))
}

// WhereExitCode applies the entql int64 predicate on the exit_code field.
func (f *TransactionFilter) WhereExitCode(p entql.Int64P) {
	f.Where(p.Field(transaction.FieldExitCode))
}

// WhereName applies the entql string predicate on the name field.
func (f *TransactionFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(transaction.FieldName))
}

// WhereFrom applies the entql string predicate on the from field.
func (f *TransactionFilter) WhereFrom(p entql.StringP) {
	f.Where(p.Field(transaction.FieldFrom))
}

// WhereTo applies the entql string predicate on the to field.
func (f *TransactionFilter) WhereTo(p entql.StringP) {
	f.Where(p.Field(transaction.FieldTo))
}

// WhereMemo applies the entql string predicate on the memo field.
func (f *TransactionFilter) WhereMemo(p entql.StringP) {
	f.Where(p.Field(transaction.FieldMemo))
}

// WhereAmount applies the entql uint64 predicate on the amount field.
func (f *TransactionFilter) WhereAmount(p entql.Uint64P) {
	f.Where(p.Field(transaction.FieldAmount))
}

// WherePayload applies the entql []byte predicate on the payload field.
func (f *TransactionFilter) WherePayload(p entql.BytesP) {
	f.Where(p.Field(transaction.FieldPayload))
}

// WhereState applies the entql int32 predicate on the state field.
func (f *TransactionFilter) WhereState(p entql.Int32P) {
	f.Where(p.Field(transaction.FieldState))
}

// WhereLockTime applies the entql uint32 predicate on the lock_time field.
func (f *TransactionFilter) WhereLockTime(p entql.Uint32P) {
	f.Where(p.Field(transaction.FieldLockTime))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *TransactionFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(transaction.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *TransactionFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(transaction.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *TransactionFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(transaction.FieldDeletedAt))
}
