package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Int("uid").
			Positive().Comment("用户ID"),
		field.String("url").
			Default("").Comment("资源地址"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return nil
}
