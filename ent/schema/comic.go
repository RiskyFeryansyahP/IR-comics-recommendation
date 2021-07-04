package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comic holds the schema definition for the Comic entity.
type Comic struct {
	ent.Schema
}

// Fields of the Comic.
func (Comic) Fields() []ent.Field {
	return []ent.Field{
		field.String("Title"),
		field.String("Author"),
		field.String("Like"),
	}
}

// Edges of the Comic.
func (Comic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("genres", Genre.Type),
	}
}
