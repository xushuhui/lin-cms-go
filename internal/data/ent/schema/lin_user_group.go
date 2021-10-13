package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type LinUserGroup struct {
	ent.Schema
}
func (LinUserGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "lin_user_group"},
	}
}
func (LinUserGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("用户id").Unique(),
		field.Int("group_id").Comment("分组id"),
	}
}
