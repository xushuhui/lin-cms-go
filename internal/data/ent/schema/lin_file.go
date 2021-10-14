package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type LinFile struct {
	ent.Schema
}

func (LinFile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "lin_file"},
	}
}
func (LinFile) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").Comment(""),
		field.Int8("type").Comment("1 LOCAL 本地，2 REMOTE 远程"),
		field.String("name").Comment(""),
		field.String("extension").Comment(""),
		field.Int("size").Comment(""),
		field.String("md5").Comment("md5值，防止上传重复文件").Unique(),
	}
}
