// Code generated by entc, DO NOT EDIT.

package role

import (
	"github.com/Sujitnapa21/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Rolename applies equality check predicate on the "rolename" field. It's identical to RolenameEQ.
func Rolename(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRolename), v))
	})
}

// RolenameEQ applies the EQ predicate on the "rolename" field.
func RolenameEQ(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRolename), v))
	})
}

// RolenameNEQ applies the NEQ predicate on the "rolename" field.
func RolenameNEQ(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRolename), v))
	})
}

// RolenameIn applies the In predicate on the "rolename" field.
func RolenameIn(vs ...string) predicate.Role {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Role(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRolename), v...))
	})
}

// RolenameNotIn applies the NotIn predicate on the "rolename" field.
func RolenameNotIn(vs ...string) predicate.Role {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Role(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRolename), v...))
	})
}

// RolenameGT applies the GT predicate on the "rolename" field.
func RolenameGT(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRolename), v))
	})
}

// RolenameGTE applies the GTE predicate on the "rolename" field.
func RolenameGTE(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRolename), v))
	})
}

// RolenameLT applies the LT predicate on the "rolename" field.
func RolenameLT(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRolename), v))
	})
}

// RolenameLTE applies the LTE predicate on the "rolename" field.
func RolenameLTE(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRolename), v))
	})
}

// RolenameContains applies the Contains predicate on the "rolename" field.
func RolenameContains(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRolename), v))
	})
}

// RolenameHasPrefix applies the HasPrefix predicate on the "rolename" field.
func RolenameHasPrefix(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRolename), v))
	})
}

// RolenameHasSuffix applies the HasSuffix predicate on the "rolename" field.
func RolenameHasSuffix(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRolename), v))
	})
}

// RolenameEqualFold applies the EqualFold predicate on the "rolename" field.
func RolenameEqualFold(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRolename), v))
	})
}

// RolenameContainsFold applies the ContainsFold predicate on the "rolename" field.
func RolenameContainsFold(v string) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRolename), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
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
func And(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
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
func Not(p predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		p(s.Not())
	})
}
