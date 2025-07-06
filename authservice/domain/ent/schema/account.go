package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Account struct {
	ent.Schema
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.String("username").
			Unique(),
		field.String("password").
			Sensitive(),
		field.String("email").
			Unique(),
		field.String("phoneNumber"),
		field.Enum("oAuthType").
			Values("GOOGLE", "FACEBOOK", "GITHUB").
			Default("GOOGLE"),
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).
			Ref("accounts").
			Unique().
			Required(),
	}
}
