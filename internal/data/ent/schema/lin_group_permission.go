package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type LinGroupPermission struct {
	ent.Schema
}
func (LinGroupPermission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "lin_group_permission"},
	}
}
func (LinGroupPermission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("group_id").Comment("分组id").Unique(),
		field.Int("permission_id").Comment("权限id"),
	}
}
