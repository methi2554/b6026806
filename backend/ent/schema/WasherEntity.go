package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Washer holds the schema definition for the Washer entity.
type Washer struct {
	ent.Schema
}

// Fields of the Washer.
func (Washer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("coins").NotEmpty(),
	}
}

// Edges of the Washer.
func (Washer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("Washers").Unique(),
	}
}
