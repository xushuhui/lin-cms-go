package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"entgo.io/ent/schema/mixin"
	"time"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").
			Immutable().
			Default(time.Now),
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("delete_time").
			Optional(),
	}
}
