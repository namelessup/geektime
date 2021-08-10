package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("username").
			Unique().
			Immutable(),
		field.String("password").
			Sensitive().
			Default("123456"),
		field.Time("createdTime").
			StorageKey("created_time"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
