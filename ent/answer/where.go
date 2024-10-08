// Code generated by ent, DO NOT EDIT.

package answer

import (
	"questionnaire/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldUpdatedAt, v))
}

// AnswerChoice applies equality check predicate on the "answer_choice" field. It's identical to AnswerChoiceEQ.
func AnswerChoice(v int) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldAnswerChoice, v))
}

// AnswerText applies equality check predicate on the "answer_text" field. It's identical to AnswerTextEQ.
func AnswerText(v string) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldAnswerText, v))
}

// QuestionnaireID applies equality check predicate on the "questionnaire_id" field. It's identical to QuestionnaireIDEQ.
func QuestionnaireID(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldQuestionnaireID, v))
}

// QuestionID applies equality check predicate on the "question_id" field. It's identical to QuestionIDEQ.
func QuestionID(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldQuestionID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldUpdatedAt, v))
}

// AnswerChoiceEQ applies the EQ predicate on the "answer_choice" field.
func AnswerChoiceEQ(v int) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldAnswerChoice, v))
}

// AnswerChoiceNEQ applies the NEQ predicate on the "answer_choice" field.
func AnswerChoiceNEQ(v int) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldAnswerChoice, v))
}

// AnswerChoiceIn applies the In predicate on the "answer_choice" field.
func AnswerChoiceIn(vs ...int) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldAnswerChoice, vs...))
}

// AnswerChoiceNotIn applies the NotIn predicate on the "answer_choice" field.
func AnswerChoiceNotIn(vs ...int) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldAnswerChoice, vs...))
}

// AnswerChoiceGT applies the GT predicate on the "answer_choice" field.
func AnswerChoiceGT(v int) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldAnswerChoice, v))
}

// AnswerChoiceGTE applies the GTE predicate on the "answer_choice" field.
func AnswerChoiceGTE(v int) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldAnswerChoice, v))
}

// AnswerChoiceLT applies the LT predicate on the "answer_choice" field.
func AnswerChoiceLT(v int) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldAnswerChoice, v))
}

// AnswerChoiceLTE applies the LTE predicate on the "answer_choice" field.
func AnswerChoiceLTE(v int) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldAnswerChoice, v))
}

// AnswerTextEQ applies the EQ predicate on the "answer_text" field.
func AnswerTextEQ(v string) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldAnswerText, v))
}

// AnswerTextNEQ applies the NEQ predicate on the "answer_text" field.
func AnswerTextNEQ(v string) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldAnswerText, v))
}

// AnswerTextIn applies the In predicate on the "answer_text" field.
func AnswerTextIn(vs ...string) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldAnswerText, vs...))
}

// AnswerTextNotIn applies the NotIn predicate on the "answer_text" field.
func AnswerTextNotIn(vs ...string) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldAnswerText, vs...))
}

// AnswerTextGT applies the GT predicate on the "answer_text" field.
func AnswerTextGT(v string) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldAnswerText, v))
}

// AnswerTextGTE applies the GTE predicate on the "answer_text" field.
func AnswerTextGTE(v string) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldAnswerText, v))
}

// AnswerTextLT applies the LT predicate on the "answer_text" field.
func AnswerTextLT(v string) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldAnswerText, v))
}

// AnswerTextLTE applies the LTE predicate on the "answer_text" field.
func AnswerTextLTE(v string) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldAnswerText, v))
}

// AnswerTextContains applies the Contains predicate on the "answer_text" field.
func AnswerTextContains(v string) predicate.Answer {
	return predicate.Answer(sql.FieldContains(FieldAnswerText, v))
}

// AnswerTextHasPrefix applies the HasPrefix predicate on the "answer_text" field.
func AnswerTextHasPrefix(v string) predicate.Answer {
	return predicate.Answer(sql.FieldHasPrefix(FieldAnswerText, v))
}

// AnswerTextHasSuffix applies the HasSuffix predicate on the "answer_text" field.
func AnswerTextHasSuffix(v string) predicate.Answer {
	return predicate.Answer(sql.FieldHasSuffix(FieldAnswerText, v))
}

// AnswerTextEqualFold applies the EqualFold predicate on the "answer_text" field.
func AnswerTextEqualFold(v string) predicate.Answer {
	return predicate.Answer(sql.FieldEqualFold(FieldAnswerText, v))
}

// AnswerTextContainsFold applies the ContainsFold predicate on the "answer_text" field.
func AnswerTextContainsFold(v string) predicate.Answer {
	return predicate.Answer(sql.FieldContainsFold(FieldAnswerText, v))
}

// QuestionnaireIDEQ applies the EQ predicate on the "questionnaire_id" field.
func QuestionnaireIDEQ(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldQuestionnaireID, v))
}

// QuestionnaireIDNEQ applies the NEQ predicate on the "questionnaire_id" field.
func QuestionnaireIDNEQ(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldQuestionnaireID, v))
}

// QuestionnaireIDIn applies the In predicate on the "questionnaire_id" field.
func QuestionnaireIDIn(vs ...int64) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldQuestionnaireID, vs...))
}

// QuestionnaireIDNotIn applies the NotIn predicate on the "questionnaire_id" field.
func QuestionnaireIDNotIn(vs ...int64) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldQuestionnaireID, vs...))
}

// QuestionnaireIDGT applies the GT predicate on the "questionnaire_id" field.
func QuestionnaireIDGT(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldQuestionnaireID, v))
}

// QuestionnaireIDGTE applies the GTE predicate on the "questionnaire_id" field.
func QuestionnaireIDGTE(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldQuestionnaireID, v))
}

// QuestionnaireIDLT applies the LT predicate on the "questionnaire_id" field.
func QuestionnaireIDLT(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldQuestionnaireID, v))
}

// QuestionnaireIDLTE applies the LTE predicate on the "questionnaire_id" field.
func QuestionnaireIDLTE(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldQuestionnaireID, v))
}

// QuestionIDEQ applies the EQ predicate on the "question_id" field.
func QuestionIDEQ(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldQuestionID, v))
}

// QuestionIDNEQ applies the NEQ predicate on the "question_id" field.
func QuestionIDNEQ(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldQuestionID, v))
}

// QuestionIDIn applies the In predicate on the "question_id" field.
func QuestionIDIn(vs ...int64) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldQuestionID, vs...))
}

// QuestionIDNotIn applies the NotIn predicate on the "question_id" field.
func QuestionIDNotIn(vs ...int64) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldQuestionID, vs...))
}

// QuestionIDGT applies the GT predicate on the "question_id" field.
func QuestionIDGT(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldQuestionID, v))
}

// QuestionIDGTE applies the GTE predicate on the "question_id" field.
func QuestionIDGTE(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldQuestionID, v))
}

// QuestionIDLT applies the LT predicate on the "question_id" field.
func QuestionIDLT(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldQuestionID, v))
}

// QuestionIDLTE applies the LTE predicate on the "question_id" field.
func QuestionIDLTE(v int64) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldQuestionID, v))
}

// HasQuestion applies the HasEdge predicate on the "question" edge.
func HasQuestion() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, QuestionTable, QuestionPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionWith applies the HasEdge predicate on the "question" edge with a given conditions (other predicates).
func HasQuestionWith(preds ...predicate.Question) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := newQuestionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Answer) predicate.Answer {
	return predicate.Answer(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Answer) predicate.Answer {
	return predicate.Answer(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Answer) predicate.Answer {
	return predicate.Answer(sql.NotPredicates(p))
}
