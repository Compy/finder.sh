package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// DataSource holds the schema definition for the DataSource entity.
type DataSource struct {
	ent.Schema
}

// Fields of the DataSource.
func (DataSource) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("type").Default("unknown"),
		field.String("status").Default("idle"),
		field.String("config").Optional(),
		field.Time("last_indexed").Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),
		field.Time("date_added").Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),
	}
}

// Edges of the DataSource.
func (DataSource) Edges() []ent.Edge {
	return nil
}
