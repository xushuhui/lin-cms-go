package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Book struct {
	ent.Schema
}

func (Book) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "book"},
	}
}
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment(""),
		field.String("author").Comment(""),
		field.String("summary").Comment(""),
		field.String("image").Comment(""),
	}
}
