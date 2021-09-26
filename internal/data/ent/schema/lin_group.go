package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type LinGroup struct {
	ent.Schema
}

func (LinGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("分组名称，例如：搬砖者").Unique(),
		field.String("info").Comment("分组信息：例如：搬砖的人"),
		field.Int8("level").Comment("分组级别 1：root 2：guest 3：user（root、guest分组只能存在一个)"),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
	}
}
