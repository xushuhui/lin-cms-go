package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type TimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_time").
			Immutable().
			Default(time.Now),
		field.Time("updated_time").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

//type LinUserGroup struct {
//	ent.Schema
//}
//func (LinUserGroup) Annotations() []schema.Annotation {
//	return []schema.Annotation{
//		entsql.Annotation{Table: "lin_user_group"},
//	}
//}
//func (LinUserGroup) Fields() []ent.Field {
//	return []ent.Field{
//		field.Int("user_id").Comment("用户id").Unique(),
//		field.Int("group_id").Comment("分组id"),
//	}
//}
