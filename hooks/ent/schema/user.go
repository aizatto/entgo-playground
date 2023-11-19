package schema

import (
	"context"
	"errors"
	"log"

	gen "hooks/ent"
	"hooks/ent/hook"
	"hooks/ent/user"

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
		field.Enum("fail").
			Values("before", "after", "never"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		// First hook.
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					state, _ := m.Fail()
					if state == user.FailBefore {
						log.Println("hook: failed before next.Mutate()")
						return nil, errors.New("hook error: failed before next.Mutate()")
					}

					v, err := next.Mutate(ctx, m)
					if state == user.FailAfter {
						log.Println("hook: failed after next.Mutate()")
						return nil, errors.New("hook error: failed after next.Mutate()")
					}

					return v, err
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
