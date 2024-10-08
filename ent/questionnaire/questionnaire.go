// Code generated by ent, DO NOT EDIT.

package questionnaire

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the questionnaire type in the database.
	Label = "questionnaire"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldObject holds the string denoting the object field in the database.
	FieldObject = "object"
	// FieldAnonymous holds the string denoting the anonymous field in the database.
	FieldAnonymous = "anonymous"
	// EdgeInclude holds the string denoting the include edge name in mutations.
	EdgeInclude = "include"
	// EdgeCreatedBy holds the string denoting the created_by edge name in mutations.
	EdgeCreatedBy = "created_by"
	// Table holds the table name of the questionnaire in the database.
	Table = "questionnaires"
	// IncludeTable is the table that holds the include relation/edge. The primary key declared below.
	IncludeTable = "questionnaire_include"
	// IncludeInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	IncludeInverseTable = "questions"
	// CreatedByTable is the table that holds the created_by relation/edge.
	CreatedByTable = "questionnaires"
	// CreatedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatedByInverseTable = "users"
	// CreatedByColumn is the table column denoting the created_by relation/edge.
	CreatedByColumn = "user_create"
)

// Columns holds all SQL columns for questionnaire fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDescription,
	FieldObject,
	FieldAnonymous,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "questionnaires"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_create",
}

var (
	// IncludePrimaryKey and IncludeColumn2 are the table columns denoting the
	// primary key for the include relation (M2M).
	IncludePrimaryKey = []string{"questionnaire_id", "question_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultAnonymous holds the default value on creation for the "anonymous" field.
	DefaultAnonymous int
)

// OrderOption defines the ordering options for the Questionnaire queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByObject orders the results by the object field.
func ByObject(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldObject, opts...).ToFunc()
}

// ByAnonymous orders the results by the anonymous field.
func ByAnonymous(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAnonymous, opts...).ToFunc()
}

// ByIncludeCount orders the results by include count.
func ByIncludeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIncludeStep(), opts...)
	}
}

// ByInclude orders the results by include terms.
func ByInclude(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIncludeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatedByField orders the results by created_by field.
func ByCreatedByField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedByStep(), sql.OrderByField(field, opts...))
	}
}
func newIncludeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IncludeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, IncludeTable, IncludePrimaryKey...),
	)
}
func newCreatedByStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedByInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CreatedByTable, CreatedByColumn),
	)
}
