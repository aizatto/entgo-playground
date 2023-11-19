// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFail holds the string denoting the fail field in the database.
	FieldFail = "fail"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldFail,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "hooks/ent/runtime"
var (
	Hooks [1]ent.Hook
)

// Fail defines the type for the "fail" enum field.
type Fail string

// Fail values.
const (
	FailBefore Fail = "before"
	FailAfter  Fail = "after"
	FailNever  Fail = "never"
)

func (f Fail) String() string {
	return string(f)
}

// FailValidator is a validator for the "fail" field enum values. It is called by the builders before save.
func FailValidator(f Fail) error {
	switch f {
	case FailBefore, FailAfter, FailNever:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for fail field: %q", f)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFail orders the results by the fail field.
func ByFail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFail, opts...).ToFunc()
}
