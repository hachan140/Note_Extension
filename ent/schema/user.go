package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now()),
		field.String("created_by"),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now),
		field.String("update_by"),
		field.Time("deleted_at").Nillable(),
		field.String("fullname"),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("avatar_url").Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
