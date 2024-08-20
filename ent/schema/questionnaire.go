package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Questionnaire holds the schema definition for the Questionnaire entity.
type Questionnaire struct {
	ent.Schema
}

// Fields of the Questionnaire.
func (Questionnaire) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.String("object"),
		field.Int("anonymous").Default(2).Comment("是否为匿名填写"),
	}
}

// Edges of the Questionnaire.
func (Questionnaire) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("include", Question.Type),
		edge.From("created_by", User.Type).Ref("create").Unique(),
	}
}

func (Questionnaire) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
