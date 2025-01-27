package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("created_at").Immutable().Default(0).Comment("创建时间"),
		field.Int("updated_at").Default(0).Comment("修改时间"),
		field.Int("deleted_at").Optional().Nillable().Comment("删除时间"),
	}
}
