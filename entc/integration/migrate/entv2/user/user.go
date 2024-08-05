// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldMixedString holds the string denoting the mixed_string field in the database.
	FieldMixedString = "mixed_string"
	// FieldMixedEnum holds the string denoting the mixed_enum field in the database.
	FieldMixedEnum = "mixed_enum"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldBuffer holds the string denoting the buffer field in the database.
	FieldBuffer = "buffer"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldNewName holds the string denoting the new_name field in the database.
	FieldNewName = "renamed"
	// FieldNewToken holds the string denoting the new_token field in the database.
	FieldNewToken = "new_token"
	// FieldBlob holds the string denoting the blob field in the database.
	FieldBlob = "blob"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldWorkplace holds the string denoting the workplace field in the database.
	FieldWorkplace = "workplace"
	// FieldRoles holds the string denoting the roles field in the database.
	FieldRoles = "roles"
	// FieldDefaultExpr holds the string denoting the default_expr field in the database.
	FieldDefaultExpr = "default_expr"
	// FieldDefaultExprs holds the string denoting the default_exprs field in the database.
	FieldDefaultExprs = "default_exprs"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDropOptional holds the string denoting the drop_optional field in the database.
	FieldDropOptional = "drop_optional"
	// EdgeCar holds the string denoting the car edge name in mutations.
	EdgeCar = "car"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// CarFieldID holds the string denoting the ID field of the Car.
	CarFieldID = "id"
	// PetFieldID holds the string denoting the ID field of the Pet.
	PetFieldID = "id"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CarTable is the table that holds the car relation/edge.
	CarTable = "Car"
	// CarInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarInverseTable = "Car"
	// CarColumn is the table column denoting the car relation/edge.
	CarColumn = "user_car"
	// PetsTable is the table that holds the pets relation/edge.
	PetsTable = "pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "owner_id"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "friends"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldMixedString,
	FieldMixedEnum,
	FieldActive,
	FieldAge,
	FieldName,
	FieldDescription,
	FieldNickname,
	FieldPhone,
	FieldBuffer,
	FieldTitle,
	FieldNewName,
	FieldNewToken,
	FieldBlob,
	FieldState,
	FieldStatus,
	FieldWorkplace,
	FieldRoles,
	FieldDefaultExpr,
	FieldDefaultExprs,
	FieldCreatedAt,
	FieldDropOptional,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"blog_admins",
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user", "friend"}
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
	// DefaultMixedString holds the default value on creation for the "mixed_string" field.
	DefaultMixedString string
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// DefaultPhone holds the default value on creation for the "phone" field.
	DefaultPhone string
	// DefaultBuffer holds the default value on creation for the "buffer" field.
	DefaultBuffer func() []byte
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultNewToken holds the default value on creation for the "new_token" field.
	DefaultNewToken func() string
	// BlobValidator is a validator for the "blob" field. It is called by the builders before save.
	BlobValidator func([]byte) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultDropOptional holds the default value on creation for the "drop_optional" field.
	DefaultDropOptional func() string
)

// MixedEnum defines the type for the "mixed_enum" enum field.
type MixedEnum string

// MixedEnumOn is the default value of the MixedEnum enum.
const DefaultMixedEnum = MixedEnumOn

// MixedEnum values.
const (
	MixedEnumOn  MixedEnum = "on"
	MixedEnumOff MixedEnum = "off"
)

func (me MixedEnum) String() string {
	return string(me)
}

// MixedEnumValidator is a validator for the "mixed_enum" field enum values. It is called by the builders before save.
func MixedEnumValidator(me MixedEnum) error {
	switch me {
	case MixedEnumOn, MixedEnumOff:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for mixed_enum field: %q", me)
	}
}

// State defines the type for the "state" enum field.
type State string

// StateLoggedIn is the default value of the State enum.
const DefaultState = StateLoggedIn

// State values.
const (
	StateLoggedIn  State = "logged_in"
	StateLoggedOut State = "logged_out"
	StateOnline    State = "online"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateLoggedIn, StateLoggedOut, StateOnline:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for state field: %q", s)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusDone    Status = "done"
	StatusPending Status = "pending"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusDone, StatusPending:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMixedString orders the results by the mixed_string field.
func ByMixedString(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMixedString, opts...).ToFunc()
}

// ByMixedEnum orders the results by the mixed_enum field.
func ByMixedEnum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMixedEnum, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByNewName orders the results by the new_name field.
func ByNewName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNewName, opts...).ToFunc()
}

// ByNewToken orders the results by the new_token field.
func ByNewToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNewToken, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByWorkplace orders the results by the workplace field.
func ByWorkplace(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkplace, opts...).ToFunc()
}

// ByDefaultExpr orders the results by the default_expr field.
func ByDefaultExpr(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultExpr, opts...).ToFunc()
}

// ByDefaultExprs orders the results by the default_exprs field.
func ByDefaultExprs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultExprs, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByDropOptional orders the results by the drop_optional field.
func ByDropOptional(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDropOptional, opts...).ToFunc()
}

// ByCarCount orders the results by car count.
func ByCarCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCarStep(), opts...)
	}
}

// ByCar orders the results by car terms.
func ByCar(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCarStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPetsField orders the results by pets field.
func ByPetsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPetsStep(), sql.OrderByField(field, opts...))
	}
}

// ByFriendsCount orders the results by friends count.
func ByFriendsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFriendsStep(), opts...)
	}
}

// ByFriends orders the results by friends terms.
func ByFriends(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFriendsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCarStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CarInverseTable, CarFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CarTable, CarColumn),
	)
}
func newPetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PetsInverseTable, PetFieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, PetsTable, PetsColumn),
	)
}
func newFriendsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, FriendsTable, FriendsPrimaryKey...),
	)
}
