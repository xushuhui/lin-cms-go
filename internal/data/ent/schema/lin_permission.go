package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type LinPermission struct {
	ent.Schema
}

func (LinPermission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("权限名称，例如：访问首页"),
		field.String("module").Comment("权限所属模块，例如：人员管理"),
		field.Int8("mount").Comment("0：关闭 1：开启"),
		field.Time("create_time").Comment(""),
		field.Time("update_time").Comment(""),
		field.Time("delete_time").Comment(""),
	}
}
