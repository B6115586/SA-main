package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// ProblemStatus holds the schema definition for the ProblemStatus entity.
type ProblemStatus struct {
	ent.Schema
}

// Fields of the ProblemStatus.
func (ProblemStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("problemstatus").NotEmpty(),
	}
}

// Edges of the ProblemStatus.
func (ProblemStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("problemstatus2problem", Problem.Type).StorageKey(edge.Column("problemstatus")),
	}
}
