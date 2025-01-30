package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "permissions"},
		entsql.WithComments(true),
		schema.Comment("权限"),
	}
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("created_at").Immutable().Default(0).Comment("创建时间"),
		field.Int("updated_at").Default(0).Comment("修改时间"),
		field.Int("deleted_at").Optional().Nillable().Comment("删除时间"),
		field.String("tenant_code").NotEmpty().Comment("租户编码"),

		field.String("name").NotEmpty().Comment("权限名称"),
		field.String("code").Unique().NotEmpty().Comment("权限编码"),
		field.Int("type").Comment("类型 1:菜单menu 2:按钮buttn 3:接口 api"),
		field.String("path").Optional().Comment("路径"),
		field.Int("action").Default(1).Optional().Comment("操作类型: 1:all 2:read 3:write"),
		field.Int("parent_id").Optional().Comment("父级ID"),
		field.String("description").Optional().Comment("描述"),
		field.Int("status").Default(1).Comment("状态 1:启用 2:禁用"),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{}
}
