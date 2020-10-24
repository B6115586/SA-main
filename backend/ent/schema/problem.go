package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.String("probleminfo").NotEmpty(),
		field.Time("adddate"),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user2problem").Unique(),
		edge.From("room", Room.Type).Ref("room2problem").Unique(),
		edge.From("problemtitle", ProblemTitle.Type).Ref("problemtitle2problem").Unique(),
		edge.From("problemstatus", ProblemStatus.Type).Ref("problemstatus2problem").Unique(),
	}

}
