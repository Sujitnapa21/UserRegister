package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Sex schema.
type Sex struct {
	ent.Schema
}

// Fields of the Sex.
func (Sex) Fields() []ent.Field {
	return []ent.Field{
		field.String("sexname").
			NotEmpty(),
	}
}

// Edges of the Sex.
func (Sex) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
	}
}
