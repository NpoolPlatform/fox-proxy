// Code generated by ent, DO NOT EDIT.

package regcoininfo

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// ChainType applies equality check predicate on the "chain_type" field. It's identical to ChainTypeEQ.
func ChainType(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainType), v))
	})
}

// CoinType applies equality check predicate on the "coin_type" field. It's identical to CoinTypeEQ.
func CoinType(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinType), v))
	})
}

// TempName applies equality check predicate on the "temp_name" field. It's identical to TempNameEQ.
func TempName(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTempName), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Env applies equality check predicate on the "env" field. It's identical to EnvEQ.
func Env(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnv), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// ChainTypeEQ applies the EQ predicate on the "chain_type" field.
func ChainTypeEQ(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainType), v))
	})
}

// ChainTypeNEQ applies the NEQ predicate on the "chain_type" field.
func ChainTypeNEQ(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChainType), v))
	})
}

// ChainTypeIn applies the In predicate on the "chain_type" field.
func ChainTypeIn(vs ...int32) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldChainType), v...))
	})
}

// ChainTypeNotIn applies the NotIn predicate on the "chain_type" field.
func ChainTypeNotIn(vs ...int32) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldChainType), v...))
	})
}

// ChainTypeGT applies the GT predicate on the "chain_type" field.
func ChainTypeGT(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChainType), v))
	})
}

// ChainTypeGTE applies the GTE predicate on the "chain_type" field.
func ChainTypeGTE(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChainType), v))
	})
}

// ChainTypeLT applies the LT predicate on the "chain_type" field.
func ChainTypeLT(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChainType), v))
	})
}

// ChainTypeLTE applies the LTE predicate on the "chain_type" field.
func ChainTypeLTE(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChainType), v))
	})
}

// ChainTypeIsNil applies the IsNil predicate on the "chain_type" field.
func ChainTypeIsNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldChainType)))
	})
}

// ChainTypeNotNil applies the NotNil predicate on the "chain_type" field.
func ChainTypeNotNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldChainType)))
	})
}

// CoinTypeEQ applies the EQ predicate on the "coin_type" field.
func CoinTypeEQ(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinType), v))
	})
}

// CoinTypeNEQ applies the NEQ predicate on the "coin_type" field.
func CoinTypeNEQ(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinType), v))
	})
}

// CoinTypeIn applies the In predicate on the "coin_type" field.
func CoinTypeIn(vs ...int32) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinType), v...))
	})
}

// CoinTypeNotIn applies the NotIn predicate on the "coin_type" field.
func CoinTypeNotIn(vs ...int32) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinType), v...))
	})
}

// CoinTypeGT applies the GT predicate on the "coin_type" field.
func CoinTypeGT(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinType), v))
	})
}

// CoinTypeGTE applies the GTE predicate on the "coin_type" field.
func CoinTypeGTE(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinType), v))
	})
}

// CoinTypeLT applies the LT predicate on the "coin_type" field.
func CoinTypeLT(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinType), v))
	})
}

// CoinTypeLTE applies the LTE predicate on the "coin_type" field.
func CoinTypeLTE(v int32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinType), v))
	})
}

// CoinTypeIsNil applies the IsNil predicate on the "coin_type" field.
func CoinTypeIsNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinType)))
	})
}

// CoinTypeNotNil applies the NotNil predicate on the "coin_type" field.
func CoinTypeNotNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinType)))
	})
}

// TempNameEQ applies the EQ predicate on the "temp_name" field.
func TempNameEQ(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTempName), v))
	})
}

// TempNameNEQ applies the NEQ predicate on the "temp_name" field.
func TempNameNEQ(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTempName), v))
	})
}

// TempNameIn applies the In predicate on the "temp_name" field.
func TempNameIn(vs ...string) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTempName), v...))
	})
}

// TempNameNotIn applies the NotIn predicate on the "temp_name" field.
func TempNameNotIn(vs ...string) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTempName), v...))
	})
}

// TempNameGT applies the GT predicate on the "temp_name" field.
func TempNameGT(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTempName), v))
	})
}

// TempNameGTE applies the GTE predicate on the "temp_name" field.
func TempNameGTE(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTempName), v))
	})
}

// TempNameLT applies the LT predicate on the "temp_name" field.
func TempNameLT(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTempName), v))
	})
}

// TempNameLTE applies the LTE predicate on the "temp_name" field.
func TempNameLTE(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTempName), v))
	})
}

// TempNameContains applies the Contains predicate on the "temp_name" field.
func TempNameContains(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTempName), v))
	})
}

// TempNameHasPrefix applies the HasPrefix predicate on the "temp_name" field.
func TempNameHasPrefix(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTempName), v))
	})
}

// TempNameHasSuffix applies the HasSuffix predicate on the "temp_name" field.
func TempNameHasSuffix(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTempName), v))
	})
}

// TempNameIsNil applies the IsNil predicate on the "temp_name" field.
func TempNameIsNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTempName)))
	})
}

// TempNameNotNil applies the NotNil predicate on the "temp_name" field.
func TempNameNotNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTempName)))
	})
}

// TempNameEqualFold applies the EqualFold predicate on the "temp_name" field.
func TempNameEqualFold(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTempName), v))
	})
}

// TempNameContainsFold applies the ContainsFold predicate on the "temp_name" field.
func TempNameContainsFold(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTempName), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// EnvEQ applies the EQ predicate on the "env" field.
func EnvEQ(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnv), v))
	})
}

// EnvNEQ applies the NEQ predicate on the "env" field.
func EnvNEQ(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnv), v))
	})
}

// EnvIn applies the In predicate on the "env" field.
func EnvIn(vs ...string) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEnv), v...))
	})
}

// EnvNotIn applies the NotIn predicate on the "env" field.
func EnvNotIn(vs ...string) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEnv), v...))
	})
}

// EnvGT applies the GT predicate on the "env" field.
func EnvGT(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnv), v))
	})
}

// EnvGTE applies the GTE predicate on the "env" field.
func EnvGTE(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnv), v))
	})
}

// EnvLT applies the LT predicate on the "env" field.
func EnvLT(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnv), v))
	})
}

// EnvLTE applies the LTE predicate on the "env" field.
func EnvLTE(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnv), v))
	})
}

// EnvContains applies the Contains predicate on the "env" field.
func EnvContains(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEnv), v))
	})
}

// EnvHasPrefix applies the HasPrefix predicate on the "env" field.
func EnvHasPrefix(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEnv), v))
	})
}

// EnvHasSuffix applies the HasSuffix predicate on the "env" field.
func EnvHasSuffix(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEnv), v))
	})
}

// EnvIsNil applies the IsNil predicate on the "env" field.
func EnvIsNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEnv)))
	})
}

// EnvNotNil applies the NotNil predicate on the "env" field.
func EnvNotNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEnv)))
	})
}

// EnvEqualFold applies the EqualFold predicate on the "env" field.
func EnvEqualFold(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEnv), v))
	})
}

// EnvContainsFold applies the ContainsFold predicate on the "env" field.
func EnvContainsFold(v string) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEnv), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedAt)))
	})
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedAt)))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.RegCoinInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RegCoinInfo) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RegCoinInfo) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RegCoinInfo) predicate.RegCoinInfo {
	return predicate.RegCoinInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}