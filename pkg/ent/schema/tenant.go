package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

func (Tenant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tenants"},
		entsql.WithComments(true),
		schema.Comment("租户"),
	}
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.Int("created_at").Immutable().Default(0).Comment("创建时间"),
		field.Int("updated_at").Default(0).Comment("修改时间"),
		field.Int("deleted_at").Optional().Nillable().Comment("删除时间"),

		// field.Uint64("tenant_id").Unique().Immutable().Comment("租户ID"),
		field.String("name").NotEmpty().Default("").Comment("租户名称"),
		field.String("code").NotEmpty().Default("").Comment("租户编码"),
		field.String("description").Default("").Comment("租户描述"),
		field.Int("status").Default(1).Comment("租户状态：1: 启用, 2: 禁用"),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Tenant) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("tenant_id"),
		index.Fields("code").Unique(),
	}
}
