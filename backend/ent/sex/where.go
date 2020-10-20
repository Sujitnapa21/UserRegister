// Code generated by entc, DO NOT EDIT.

package sex

import (
	"github.com/Sujitnapa21/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Sexname applies equality check predicate on the "sexname" field. It's identical to SexnameEQ.
func Sexname(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSexname), v))
	})
}

// SexnameEQ applies the EQ predicate on the "sexname" field.
func SexnameEQ(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSexname), v))
	})
}

// SexnameNEQ applies the NEQ predicate on the "sexname" field.
func SexnameNEQ(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSexname), v))
	})
}

// SexnameIn applies the In predicate on the "sexname" field.
func SexnameIn(vs ...string) predicate.Sex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSexname), v...))
	})
}

// SexnameNotIn applies the NotIn predicate on the "sexname" field.
func SexnameNotIn(vs ...string) predicate.Sex {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sex(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSexname), v...))
	})
}

// SexnameGT applies the GT predicate on the "sexname" field.
func SexnameGT(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSexname), v))
	})
}

// SexnameGTE applies the GTE predicate on the "sexname" field.
func SexnameGTE(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSexname), v))
	})
}

// SexnameLT applies the LT predicate on the "sexname" field.
func SexnameLT(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSexname), v))
	})
}

// SexnameLTE applies the LTE predicate on the "sexname" field.
func SexnameLTE(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSexname), v))
	})
}

// SexnameContains applies the Contains predicate on the "sexname" field.
func SexnameContains(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSexname), v))
	})
}

// SexnameHasPrefix applies the HasPrefix predicate on the "sexname" field.
func SexnameHasPrefix(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSexname), v))
	})
}

// SexnameHasSuffix applies the HasSuffix predicate on the "sexname" field.
func SexnameHasSuffix(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSexname), v))
	})
}

// SexnameEqualFold applies the EqualFold predicate on the "sexname" field.
func SexnameEqualFold(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSexname), v))
	})
}

// SexnameContainsFold applies the ContainsFold predicate on the "sexname" field.
func SexnameContainsFold(v string) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSexname), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Sex) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Sex) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
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
func Not(p predicate.Sex) predicate.Sex {
	return predicate.Sex(func(s *sql.Selector) {
		p(s.Not())
	})
}
