package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		index.Fields("email", "phone"),
		// 唯一约束索引
		index.Fields("id").
			Unique(),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.String("nickname"),
		field.String("avatar").Optional(),
		field.String("password").Sensitive(),
		field.String("role").Default("user"),
		field.Bool("ban").Default(false),
		field.Int("invite_uid").Optional(),
		field.Time("login_time").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
