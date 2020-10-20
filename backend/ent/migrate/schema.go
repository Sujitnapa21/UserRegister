// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "rolename", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:        "roles",
		Columns:     RolesColumns,
		PrimaryKey:  []*schema.Column{RolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SexesColumns holds the columns for the "sexes" table.
	SexesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "sexname", Type: field.TypeString},
	}
	// SexesTable holds the schema information for the "sexes" table.
	SexesTable = &schema.Table{
		Name:        "sexes",
		Columns:     SexesColumns,
		PrimaryKey:  []*schema.Column{SexesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TitlesColumns holds the columns for the "titles" table.
	TitlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "titlename", Type: field.TypeString},
	}
	// TitlesTable holds the schema information for the "titles" table.
	TitlesTable = &schema.Table{
		Name:        "titles",
		Columns:     TitlesColumns,
		PrimaryKey:  []*schema.Column{TitlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "age", Type: field.TypeString},
		{Name: "birthday", Type: field.TypeTime},
		{Name: "tel", Type: field.TypeString},
		{Name: "role_user", Type: field.TypeInt, Nullable: true},
		{Name: "sex_user", Type: field.TypeInt, Nullable: true},
		{Name: "title_user", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_roles_user",
				Columns: []*schema.Column{UsersColumns[7]},

				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_sexes_user",
				Columns: []*schema.Column{UsersColumns[8]},

				RefColumns: []*schema.Column{SexesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_titles_user",
				Columns: []*schema.Column{UsersColumns[9]},

				RefColumns: []*schema.Column{TitlesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RolesTable,
		SexesTable,
		TitlesTable,
		UsersTable,
	}
)

func init() {
	UsersTable.ForeignKeys[0].RefTable = RolesTable
	UsersTable.ForeignKeys[1].RefTable = SexesTable
	UsersTable.ForeignKeys[2].RefTable = TitlesTable
}
