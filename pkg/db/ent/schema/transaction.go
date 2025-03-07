package schema

import (
	"math"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

func (Transaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("coin_type").
			Optional().
			Default(0),
		field.Int32("chain_type").
			Optional().
			Default(0),
		field.Int32("client_type").
			Optional().
			Default(0),
		field.String("transaction_id").
			Unique(),
		field.String("cid").
			Optional().
			Default(""),
		field.Int64("exit_code").
			Optional().
			Default(0),
		field.String("name").
			Optional().
			Default(""),
		field.String("from").
			Optional().
			Default(""),
		field.String("to").
			Optional().
			Default(""),
		field.String("memo").
			Optional().
			Default(""),
		field.Uint64("amount").
			Optional().
			Default(0),
		field.Bytes("payload").
			Optional().
			MaxLen(math.MaxUint32).
			Default([]byte{}).
			Comment("save nonce or sign info"),
		field.Int32("state").
			Optional().
			Default(0),
		field.Uint32("lock_time").
			Optional().
			Default(0),
		field.Uint32("created_at").
			Optional().
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("updated_at").
			Optional().
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("deleted_at").
			Optional().
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}
