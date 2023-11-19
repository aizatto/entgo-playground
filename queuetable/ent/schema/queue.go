package schema

import (
	"queuetable/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Queue holds the schema definition for the Queue entity.
type Queue struct {
	ent.Schema
}

// Fields of the Queue.
func (Queue) Fields() []ent.Field {
	return []ent.Field{
		field.String("instruction"),
		field.Enum("obj_table").
			Values(user.Table),
		field.UUID("obj_id", uuid.UUID{}).
			Unique(),
	}
}

// Edges of the Queue.
func (Queue) Edges() []ent.Edge {
	return nil
}
