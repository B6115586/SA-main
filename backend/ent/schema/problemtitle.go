package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// ProblemTitle holds the schema definition for the ProblemTitle entity.
type ProblemTitle struct {
	ent.Schema
}

// Fields of the ProblemTitle.
func (ProblemTitle) Fields() []ent.Field {
	return []ent.Field{
		field.String("problemtitle").NotEmpty(),
	}
}

// Edges of the ProblemTitle.
func (ProblemTitle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("problemtitle2problem", Problem.Type).StorageKey(edge.Column("problemtitle")),
	}
}
