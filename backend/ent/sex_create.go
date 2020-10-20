// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Sujitnapa21/app/ent/sex"
	"github.com/Sujitnapa21/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// SexCreate is the builder for creating a Sex entity.
type SexCreate struct {
	config
	mutation *SexMutation
	hooks    []Hook
}

// SetSexname sets the sexname field.
func (sc *SexCreate) SetSexname(s string) *SexCreate {
	sc.mutation.SetSexname(s)
	return sc
}

// AddUserIDs adds the user edge to User by ids.
func (sc *SexCreate) AddUserIDs(ids ...int) *SexCreate {
	sc.mutation.AddUserIDs(ids...)
	return sc
}

// AddUser adds the user edges to User.
func (sc *SexCreate) AddUser(u ...*User) *SexCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return sc.AddUserIDs(ids...)
}

// Mutation returns the SexMutation object of the builder.
func (sc *SexCreate) Mutation() *SexMutation {
	return sc.mutation
}

// Save creates the Sex in the database.
func (sc *SexCreate) Save(ctx context.Context) (*Sex, error) {
	if _, ok := sc.mutation.Sexname(); !ok {
		return nil, &ValidationError{Name: "sexname", err: errors.New("ent: missing required field \"sexname\"")}
	}
	if v, ok := sc.mutation.Sexname(); ok {
		if err := sex.SexnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "sexname", err: fmt.Errorf("ent: validator failed for field \"sexname\": %w", err)}
		}
	}
	var (
		err  error
		node *Sex
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SexMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SexCreate) SaveX(ctx context.Context) *Sex {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SexCreate) sqlSave(ctx context.Context) (*Sex, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *SexCreate) createSpec() (*Sex, *sqlgraph.CreateSpec) {
	var (
		s     = &Sex{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sex.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sex.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Sexname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sex.FieldSexname,
		})
		s.Sexname = value
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sex.UserTable,
			Columns: []string{sex.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}
