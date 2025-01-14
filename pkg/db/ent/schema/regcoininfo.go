package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
)

// RegCoinInfo holds the schema definition for the RegCoinInfo entity.
type RegCoinInfo struct {
	ent.Schema
}

func (RegCoinInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the RegCoinInfo.
func (RegCoinInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("chain_type").
			Optional().
			Default(0),
		field.Int32("coin_type").
			Optional().
			Default(0),
		field.String("temp_name").
			Optional().
			Default(""),
		field.String("name").
			Unique(),
		field.String("env").
			Optional().
			Default(""),
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

func (RegCoinInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}
