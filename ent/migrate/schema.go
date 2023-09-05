// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DataSourcesColumns holds the columns for the "data_sources" table.
	DataSourcesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString, Default: "unknown"},
		{Name: "status", Type: field.TypeString, Default: "idle"},
		{Name: "config", Type: field.TypeString, Nullable: true},
		{Name: "last_indexed", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
		{Name: "date_added", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
	}
	// DataSourcesTable holds the schema information for the "data_sources" table.
	DataSourcesTable = &schema.Table{
		Name:       "data_sources",
		Columns:    DataSourcesColumns,
		PrimaryKey: []*schema.Column{DataSourcesColumns[0]},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// PasswordTokensColumns holds the columns for the "password_tokens" table.
	PasswordTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hash", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "password_token_user", Type: field.TypeInt},
	}
	// PasswordTokensTable holds the schema information for the "password_tokens" table.
	PasswordTokensTable = &schema.Table{
		Name:       "password_tokens",
		Columns:    PasswordTokensColumns,
		PrimaryKey: []*schema.Column{PasswordTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "password_tokens_users_user",
				Columns:    []*schema.Column{PasswordTokensColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "verified", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// GroupUsersColumns holds the columns for the "group_users" table.
	GroupUsersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupUsersTable holds the schema information for the "group_users" table.
	GroupUsersTable = &schema.Table{
		Name:       "group_users",
		Columns:    GroupUsersColumns,
		PrimaryKey: []*schema.Column{GroupUsersColumns[0], GroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_users_group_id",
				Columns:    []*schema.Column{GroupUsersColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_users_user_id",
				Columns:    []*schema.Column{GroupUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DataSourcesTable,
		GroupsTable,
		PasswordTokensTable,
		UsersTable,
		GroupUsersTable,
	}
)

func init() {
	PasswordTokensTable.ForeignKeys[0].RefTable = UsersTable
	GroupUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[1].RefTable = UsersTable
}