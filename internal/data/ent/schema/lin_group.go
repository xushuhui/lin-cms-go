package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type LinGroup struct {
	ent.Schema
}
func (LinGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "lin_user_group"},
	}
}
func (LinGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("分组名称，例如：搬砖者").Unique(),
		field.String("info").Comment("分组信息：例如：搬砖的人"),
		field.Int8("level").Comment("分组级别 1：root 2：guest 3：user（root、guest分组只能存在一个)"),
		field.String("create_time").Comment(""),
		field.String("update_time").Comment(""),
		field.String("delete_time").Comment(""),
	}
}
func (LinGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lin_user", LinUser.Type),
		edge.From("lin_permission",LinPermission.Type).Ref("lin_group"),
	}
}
