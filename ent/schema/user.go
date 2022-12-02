// schema/user.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(30)",
			}),
		field.String("last_name").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(30)",
			}),
		field.String("email").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(30)",
			}).
			Optional(),
		field.Int("age").
			SchemaType(map[string]string{
				dialect.Postgres: "int",
			}).
			Optional(),
		field.Int("company_id"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("company", Company.Type).
			Unique().
			Field("company_id").
			Required(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("first_name", "last_name"),
	}
}
