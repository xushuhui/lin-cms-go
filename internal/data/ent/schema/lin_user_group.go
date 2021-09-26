package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type LinUserGroup struct {
	ent.Schema
}

func (LinUserGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("用户id").Unique(),
		field.Int("group_id").Comment("分组id"),
	}
}
