// Code generated by entc, DO NOT EDIT.

package problem

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/thanawat/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Probleminfo applies equality check predicate on the "probleminfo" field. It's identical to ProbleminfoEQ.
func Probleminfo(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProbleminfo), v))
	})
}

// Adddate applies equality check predicate on the "adddate" field. It's identical to AdddateEQ.
func Adddate(v time.Time) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdddate), v))
	})
}

// ProbleminfoEQ applies the EQ predicate on the "probleminfo" field.
func ProbleminfoEQ(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProbleminfo), v))
	})
}

// ProbleminfoNEQ applies the NEQ predicate on the "probleminfo" field.
func ProbleminfoNEQ(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProbleminfo), v))
	})
}

// ProbleminfoIn applies the In predicate on the "probleminfo" field.
func ProbleminfoIn(vs ...string) predicate.Problem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Problem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProbleminfo), v...))
	})
}

// ProbleminfoNotIn applies the NotIn predicate on the "probleminfo" field.
func ProbleminfoNotIn(vs ...string) predicate.Problem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Problem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProbleminfo), v...))
	})
}

// ProbleminfoGT applies the GT predicate on the "probleminfo" field.
func ProbleminfoGT(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProbleminfo), v))
	})
}

// ProbleminfoGTE applies the GTE predicate on the "probleminfo" field.
func ProbleminfoGTE(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProbleminfo), v))
	})
}

// ProbleminfoLT applies the LT predicate on the "probleminfo" field.
func ProbleminfoLT(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProbleminfo), v))
	})
}

// ProbleminfoLTE applies the LTE predicate on the "probleminfo" field.
func ProbleminfoLTE(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProbleminfo), v))
	})
}

// ProbleminfoContains applies the Contains predicate on the "probleminfo" field.
func ProbleminfoContains(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProbleminfo), v))
	})
}

// ProbleminfoHasPrefix applies the HasPrefix predicate on the "probleminfo" field.
func ProbleminfoHasPrefix(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProbleminfo), v))
	})
}

// ProbleminfoHasSuffix applies the HasSuffix predicate on the "probleminfo" field.
func ProbleminfoHasSuffix(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProbleminfo), v))
	})
}

// ProbleminfoEqualFold applies the EqualFold predicate on the "probleminfo" field.
func ProbleminfoEqualFold(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProbleminfo), v))
	})
}

// ProbleminfoContainsFold applies the ContainsFold predicate on the "probleminfo" field.
func ProbleminfoContainsFold(v string) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProbleminfo), v))
	})
}

// AdddateEQ applies the EQ predicate on the "adddate" field.
func AdddateEQ(v time.Time) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdddate), v))
	})
}

// AdddateNEQ applies the NEQ predicate on the "adddate" field.
func AdddateNEQ(v time.Time) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdddate), v))
	})
}

// AdddateIn applies the In predicate on the "adddate" field.
func AdddateIn(vs ...time.Time) predicate.Problem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Problem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAdddate), v...))
	})
}

// AdddateNotIn applies the NotIn predicate on the "adddate" field.
func AdddateNotIn(vs ...time.Time) predicate.Problem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Problem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAdddate), v...))
	})
}

// AdddateGT applies the GT predicate on the "adddate" field.
func AdddateGT(v time.Time) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAdddate), v))
	})
}

// AdddateGTE applies the GTE predicate on the "adddate" field.
func AdddateGTE(v time.Time) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAdddate), v))
	})
}

// AdddateLT applies the LT predicate on the "adddate" field.
func AdddateLT(v time.Time) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAdddate), v))
	})
}

// AdddateLTE applies the LTE predicate on the "adddate" field.
func AdddateLTE(v time.Time) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAdddate), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoom applies the HasEdge predicate on the "room" edge.
func HasRoom() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomWith applies the HasEdge predicate on the "room" edge with a given conditions (other predicates).
func HasRoomWith(preds ...predicate.Room) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProblemtitle applies the HasEdge predicate on the "problemtitle" edge.
func HasProblemtitle() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProblemtitleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProblemtitleTable, ProblemtitleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProblemtitleWith applies the HasEdge predicate on the "problemtitle" edge with a given conditions (other predicates).
func HasProblemtitleWith(preds ...predicate.ProblemTitle) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProblemtitleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProblemtitleTable, ProblemtitleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProblemstatus applies the HasEdge predicate on the "problemstatus" edge.
func HasProblemstatus() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProblemstatusTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProblemstatusTable, ProblemstatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProblemstatusWith applies the HasEdge predicate on the "problemstatus" edge with a given conditions (other predicates).
func HasProblemstatusWith(preds ...predicate.ProblemStatus) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProblemstatusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProblemstatusTable, ProblemstatusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Problem) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Problem) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Problem) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		p(s.Not())
	})
}