// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"questionnaire/ent/predicate"
	"questionnaire/ent/question"
	"questionnaire/ent/questionnaire"
	"questionnaire/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QuestionnaireUpdate is the builder for updating Questionnaire entities.
type QuestionnaireUpdate struct {
	config
	hooks    []Hook
	mutation *QuestionnaireMutation
}

// Where appends a list predicates to the QuestionnaireUpdate builder.
func (qu *QuestionnaireUpdate) Where(ps ...predicate.Questionnaire) *QuestionnaireUpdate {
	qu.mutation.Where(ps...)
	return qu
}

// SetCreatedAt sets the "created_at" field.
func (qu *QuestionnaireUpdate) SetCreatedAt(t time.Time) *QuestionnaireUpdate {
	qu.mutation.SetCreatedAt(t)
	return qu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (qu *QuestionnaireUpdate) SetNillableCreatedAt(t *time.Time) *QuestionnaireUpdate {
	if t != nil {
		qu.SetCreatedAt(*t)
	}
	return qu
}

// SetUpdatedAt sets the "updated_at" field.
func (qu *QuestionnaireUpdate) SetUpdatedAt(t time.Time) *QuestionnaireUpdate {
	qu.mutation.SetUpdatedAt(t)
	return qu
}

// SetName sets the "name" field.
func (qu *QuestionnaireUpdate) SetName(s string) *QuestionnaireUpdate {
	qu.mutation.SetName(s)
	return qu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (qu *QuestionnaireUpdate) SetNillableName(s *string) *QuestionnaireUpdate {
	if s != nil {
		qu.SetName(*s)
	}
	return qu
}

// SetDescription sets the "description" field.
func (qu *QuestionnaireUpdate) SetDescription(s string) *QuestionnaireUpdate {
	qu.mutation.SetDescription(s)
	return qu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (qu *QuestionnaireUpdate) SetNillableDescription(s *string) *QuestionnaireUpdate {
	if s != nil {
		qu.SetDescription(*s)
	}
	return qu
}

// SetObject sets the "object" field.
func (qu *QuestionnaireUpdate) SetObject(s string) *QuestionnaireUpdate {
	qu.mutation.SetObject(s)
	return qu
}

// SetNillableObject sets the "object" field if the given value is not nil.
func (qu *QuestionnaireUpdate) SetNillableObject(s *string) *QuestionnaireUpdate {
	if s != nil {
		qu.SetObject(*s)
	}
	return qu
}

// SetAnonymous sets the "anonymous" field.
func (qu *QuestionnaireUpdate) SetAnonymous(i int) *QuestionnaireUpdate {
	qu.mutation.ResetAnonymous()
	qu.mutation.SetAnonymous(i)
	return qu
}

// SetNillableAnonymous sets the "anonymous" field if the given value is not nil.
func (qu *QuestionnaireUpdate) SetNillableAnonymous(i *int) *QuestionnaireUpdate {
	if i != nil {
		qu.SetAnonymous(*i)
	}
	return qu
}

// AddAnonymous adds i to the "anonymous" field.
func (qu *QuestionnaireUpdate) AddAnonymous(i int) *QuestionnaireUpdate {
	qu.mutation.AddAnonymous(i)
	return qu
}

// AddIncludeIDs adds the "include" edge to the Question entity by IDs.
func (qu *QuestionnaireUpdate) AddIncludeIDs(ids ...int) *QuestionnaireUpdate {
	qu.mutation.AddIncludeIDs(ids...)
	return qu
}

// AddInclude adds the "include" edges to the Question entity.
func (qu *QuestionnaireUpdate) AddInclude(q ...*Question) *QuestionnaireUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qu.AddIncludeIDs(ids...)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (qu *QuestionnaireUpdate) SetCreatedByID(id int) *QuestionnaireUpdate {
	qu.mutation.SetCreatedByID(id)
	return qu
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (qu *QuestionnaireUpdate) SetNillableCreatedByID(id *int) *QuestionnaireUpdate {
	if id != nil {
		qu = qu.SetCreatedByID(*id)
	}
	return qu
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (qu *QuestionnaireUpdate) SetCreatedBy(u *User) *QuestionnaireUpdate {
	return qu.SetCreatedByID(u.ID)
}

// Mutation returns the QuestionnaireMutation object of the builder.
func (qu *QuestionnaireUpdate) Mutation() *QuestionnaireMutation {
	return qu.mutation
}

// ClearInclude clears all "include" edges to the Question entity.
func (qu *QuestionnaireUpdate) ClearInclude() *QuestionnaireUpdate {
	qu.mutation.ClearInclude()
	return qu
}

// RemoveIncludeIDs removes the "include" edge to Question entities by IDs.
func (qu *QuestionnaireUpdate) RemoveIncludeIDs(ids ...int) *QuestionnaireUpdate {
	qu.mutation.RemoveIncludeIDs(ids...)
	return qu
}

// RemoveInclude removes "include" edges to Question entities.
func (qu *QuestionnaireUpdate) RemoveInclude(q ...*Question) *QuestionnaireUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qu.RemoveIncludeIDs(ids...)
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (qu *QuestionnaireUpdate) ClearCreatedBy() *QuestionnaireUpdate {
	qu.mutation.ClearCreatedBy()
	return qu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qu *QuestionnaireUpdate) Save(ctx context.Context) (int, error) {
	qu.defaults()
	return withHooks(ctx, qu.sqlSave, qu.mutation, qu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (qu *QuestionnaireUpdate) SaveX(ctx context.Context) int {
	affected, err := qu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qu *QuestionnaireUpdate) Exec(ctx context.Context) error {
	_, err := qu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qu *QuestionnaireUpdate) ExecX(ctx context.Context) {
	if err := qu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qu *QuestionnaireUpdate) defaults() {
	if _, ok := qu.mutation.UpdatedAt(); !ok {
		v := questionnaire.UpdateDefaultUpdatedAt()
		qu.mutation.SetUpdatedAt(v)
	}
}

func (qu *QuestionnaireUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(questionnaire.Table, questionnaire.Columns, sqlgraph.NewFieldSpec(questionnaire.FieldID, field.TypeInt))
	if ps := qu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qu.mutation.CreatedAt(); ok {
		_spec.SetField(questionnaire.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := qu.mutation.UpdatedAt(); ok {
		_spec.SetField(questionnaire.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := qu.mutation.Name(); ok {
		_spec.SetField(questionnaire.FieldName, field.TypeString, value)
	}
	if value, ok := qu.mutation.Description(); ok {
		_spec.SetField(questionnaire.FieldDescription, field.TypeString, value)
	}
	if value, ok := qu.mutation.Object(); ok {
		_spec.SetField(questionnaire.FieldObject, field.TypeString, value)
	}
	if value, ok := qu.mutation.Anonymous(); ok {
		_spec.SetField(questionnaire.FieldAnonymous, field.TypeInt, value)
	}
	if value, ok := qu.mutation.AddedAnonymous(); ok {
		_spec.AddField(questionnaire.FieldAnonymous, field.TypeInt, value)
	}
	if qu.mutation.IncludeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   questionnaire.IncludeTable,
			Columns: questionnaire.IncludePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.RemovedIncludeIDs(); len(nodes) > 0 && !qu.mutation.IncludeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   questionnaire.IncludeTable,
			Columns: questionnaire.IncludePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.IncludeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   questionnaire.IncludeTable,
			Columns: questionnaire.IncludePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qu.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionnaire.CreatedByTable,
			Columns: []string{questionnaire.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qu.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionnaire.CreatedByTable,
			Columns: []string{questionnaire.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questionnaire.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	qu.mutation.done = true
	return n, nil
}

// QuestionnaireUpdateOne is the builder for updating a single Questionnaire entity.
type QuestionnaireUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *QuestionnaireMutation
}

// SetCreatedAt sets the "created_at" field.
func (quo *QuestionnaireUpdateOne) SetCreatedAt(t time.Time) *QuestionnaireUpdateOne {
	quo.mutation.SetCreatedAt(t)
	return quo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (quo *QuestionnaireUpdateOne) SetNillableCreatedAt(t *time.Time) *QuestionnaireUpdateOne {
	if t != nil {
		quo.SetCreatedAt(*t)
	}
	return quo
}

// SetUpdatedAt sets the "updated_at" field.
func (quo *QuestionnaireUpdateOne) SetUpdatedAt(t time.Time) *QuestionnaireUpdateOne {
	quo.mutation.SetUpdatedAt(t)
	return quo
}

// SetName sets the "name" field.
func (quo *QuestionnaireUpdateOne) SetName(s string) *QuestionnaireUpdateOne {
	quo.mutation.SetName(s)
	return quo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (quo *QuestionnaireUpdateOne) SetNillableName(s *string) *QuestionnaireUpdateOne {
	if s != nil {
		quo.SetName(*s)
	}
	return quo
}

// SetDescription sets the "description" field.
func (quo *QuestionnaireUpdateOne) SetDescription(s string) *QuestionnaireUpdateOne {
	quo.mutation.SetDescription(s)
	return quo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (quo *QuestionnaireUpdateOne) SetNillableDescription(s *string) *QuestionnaireUpdateOne {
	if s != nil {
		quo.SetDescription(*s)
	}
	return quo
}

// SetObject sets the "object" field.
func (quo *QuestionnaireUpdateOne) SetObject(s string) *QuestionnaireUpdateOne {
	quo.mutation.SetObject(s)
	return quo
}

// SetNillableObject sets the "object" field if the given value is not nil.
func (quo *QuestionnaireUpdateOne) SetNillableObject(s *string) *QuestionnaireUpdateOne {
	if s != nil {
		quo.SetObject(*s)
	}
	return quo
}

// SetAnonymous sets the "anonymous" field.
func (quo *QuestionnaireUpdateOne) SetAnonymous(i int) *QuestionnaireUpdateOne {
	quo.mutation.ResetAnonymous()
	quo.mutation.SetAnonymous(i)
	return quo
}

// SetNillableAnonymous sets the "anonymous" field if the given value is not nil.
func (quo *QuestionnaireUpdateOne) SetNillableAnonymous(i *int) *QuestionnaireUpdateOne {
	if i != nil {
		quo.SetAnonymous(*i)
	}
	return quo
}

// AddAnonymous adds i to the "anonymous" field.
func (quo *QuestionnaireUpdateOne) AddAnonymous(i int) *QuestionnaireUpdateOne {
	quo.mutation.AddAnonymous(i)
	return quo
}

// AddIncludeIDs adds the "include" edge to the Question entity by IDs.
func (quo *QuestionnaireUpdateOne) AddIncludeIDs(ids ...int) *QuestionnaireUpdateOne {
	quo.mutation.AddIncludeIDs(ids...)
	return quo
}

// AddInclude adds the "include" edges to the Question entity.
func (quo *QuestionnaireUpdateOne) AddInclude(q ...*Question) *QuestionnaireUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return quo.AddIncludeIDs(ids...)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (quo *QuestionnaireUpdateOne) SetCreatedByID(id int) *QuestionnaireUpdateOne {
	quo.mutation.SetCreatedByID(id)
	return quo
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (quo *QuestionnaireUpdateOne) SetNillableCreatedByID(id *int) *QuestionnaireUpdateOne {
	if id != nil {
		quo = quo.SetCreatedByID(*id)
	}
	return quo
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (quo *QuestionnaireUpdateOne) SetCreatedBy(u *User) *QuestionnaireUpdateOne {
	return quo.SetCreatedByID(u.ID)
}

// Mutation returns the QuestionnaireMutation object of the builder.
func (quo *QuestionnaireUpdateOne) Mutation() *QuestionnaireMutation {
	return quo.mutation
}

// ClearInclude clears all "include" edges to the Question entity.
func (quo *QuestionnaireUpdateOne) ClearInclude() *QuestionnaireUpdateOne {
	quo.mutation.ClearInclude()
	return quo
}

// RemoveIncludeIDs removes the "include" edge to Question entities by IDs.
func (quo *QuestionnaireUpdateOne) RemoveIncludeIDs(ids ...int) *QuestionnaireUpdateOne {
	quo.mutation.RemoveIncludeIDs(ids...)
	return quo
}

// RemoveInclude removes "include" edges to Question entities.
func (quo *QuestionnaireUpdateOne) RemoveInclude(q ...*Question) *QuestionnaireUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return quo.RemoveIncludeIDs(ids...)
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (quo *QuestionnaireUpdateOne) ClearCreatedBy() *QuestionnaireUpdateOne {
	quo.mutation.ClearCreatedBy()
	return quo
}

// Where appends a list predicates to the QuestionnaireUpdate builder.
func (quo *QuestionnaireUpdateOne) Where(ps ...predicate.Questionnaire) *QuestionnaireUpdateOne {
	quo.mutation.Where(ps...)
	return quo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (quo *QuestionnaireUpdateOne) Select(field string, fields ...string) *QuestionnaireUpdateOne {
	quo.fields = append([]string{field}, fields...)
	return quo
}

// Save executes the query and returns the updated Questionnaire entity.
func (quo *QuestionnaireUpdateOne) Save(ctx context.Context) (*Questionnaire, error) {
	quo.defaults()
	return withHooks(ctx, quo.sqlSave, quo.mutation, quo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (quo *QuestionnaireUpdateOne) SaveX(ctx context.Context) *Questionnaire {
	node, err := quo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (quo *QuestionnaireUpdateOne) Exec(ctx context.Context) error {
	_, err := quo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (quo *QuestionnaireUpdateOne) ExecX(ctx context.Context) {
	if err := quo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (quo *QuestionnaireUpdateOne) defaults() {
	if _, ok := quo.mutation.UpdatedAt(); !ok {
		v := questionnaire.UpdateDefaultUpdatedAt()
		quo.mutation.SetUpdatedAt(v)
	}
}

func (quo *QuestionnaireUpdateOne) sqlSave(ctx context.Context) (_node *Questionnaire, err error) {
	_spec := sqlgraph.NewUpdateSpec(questionnaire.Table, questionnaire.Columns, sqlgraph.NewFieldSpec(questionnaire.FieldID, field.TypeInt))
	id, ok := quo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Questionnaire.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := quo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, questionnaire.FieldID)
		for _, f := range fields {
			if !questionnaire.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != questionnaire.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := quo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := quo.mutation.CreatedAt(); ok {
		_spec.SetField(questionnaire.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := quo.mutation.UpdatedAt(); ok {
		_spec.SetField(questionnaire.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := quo.mutation.Name(); ok {
		_spec.SetField(questionnaire.FieldName, field.TypeString, value)
	}
	if value, ok := quo.mutation.Description(); ok {
		_spec.SetField(questionnaire.FieldDescription, field.TypeString, value)
	}
	if value, ok := quo.mutation.Object(); ok {
		_spec.SetField(questionnaire.FieldObject, field.TypeString, value)
	}
	if value, ok := quo.mutation.Anonymous(); ok {
		_spec.SetField(questionnaire.FieldAnonymous, field.TypeInt, value)
	}
	if value, ok := quo.mutation.AddedAnonymous(); ok {
		_spec.AddField(questionnaire.FieldAnonymous, field.TypeInt, value)
	}
	if quo.mutation.IncludeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   questionnaire.IncludeTable,
			Columns: questionnaire.IncludePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.RemovedIncludeIDs(); len(nodes) > 0 && !quo.mutation.IncludeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   questionnaire.IncludeTable,
			Columns: questionnaire.IncludePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.IncludeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   questionnaire.IncludeTable,
			Columns: questionnaire.IncludePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if quo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionnaire.CreatedByTable,
			Columns: []string{questionnaire.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := quo.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionnaire.CreatedByTable,
			Columns: []string{questionnaire.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Questionnaire{config: quo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, quo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questionnaire.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	quo.mutation.done = true
	return _node, nil
}
