// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package group

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldID, id))
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldActive, v))
}

// Expire applies equality check predicate on the "expire" field. It's identical to ExpireEQ.
func Expire(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldExpire, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldType, v))
}

// MaxUsers applies equality check predicate on the "max_users" field. It's identical to MaxUsersEQ.
func MaxUsers(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldMaxUsers, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldName, v))
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldActive, v))
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldActive, v))
}

// ExpireEQ applies the EQ predicate on the "expire" field.
func ExpireEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldExpire, v))
}

// ExpireNEQ applies the NEQ predicate on the "expire" field.
func ExpireNEQ(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldExpire, v))
}

// ExpireIn applies the In predicate on the "expire" field.
func ExpireIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldExpire, vs...))
}

// ExpireNotIn applies the NotIn predicate on the "expire" field.
func ExpireNotIn(vs ...time.Time) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldExpire, vs...))
}

// ExpireGT applies the GT predicate on the "expire" field.
func ExpireGT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldExpire, v))
}

// ExpireGTE applies the GTE predicate on the "expire" field.
func ExpireGTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldExpire, v))
}

// ExpireLT applies the LT predicate on the "expire" field.
func ExpireLT(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldExpire, v))
}

// ExpireLTE applies the LTE predicate on the "expire" field.
func ExpireLTE(v time.Time) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldExpire, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Group {
	return predicate.Group(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Group {
	return predicate.Group(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldType, v))
}

// MaxUsersEQ applies the EQ predicate on the "max_users" field.
func MaxUsersEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldMaxUsers, v))
}

// MaxUsersNEQ applies the NEQ predicate on the "max_users" field.
func MaxUsersNEQ(v int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldMaxUsers, v))
}

// MaxUsersIn applies the In predicate on the "max_users" field.
func MaxUsersIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldMaxUsers, vs...))
}

// MaxUsersNotIn applies the NotIn predicate on the "max_users" field.
func MaxUsersNotIn(vs ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldMaxUsers, vs...))
}

// MaxUsersGT applies the GT predicate on the "max_users" field.
func MaxUsersGT(v int) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldMaxUsers, v))
}

// MaxUsersGTE applies the GTE predicate on the "max_users" field.
func MaxUsersGTE(v int) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldMaxUsers, v))
}

// MaxUsersLT applies the LT predicate on the "max_users" field.
func MaxUsersLT(v int) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldMaxUsers, v))
}

// MaxUsersLTE applies the LTE predicate on the "max_users" field.
func MaxUsersLTE(v int) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldMaxUsers, v))
}

// MaxUsersIsNil applies the IsNil predicate on the "max_users" field.
func MaxUsersIsNil() predicate.Group {
	return predicate.Group(sql.FieldIsNull(FieldMaxUsers))
}

// MaxUsersNotNil applies the NotNil predicate on the "max_users" field.
func MaxUsersNotNil() predicate.Group {
	return predicate.Group(sql.FieldNotNull(FieldMaxUsers))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldName, v))
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FilesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBlocked applies the HasEdge predicate on the "blocked" edge.
func HasBlocked() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BlockedTable, BlockedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlockedWith applies the HasEdge predicate on the "blocked" edge with a given conditions (other predicates).
func HasBlockedWith(preds ...predicate.User) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlockedInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BlockedTable, BlockedColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInfo applies the HasEdge predicate on the "info" edge.
func HasInfo() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, InfoTable, InfoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInfoWith applies the HasEdge predicate on the "info" edge with a given conditions (other predicates).
func HasInfoWith(preds ...predicate.GroupInfo) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InfoInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, InfoTable, InfoColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		p(s.Not())
	})
}
