package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type RefreshToken struct {
	ent.Schema
}

func (RefreshToken) Fields() []ent.Field{
return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.String("username").
			Unique(),
		field.String("token"),
		field.String("expiredAt"),
		field.Enum("roleName").
			Values("ADMIN", "MANAGER", "CUSTOMER").
			Default("CUSTOMER"),
	}
}
