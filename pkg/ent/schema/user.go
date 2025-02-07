package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
		entsql.WithComments(true),
		schema.Comment("用户"),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("created_at").Immutable().Default(0).Comment("创建时间"),
		field.Int64("updated_at").Default(0).Comment("修改时间"),
		field.Int64("deleted_at").Optional().Nillable().Comment("删除时间"),
		field.Uint64("tenant_id").Comment("租户ID"),

		field.Uint64("user_id").Unique().Comment("用户ID"),
		field.String("user_name").NotEmpty().Default("").Comment("用户名"),
		field.String("password").NotEmpty().Default("").Comment("密码"),
		field.String("salt").NotEmpty().Default("").Comment("密码加盐"),
		field.String("token").Default("").Comment("登录后的token信息"),

		field.String("nick_name").Default("").Comment("昵称"),
		field.String("avatar").Default("").Comment("头像"),
		field.String("phone").NotEmpty().Default("").Comment("电话"),
		field.String("email").Default("").Comment("邮箱"),
		field.Int("sex").Default(0).Comment("性别: 1：男 2：女"),
		field.Int("status").Default(1).Comment("状态: 1:启用, 2:禁用"),

		field.Uint64("role_id").Default(0).Comment("角色ID"),
		field.Uint64("dept_id").Default(0).Comment("部门ID"),
		field.Uint64("post_id").Default(0).Comment("岗位ID"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("phone"),
		index.Fields("user_id"),
	}
}
