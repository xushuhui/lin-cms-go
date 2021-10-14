package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type LinUser struct {
	ent.Schema
}

func (LinUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "lin_user"},
	}
}
func (LinUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
func (LinUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Comment("用户名，唯一").Unique(),
		field.String("nickname").Comment("用户昵称"),
		field.String("avatar").Default("").Comment("头像url"),
		field.String("email").Comment("邮箱").Unique(),
	}
}
func (LinUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lin_user_identiy", LinUserIdentiy.Type),
		edge.From("lin_group", LinGroup.Type).Ref("lin_user"),
	}
}
