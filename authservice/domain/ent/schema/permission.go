package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Permission struct {
	ent.Schema
}

func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.String("permissionName").
			NotEmpty(),

		field.String("uri").
			NotEmpty(),

		field.String("method").
			NotEmpty(),
	}
}

func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permissions", Permission.Type), // real edge name used by RolePermission
	}
}
