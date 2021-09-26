package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type LinGroupPermission struct {
	ent.Schema
}

func (LinGroupPermission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("group_id").Comment("分组id").Unique(),
		field.Int("permission_id").Comment("权限id"),
	}
}
