// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"questionnaire/ent/answer"
	"questionnaire/ent/predicate"
	"questionnaire/ent/question"
	"questionnaire/ent/questionnaire"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QuestionQuery is the builder for querying Question entities.
type QuestionQuery struct {
	config
	ctx          *QueryContext
	order        []question.OrderOption
	inters       []Interceptor
	predicates   []predicate.Question
	withOwn      *AnswerQuery
	withIncluded *QuestionnaireQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the QuestionQuery builder.
func (qq *QuestionQuery) Where(ps ...predicate.Question) *QuestionQuery {
	qq.predicates = append(qq.predicates, ps...)
	return qq
}

// Limit the number of records to be returned by this query.
func (qq *QuestionQuery) Limit(limit int) *QuestionQuery {
	qq.ctx.Limit = &limit
	return qq
}

// Offset to start from.
func (qq *QuestionQuery) Offset(offset int) *QuestionQuery {
	qq.ctx.Offset = &offset
	return qq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (qq *QuestionQuery) Unique(unique bool) *QuestionQuery {
	qq.ctx.Unique = &unique
	return qq
}

// Order specifies how the records should be ordered.
func (qq *QuestionQuery) Order(o ...question.OrderOption) *QuestionQuery {
	qq.order = append(qq.order, o...)
	return qq
}

// QueryOwn chains the current query on the "own" edge.
func (qq *QuestionQuery) QueryOwn() *AnswerQuery {
	query := (&AnswerClient{config: qq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := qq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(question.Table, question.FieldID, selector),
			sqlgraph.To(answer.Table, answer.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, question.OwnTable, question.OwnPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(qq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryIncluded chains the current query on the "included" edge.
func (qq *QuestionQuery) QueryIncluded() *QuestionnaireQuery {
	query := (&QuestionnaireClient{config: qq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := qq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(question.Table, question.FieldID, selector),
			sqlgraph.To(questionnaire.Table, questionnaire.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, question.IncludedTable, question.IncludedPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(qq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Question entity from the query.
// Returns a *NotFoundError when no Question was found.
func (qq *QuestionQuery) First(ctx context.Context) (*Question, error) {
	nodes, err := qq.Limit(1).All(setContextOp(ctx, qq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{question.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qq *QuestionQuery) FirstX(ctx context.Context) *Question {
	node, err := qq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Question ID from the query.
// Returns a *NotFoundError when no Question ID was found.
func (qq *QuestionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(1).IDs(setContextOp(ctx, qq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{question.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qq *QuestionQuery) FirstIDX(ctx context.Context) int {
	id, err := qq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Question entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Question entity is found.
// Returns a *NotFoundError when no Question entities are found.
func (qq *QuestionQuery) Only(ctx context.Context) (*Question, error) {
	nodes, err := qq.Limit(2).All(setContextOp(ctx, qq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{question.Label}
	default:
		return nil, &NotSingularError{question.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qq *QuestionQuery) OnlyX(ctx context.Context) *Question {
	node, err := qq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Question ID in the query.
// Returns a *NotSingularError when more than one Question ID is found.
// Returns a *NotFoundError when no entities are found.
func (qq *QuestionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(2).IDs(setContextOp(ctx, qq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{question.Label}
	default:
		err = &NotSingularError{question.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qq *QuestionQuery) OnlyIDX(ctx context.Context) int {
	id, err := qq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Questions.
func (qq *QuestionQuery) All(ctx context.Context) ([]*Question, error) {
	ctx = setContextOp(ctx, qq.ctx, ent.OpQueryAll)
	if err := qq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Question, *QuestionQuery]()
	return withInterceptors[[]*Question](ctx, qq, qr, qq.inters)
}

// AllX is like All, but panics if an error occurs.
func (qq *QuestionQuery) AllX(ctx context.Context) []*Question {
	nodes, err := qq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Question IDs.
func (qq *QuestionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if qq.ctx.Unique == nil && qq.path != nil {
		qq.Unique(true)
	}
	ctx = setContextOp(ctx, qq.ctx, ent.OpQueryIDs)
	if err = qq.Select(question.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qq *QuestionQuery) IDsX(ctx context.Context) []int {
	ids, err := qq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qq *QuestionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, qq.ctx, ent.OpQueryCount)
	if err := qq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, qq, querierCount[*QuestionQuery](), qq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (qq *QuestionQuery) CountX(ctx context.Context) int {
	count, err := qq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qq *QuestionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, qq.ctx, ent.OpQueryExist)
	switch _, err := qq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (qq *QuestionQuery) ExistX(ctx context.Context) bool {
	exist, err := qq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the QuestionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qq *QuestionQuery) Clone() *QuestionQuery {
	if qq == nil {
		return nil
	}
	return &QuestionQuery{
		config:       qq.config,
		ctx:          qq.ctx.Clone(),
		order:        append([]question.OrderOption{}, qq.order...),
		inters:       append([]Interceptor{}, qq.inters...),
		predicates:   append([]predicate.Question{}, qq.predicates...),
		withOwn:      qq.withOwn.Clone(),
		withIncluded: qq.withIncluded.Clone(),
		// clone intermediate query.
		sql:  qq.sql.Clone(),
		path: qq.path,
	}
}

// WithOwn tells the query-builder to eager-load the nodes that are connected to
// the "own" edge. The optional arguments are used to configure the query builder of the edge.
func (qq *QuestionQuery) WithOwn(opts ...func(*AnswerQuery)) *QuestionQuery {
	query := (&AnswerClient{config: qq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	qq.withOwn = query
	return qq
}

// WithIncluded tells the query-builder to eager-load the nodes that are connected to
// the "included" edge. The optional arguments are used to configure the query builder of the edge.
func (qq *QuestionQuery) WithIncluded(opts ...func(*QuestionnaireQuery)) *QuestionQuery {
	query := (&QuestionnaireClient{config: qq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	qq.withIncluded = query
	return qq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Question.Query().
//		GroupBy(question.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (qq *QuestionQuery) GroupBy(field string, fields ...string) *QuestionGroupBy {
	qq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &QuestionGroupBy{build: qq}
	grbuild.flds = &qq.ctx.Fields
	grbuild.label = question.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Question.Query().
//		Select(question.FieldCreatedAt).
//		Scan(ctx, &v)
func (qq *QuestionQuery) Select(fields ...string) *QuestionSelect {
	qq.ctx.Fields = append(qq.ctx.Fields, fields...)
	sbuild := &QuestionSelect{QuestionQuery: qq}
	sbuild.label = question.Label
	sbuild.flds, sbuild.scan = &qq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a QuestionSelect configured with the given aggregations.
func (qq *QuestionQuery) Aggregate(fns ...AggregateFunc) *QuestionSelect {
	return qq.Select().Aggregate(fns...)
}

func (qq *QuestionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range qq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, qq); err != nil {
				return err
			}
		}
	}
	for _, f := range qq.ctx.Fields {
		if !question.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if qq.path != nil {
		prev, err := qq.path(ctx)
		if err != nil {
			return err
		}
		qq.sql = prev
	}
	return nil
}

func (qq *QuestionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Question, error) {
	var (
		nodes       = []*Question{}
		_spec       = qq.querySpec()
		loadedTypes = [2]bool{
			qq.withOwn != nil,
			qq.withIncluded != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Question).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Question{config: qq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, qq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := qq.withOwn; query != nil {
		if err := qq.loadOwn(ctx, query, nodes,
			func(n *Question) { n.Edges.Own = []*Answer{} },
			func(n *Question, e *Answer) { n.Edges.Own = append(n.Edges.Own, e) }); err != nil {
			return nil, err
		}
	}
	if query := qq.withIncluded; query != nil {
		if err := qq.loadIncluded(ctx, query, nodes,
			func(n *Question) { n.Edges.Included = []*Questionnaire{} },
			func(n *Question, e *Questionnaire) { n.Edges.Included = append(n.Edges.Included, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (qq *QuestionQuery) loadOwn(ctx context.Context, query *AnswerQuery, nodes []*Question, init func(*Question), assign func(*Question, *Answer)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Question)
	nids := make(map[int]map[*Question]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(question.OwnTable)
		s.Join(joinT).On(s.C(answer.FieldID), joinT.C(question.OwnPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(question.OwnPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(question.OwnPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Question]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Answer](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "own" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (qq *QuestionQuery) loadIncluded(ctx context.Context, query *QuestionnaireQuery, nodes []*Question, init func(*Question), assign func(*Question, *Questionnaire)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Question)
	nids := make(map[int]map[*Question]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(question.IncludedTable)
		s.Join(joinT).On(s.C(questionnaire.FieldID), joinT.C(question.IncludedPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(question.IncludedPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(question.IncludedPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Question]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Questionnaire](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "included" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (qq *QuestionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qq.querySpec()
	_spec.Node.Columns = qq.ctx.Fields
	if len(qq.ctx.Fields) > 0 {
		_spec.Unique = qq.ctx.Unique != nil && *qq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, qq.driver, _spec)
}

func (qq *QuestionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(question.Table, question.Columns, sqlgraph.NewFieldSpec(question.FieldID, field.TypeInt))
	_spec.From = qq.sql
	if unique := qq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if qq.path != nil {
		_spec.Unique = true
	}
	if fields := qq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, question.FieldID)
		for i := range fields {
			if fields[i] != question.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := qq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qq *QuestionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(qq.driver.Dialect())
	t1 := builder.Table(question.Table)
	columns := qq.ctx.Fields
	if len(columns) == 0 {
		columns = question.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if qq.sql != nil {
		selector = qq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if qq.ctx.Unique != nil && *qq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range qq.predicates {
		p(selector)
	}
	for _, p := range qq.order {
		p(selector)
	}
	if offset := qq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QuestionGroupBy is the group-by builder for Question entities.
type QuestionGroupBy struct {
	selector
	build *QuestionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qgb *QuestionGroupBy) Aggregate(fns ...AggregateFunc) *QuestionGroupBy {
	qgb.fns = append(qgb.fns, fns...)
	return qgb
}

// Scan applies the selector query and scans the result into the given value.
func (qgb *QuestionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qgb.build.ctx, ent.OpQueryGroupBy)
	if err := qgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QuestionQuery, *QuestionGroupBy](ctx, qgb.build, qgb, qgb.build.inters, v)
}

func (qgb *QuestionGroupBy) sqlScan(ctx context.Context, root *QuestionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(qgb.fns))
	for _, fn := range qgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*qgb.flds)+len(qgb.fns))
		for _, f := range *qgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*qgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// QuestionSelect is the builder for selecting fields of Question entities.
type QuestionSelect struct {
	*QuestionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (qs *QuestionSelect) Aggregate(fns ...AggregateFunc) *QuestionSelect {
	qs.fns = append(qs.fns, fns...)
	return qs
}

// Scan applies the selector query and scans the result into the given value.
func (qs *QuestionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qs.ctx, ent.OpQuerySelect)
	if err := qs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QuestionQuery, *QuestionSelect](ctx, qs.QuestionQuery, qs, qs.inters, v)
}

func (qs *QuestionSelect) sqlScan(ctx context.Context, root *QuestionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(qs.fns))
	for _, fn := range qs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*qs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
