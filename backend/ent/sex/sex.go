// Code generated by entc, DO NOT EDIT.

package sex

const (
	// Label holds the string label denoting the sex type in the database.
	Label = "sex"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSexname holds the string denoting the sexname field in the database.
	FieldSexname = "sexname"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the sex in the database.
	Table = "sexes"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "sex_user"
)

// Columns holds all SQL columns for sex fields.
var Columns = []string{
	FieldID,
	FieldSexname,
}

var (
	// SexnameValidator is a validator for the "sexname" field. It is called by the builders before save.
	SexnameValidator func(string) error
)
