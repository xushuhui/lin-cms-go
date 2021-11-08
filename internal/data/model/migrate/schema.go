// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BookColumns holds the columns for the "book" table.
	BookColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "author", Type: field.TypeString},
		{Name: "summary", Type: field.TypeString},
		{Name: "image", Type: field.TypeString},
	}
	// BookTable holds the schema information for the "book" table.
	BookTable = &schema.Table{
		Name:       "book",
		Columns:    BookColumns,
		PrimaryKey: []*schema.Column{BookColumns[0]},
	}
	// LinFileColumns holds the columns for the "lin_file" table.
	LinFileColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "path", Type: field.TypeString},
		{Name: "type", Type: field.TypeInt8},
		{Name: "name", Type: field.TypeString},
		{Name: "extension", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt},
		{Name: "md5", Type: field.TypeString, Unique: true},
	}
	// LinFileTable holds the schema information for the "lin_file" table.
	LinFileTable = &schema.Table{
		Name:       "lin_file",
		Columns:    LinFileColumns,
		PrimaryKey: []*schema.Column{LinFileColumns[0]},
	}
	// LinGroupColumns holds the columns for the "lin_group" table.
	LinGroupColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "info", Type: field.TypeString},
		{Name: "level", Type: field.TypeInt8},
	}
	// LinGroupTable holds the schema information for the "lin_group" table.
	LinGroupTable = &schema.Table{
		Name:       "lin_group",
		Columns:    LinGroupColumns,
		PrimaryKey: []*schema.Column{LinGroupColumns[0]},
	}
	// LinLogColumns holds the columns for the "lin_log" table.
	LinLogColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime},
		{Name: "message", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "username", Type: field.TypeString},
		{Name: "status_code", Type: field.TypeInt},
		{Name: "method", Type: field.TypeString},
		{Name: "path", Type: field.TypeString},
		{Name: "permission", Type: field.TypeString},
	}
	// LinLogTable holds the schema information for the "lin_log" table.
	LinLogTable = &schema.Table{
		Name:       "lin_log",
		Columns:    LinLogColumns,
		PrimaryKey: []*schema.Column{LinLogColumns[0]},
	}
	// LinPermissionColumns holds the columns for the "lin_permission" table.
	LinPermissionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "module", Type: field.TypeString},
		{Name: "mount", Type: field.TypeInt8},
	}
	// LinPermissionTable holds the schema information for the "lin_permission" table.
	LinPermissionTable = &schema.Table{
		Name:       "lin_permission",
		Columns:    LinPermissionColumns,
		PrimaryKey: []*schema.Column{LinPermissionColumns[0]},
	}
	// LinUserColumns holds the columns for the "lin_user" table.
	LinUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "nickname", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString, Default: ""},
		{Name: "email", Type: field.TypeString, Unique: true},
	}
	// LinUserTable holds the schema information for the "lin_user" table.
	LinUserTable = &schema.Table{
		Name:       "lin_user",
		Columns:    LinUserColumns,
		PrimaryKey: []*schema.Column{LinUserColumns[0]},
	}
	// LinUserIdentiyColumns holds the columns for the "lin_user_identiy" table.
	LinUserIdentiyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "identity_type", Type: field.TypeString},
		{Name: "identifier", Type: field.TypeString},
		{Name: "credential", Type: field.TypeString},
		{Name: "lin_user_lin_user_identiy", Type: field.TypeInt, Nullable: true},
	}
	// LinUserIdentiyTable holds the schema information for the "lin_user_identiy" table.
	LinUserIdentiyTable = &schema.Table{
		Name:       "lin_user_identiy",
		Columns:    LinUserIdentiyColumns,
		PrimaryKey: []*schema.Column{LinUserIdentiyColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lin_user_identiy_lin_user_lin_user_identiy",
				Columns:    []*schema.Column{LinUserIdentiyColumns[5]},
				RefColumns: []*schema.Column{LinUserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LinUserGroupColumns holds the columns for the "lin_user_group" table.
	LinUserGroupColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// LinUserGroupTable holds the schema information for the "lin_user_group" table.
	LinUserGroupTable = &schema.Table{
		Name:       "lin_user_group",
		Columns:    LinUserGroupColumns,
		PrimaryKey: []*schema.Column{LinUserGroupColumns[0], LinUserGroupColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lin_user_group_user_id",
				Columns:    []*schema.Column{LinUserGroupColumns[0]},
				RefColumns: []*schema.Column{LinGroupColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "lin_user_group_group_id",
				Columns:    []*schema.Column{LinUserGroupColumns[1]},
				RefColumns: []*schema.Column{LinUserColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LinGroupPermissionColumns holds the columns for the "lin_group_permission" table.
	LinGroupPermissionColumns = []*schema.Column{
		{Name: "permission_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// LinGroupPermissionTable holds the schema information for the "lin_group_permission" table.
	LinGroupPermissionTable = &schema.Table{
		Name:       "lin_group_permission",
		Columns:    LinGroupPermissionColumns,
		PrimaryKey: []*schema.Column{LinGroupPermissionColumns[0], LinGroupPermissionColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "lin_group_permission_permission_id",
				Columns:    []*schema.Column{LinGroupPermissionColumns[0]},
				RefColumns: []*schema.Column{LinPermissionColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "lin_group_permission_group_id",
				Columns:    []*schema.Column{LinGroupPermissionColumns[1]},
				RefColumns: []*schema.Column{LinGroupColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BookTable,
		LinFileTable,
		LinGroupTable,
		LinLogTable,
		LinPermissionTable,
		LinUserTable,
		LinUserIdentiyTable,
		LinUserGroupTable,
		LinGroupPermissionTable,
	}
)

func init() {
	BookTable.Annotation = &entsql.Annotation{
		Table: "book",
	}
	LinFileTable.Annotation = &entsql.Annotation{
		Table: "lin_file",
	}
	LinGroupTable.Annotation = &entsql.Annotation{
		Table: "lin_group",
	}
	LinLogTable.Annotation = &entsql.Annotation{
		Table: "lin_log",
	}
	LinPermissionTable.Annotation = &entsql.Annotation{
		Table: "lin_permission",
	}
	LinUserTable.Annotation = &entsql.Annotation{
		Table: "lin_user",
	}
	LinUserIdentiyTable.ForeignKeys[0].RefTable = LinUserTable
	LinUserIdentiyTable.Annotation = &entsql.Annotation{
		Table: "lin_user_identiy",
	}
	LinUserGroupTable.ForeignKeys[0].RefTable = LinGroupTable
	LinUserGroupTable.ForeignKeys[1].RefTable = LinUserTable
	LinGroupPermissionTable.ForeignKeys[0].RefTable = LinPermissionTable
	LinGroupPermissionTable.ForeignKeys[1].RefTable = LinGroupTable
}
