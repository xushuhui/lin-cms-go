package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type LinUser struct {
	ent.Schema
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
		field.String("avatar").Comment("头像url"),
		field.String("email").Comment("邮箱").Unique(),
	}
}
