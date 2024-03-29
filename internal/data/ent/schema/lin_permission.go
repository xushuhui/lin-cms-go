package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type LinPermission struct {
	ent.Schema
}

func (LinPermission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "lin_permission"},
	}
}
func (LinPermission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("权限名称，例如：访问首页"),
		field.String("module").Comment("权限所属模块，例如：人员管理"),
		field.Int8("mount").Comment("0：关闭 1：开启"),
	}
}
func (LinPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lin_group", LinGroup.Type).StorageKey(
			edge.Table("lin_group_permission"), edge.Columns("permission_id", "group_id"),
		),
	}
}
