package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("birthday").Immutable(),
		field.String("name").NotEmpty(),
		field.Float("nickname"),
		field.Float("country").NotEmpty(),
		field.Float("age").Positive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "country").Unique(),
	}
}

func (User) Config() ent.Config {
	return ent.Config{
		Table: "users",
	}
}
