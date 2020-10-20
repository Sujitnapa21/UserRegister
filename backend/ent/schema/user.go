package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User schema.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("password"),
		field.String("name"),
		field.String("age"),
		field.Time("birthday"),
		field.String("tel"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sex", Sex.Type).Ref("user").Unique(),
		edge.From("role", Role.Type).Ref("user").Unique(),
		edge.From("title", Title.Type).Ref("user").Unique(),
	}
}
