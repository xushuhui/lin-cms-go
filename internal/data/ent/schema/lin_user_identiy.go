package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type LinUserIdentiy struct {
	ent.Schema
}

func (LinUserIdentiy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "lin_user_identiy"},
	}
}
func (LinUserIdentiy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("用户id"),
		field.String("identity_type").Comment(""),
		field.String("identifier").Comment(""),
		field.String("credential").Comment(""),
	}
}

//func (LinUserIdentiy) Mixin() []ent.Mixin {
//	return []ent.Mixin{
//		mixin.Time{},
//	}
//}
