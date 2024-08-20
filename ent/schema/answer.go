package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Answer holds the schema definition for the Answer entity.
type Answer struct {
	ent.Schema
}

// Fields of the Answer.
func (Answer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("answer_choice"),
		field.String("answer_text"),
		field.Int64("questionnaire_id"),
		field.Int64("question_id"),
	}
}

// Edges of the Answer.
func (Answer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", Question.Type).Ref("own"),
	}
}

func (Answer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
