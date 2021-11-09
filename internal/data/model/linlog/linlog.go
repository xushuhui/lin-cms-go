// Code generated by entc, DO NOT EDIT.

package linlog

import (
	"time"
)

const (
	// Label holds the string label denoting the linlog type in the database.
	Label = "lin_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldStatusCode holds the string denoting the status_code field in the database.
	FieldStatusCode = "status_code"
	// FieldMethod holds the string denoting the method field in the database.
	FieldMethod = "method"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldPermission holds the string denoting the permission field in the database.
	FieldPermission = "permission"
	// Table holds the table name of the linlog in the database.
	Table = "lin_log"
)

// Columns holds all SQL columns for linlog fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeleteTime,
	FieldMessage,
	FieldUserID,
	FieldUsername,
	FieldStatusCode,
	FieldMethod,
	FieldPath,
	FieldPermission,
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
