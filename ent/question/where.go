// Code generated by ent, DO NOT EDIT.

package question

import (
	"questionnaire/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldUpdatedAt, v))
}

// Question applies equality check predicate on the "question" field. It's identical to QuestionEQ.
func Question(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldQuestion, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldType, v))
}

// QuestionnaireID applies equality check predicate on the "questionnaire_id" field. It's identical to QuestionnaireIDEQ.
func QuestionnaireID(v int64) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldQuestionnaireID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldUpdatedAt, v))
}

// QuestionEQ applies the EQ predicate on the "question" field.
func QuestionEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldQuestion, v))
}

// QuestionNEQ applies the NEQ predicate on the "question" field.
func QuestionNEQ(v string) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldQuestion, v))
}

// QuestionIn applies the In predicate on the "question" field.
func QuestionIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldQuestion, vs...))
}

// QuestionNotIn applies the NotIn predicate on the "question" field.
func QuestionNotIn(vs ...string) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldQuestion, vs...))
}

// QuestionGT applies the GT predicate on the "question" field.
func QuestionGT(v string) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldQuestion, v))
}

// QuestionGTE applies the GTE predicate on the "question" field.
func QuestionGTE(v string) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldQuestion, v))
}

// QuestionLT applies the LT predicate on the "question" field.
func QuestionLT(v string) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldQuestion, v))
}

// QuestionLTE applies the LTE predicate on the "question" field.
func QuestionLTE(v string) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldQuestion, v))
}

// QuestionContains applies the Contains predicate on the "question" field.
func QuestionContains(v string) predicate.Question {
	return predicate.Question(sql.FieldContains(FieldQuestion, v))
}

// QuestionHasPrefix applies the HasPrefix predicate on the "question" field.
func QuestionHasPrefix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasPrefix(FieldQuestion, v))
}

// QuestionHasSuffix applies the HasSuffix predicate on the "question" field.
func QuestionHasSuffix(v string) predicate.Question {
	return predicate.Question(sql.FieldHasSuffix(FieldQuestion, v))
}

// QuestionEqualFold applies the EqualFold predicate on the "question" field.
func QuestionEqualFold(v string) predicate.Question {
	return predicate.Question(sql.FieldEqualFold(FieldQuestion, v))
}

// QuestionContainsFold applies the ContainsFold predicate on the "question" field.
func QuestionContainsFold(v string) predicate.Question {
	return predicate.Question(sql.FieldContainsFold(FieldQuestion, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldType, v))
}

// QuestionnaireIDEQ applies the EQ predicate on the "questionnaire_id" field.
func QuestionnaireIDEQ(v int64) predicate.Question {
	return predicate.Question(sql.FieldEQ(FieldQuestionnaireID, v))
}

// QuestionnaireIDNEQ applies the NEQ predicate on the "questionnaire_id" field.
func QuestionnaireIDNEQ(v int64) predicate.Question {
	return predicate.Question(sql.FieldNEQ(FieldQuestionnaireID, v))
}

// QuestionnaireIDIn applies the In predicate on the "questionnaire_id" field.
func QuestionnaireIDIn(vs ...int64) predicate.Question {
	return predicate.Question(sql.FieldIn(FieldQuestionnaireID, vs...))
}

// QuestionnaireIDNotIn applies the NotIn predicate on the "questionnaire_id" field.
func QuestionnaireIDNotIn(vs ...int64) predicate.Question {
	return predicate.Question(sql.FieldNotIn(FieldQuestionnaireID, vs...))
}

// QuestionnaireIDGT applies the GT predicate on the "questionnaire_id" field.
func QuestionnaireIDGT(v int64) predicate.Question {
	return predicate.Question(sql.FieldGT(FieldQuestionnaireID, v))
}

// QuestionnaireIDGTE applies the GTE predicate on the "questionnaire_id" field.
func QuestionnaireIDGTE(v int64) predicate.Question {
	return predicate.Question(sql.FieldGTE(FieldQuestionnaireID, v))
}

// QuestionnaireIDLT applies the LT predicate on the "questionnaire_id" field.
func QuestionnaireIDLT(v int64) predicate.Question {
	return predicate.Question(sql.FieldLT(FieldQuestionnaireID, v))
}

// QuestionnaireIDLTE applies the LTE predicate on the "questionnaire_id" field.
func QuestionnaireIDLTE(v int64) predicate.Question {
	return predicate.Question(sql.FieldLTE(FieldQuestionnaireID, v))
}

// HasOwn applies the HasEdge predicate on the "own" edge.
func HasOwn() predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, OwnTable, OwnPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnWith applies the HasEdge predicate on the "own" edge with a given conditions (other predicates).
func HasOwnWith(preds ...predicate.Answer) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := newOwnStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIncluded applies the HasEdge predicate on the "included" edge.
func HasIncluded() predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, IncludedTable, IncludedPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIncludedWith applies the HasEdge predicate on the "included" edge with a given conditions (other predicates).
func HasIncludedWith(preds ...predicate.Questionnaire) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := newIncludedStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Question) predicate.Question {
	return predicate.Question(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Question) predicate.Question {
	return predicate.Question(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Question) predicate.Question {
	return predicate.Question(sql.NotPredicates(p))
}
