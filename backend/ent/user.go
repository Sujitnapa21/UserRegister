// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/Sujitnapa21/app/ent/role"
	"github.com/Sujitnapa21/app/ent/sex"
	"github.com/Sujitnapa21/app/ent/title"
	"github.com/Sujitnapa21/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age string `json:"age,omitempty"`
	// Birthday holds the value of the "birthday" field.
	Birthday time.Time `json:"birthday,omitempty"`
	// Tel holds the value of the "tel" field.
	Tel string `json:"tel,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges      UserEdges `json:"edges"`
	role_user  *int
	sex_user   *int
	title_user *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Sex holds the value of the sex edge.
	Sex *Sex
	// Role holds the value of the role edge.
	Role *Role
	// Title holds the value of the title edge.
	Title *Title
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SexOrErr returns the Sex value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SexOrErr() (*Sex, error) {
	if e.loadedTypes[0] {
		if e.Sex == nil {
			// The edge sex was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: sex.Label}
		}
		return e.Sex, nil
	}
	return nil, &NotLoadedError{edge: "sex"}
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) RoleOrErr() (*Role, error) {
	if e.loadedTypes[1] {
		if e.Role == nil {
			// The edge role was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// TitleOrErr returns the Title value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) TitleOrErr() (*Title, error) {
	if e.loadedTypes[2] {
		if e.Title == nil {
			// The edge title was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: title.Label}
		}
		return e.Title, nil
	}
	return nil, &NotLoadedError{edge: "title"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // email
		&sql.NullString{}, // password
		&sql.NullString{}, // name
		&sql.NullString{}, // age
		&sql.NullTime{},   // birthday
		&sql.NullString{}, // tel
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // role_user
		&sql.NullInt64{}, // sex_user
		&sql.NullInt64{}, // title_user
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[0])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[1])
	} else if value.Valid {
		u.Password = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[3])
	} else if value.Valid {
		u.Age = value.String
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field birthday", values[4])
	} else if value.Valid {
		u.Birthday = value.Time
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field tel", values[5])
	} else if value.Valid {
		u.Tel = value.String
	}
	values = values[6:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field role_user", value)
		} else if value.Valid {
			u.role_user = new(int)
			*u.role_user = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field sex_user", value)
		} else if value.Valid {
			u.sex_user = new(int)
			*u.sex_user = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field title_user", value)
		} else if value.Valid {
			u.title_user = new(int)
			*u.title_user = int(value.Int64)
		}
	}
	return nil
}

// QuerySex queries the sex edge of the User.
func (u *User) QuerySex() *SexQuery {
	return (&UserClient{config: u.config}).QuerySex(u)
}

// QueryRole queries the role edge of the User.
func (u *User) QueryRole() *RoleQuery {
	return (&UserClient{config: u.config}).QueryRole(u)
}

// QueryTitle queries the title edge of the User.
func (u *User) QueryTitle() *TitleQuery {
	return (&UserClient{config: u.config}).QueryTitle(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", age=")
	builder.WriteString(u.Age)
	builder.WriteString(", birthday=")
	builder.WriteString(u.Birthday.Format(time.ANSIC))
	builder.WriteString(", tel=")
	builder.WriteString(u.Tel)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
