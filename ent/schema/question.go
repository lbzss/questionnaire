package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.String("question").Comment("问题"),
		field.Int("type").Comment("问题类型，选择或文本"),
		field.Int64("questionnaire_id").Comment("问卷ID"),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("own", Answer.Type),
		edge.From("included", Questionnaire.Type).Ref("include"),
	}
}

func (Question) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
