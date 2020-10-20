// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Sujitnapa21/app/ent/role"
	"github.com/Sujitnapa21/app/ent/sex"
	"github.com/Sujitnapa21/app/ent/title"
	"github.com/Sujitnapa21/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetEmail sets the email field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetPassword sets the password field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetAge sets the age field.
func (uc *UserCreate) SetAge(s string) *UserCreate {
	uc.mutation.SetAge(s)
	return uc
}

// SetBirthday sets the birthday field.
func (uc *UserCreate) SetBirthday(t time.Time) *UserCreate {
	uc.mutation.SetBirthday(t)
	return uc
}

// SetTel sets the tel field.
func (uc *UserCreate) SetTel(s string) *UserCreate {
	uc.mutation.SetTel(s)
	return uc
}

// SetSexID sets the sex edge to Sex by id.
func (uc *UserCreate) SetSexID(id int) *UserCreate {
	uc.mutation.SetSexID(id)
	return uc
}

// SetNillableSexID sets the sex edge to Sex by id if the given value is not nil.
func (uc *UserCreate) SetNillableSexID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetSexID(*id)
	}
	return uc
}

// SetSex sets the sex edge to Sex.
func (uc *UserCreate) SetSex(s *Sex) *UserCreate {
	return uc.SetSexID(s.ID)
}

// SetRoleID sets the role edge to Role by id.
func (uc *UserCreate) SetRoleID(id int) *UserCreate {
	uc.mutation.SetRoleID(id)
	return uc
}

// SetNillableRoleID sets the role edge to Role by id if the given value is not nil.
func (uc *UserCreate) SetNillableRoleID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetRoleID(*id)
	}
	return uc
}

// SetRole sets the role edge to Role.
func (uc *UserCreate) SetRole(r *Role) *UserCreate {
	return uc.SetRoleID(r.ID)
}

// SetTitleID sets the title edge to Title by id.
func (uc *UserCreate) SetTitleID(id int) *UserCreate {
	uc.mutation.SetTitleID(id)
	return uc
}

// SetNillableTitleID sets the title edge to Title by id if the given value is not nil.
func (uc *UserCreate) SetNillableTitleID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetTitleID(*id)
	}
	return uc
}

// SetTitle sets the title edge to Title.
func (uc *UserCreate) SetTitle(t *Title) *UserCreate {
	return uc.SetTitleID(t.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if _, ok := uc.mutation.Email(); !ok {
		return nil, &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return nil, &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := uc.mutation.Age(); !ok {
		return nil, &ValidationError{Name: "age", err: errors.New("ent: missing required field \"age\"")}
	}
	if _, ok := uc.mutation.Birthday(); !ok {
		return nil, &ValidationError{Name: "birthday", err: errors.New("ent: missing required field \"birthday\"")}
	}
	if _, ok := uc.mutation.Tel(); !ok {
		return nil, &ValidationError{Name: "tel", err: errors.New("ent: missing required field \"tel\"")}
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	u, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		u.Email = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
		u.Password = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
		u.Name = value
	}
	if value, ok := uc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAge,
		})
		u.Age = value
	}
	if value, ok := uc.mutation.Birthday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldBirthday,
		})
		u.Birthday = value
	}
	if value, ok := uc.mutation.Tel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTel,
		})
		u.Tel = value
	}
	if nodes := uc.mutation.SexIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.SexTable,
			Columns: []string{user.SexColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sex.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.TitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.TitleTable,
			Columns: []string{user.TitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: title.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return u, _spec
}
