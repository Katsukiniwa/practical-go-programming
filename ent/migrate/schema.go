// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompaniesColumns holds the columns for the "companies" table.
	CompaniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(30)"}},
	}
	// CompaniesTable holds the schema information for the "companies" table.
	CompaniesTable = &schema.Table{
		Name:       "companies",
		Columns:    CompaniesColumns,
		PrimaryKey: []*schema.Column{CompaniesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "first_name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(30)"}},
		{Name: "last_name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(30)"}},
		{Name: "email", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(30)"}},
		{Name: "age", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"postgres": "int"}},
		{Name: "company_id", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_companies_company",
				Columns:    []*schema.Column{UsersColumns[5]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_first_name_last_name",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[1], UsersColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompaniesTable,
		UsersTable,
	}
)

func init() {
	UsersTable.ForeignKeys[0].RefTable = CompaniesTable
}
