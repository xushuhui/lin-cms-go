package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type LinLog struct {
	ent.Schema
}
func (LinLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "lin_log"},
	}
}
func (LinLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("message").Comment(""),
		field.Int("user_id").Comment(""),
		field.String("username").Comment(""),
		field.Int("status_code").Comment(""),
		field.String("method").Comment(""),
		field.String("path").Comment(""),
		field.String("permission").Comment(""),
		field.String("create_time").Comment(""),
		field.String("update_time").Comment(""),
		field.String("delete_time").Comment(""),
	}
}
