// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package usergroup

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldID, id))
}

// JoinedAt applies equality check predicate on the "joined_at" field. It's identical to JoinedAtEQ.
func JoinedAt(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldJoinedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldUserID, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldGroupID, v))
}

// JoinedAtEQ applies the EQ predicate on the "joined_at" field.
func JoinedAtEQ(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldJoinedAt, v))
}

// JoinedAtNEQ applies the NEQ predicate on the "joined_at" field.
func JoinedAtNEQ(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldJoinedAt, v))
}

// JoinedAtIn applies the In predicate on the "joined_at" field.
func JoinedAtIn(vs ...time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldJoinedAt, vs...))
}

// JoinedAtNotIn applies the NotIn predicate on the "joined_at" field.
func JoinedAtNotIn(vs ...time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldJoinedAt, vs...))
}

// JoinedAtGT applies the GT predicate on the "joined_at" field.
func JoinedAtGT(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldJoinedAt, v))
}

// JoinedAtGTE applies the GTE predicate on the "joined_at" field.
func JoinedAtGTE(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldJoinedAt, v))
}

// JoinedAtLT applies the LT predicate on the "joined_at" field.
func JoinedAtLT(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldJoinedAt, v))
}

// JoinedAtLTE applies the LTE predicate on the "joined_at" field.
func JoinedAtLTE(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldJoinedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldUserID, vs...))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldGroupID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserGroup {
	return predicate.UserGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserGroup {
	return predicate.UserGroup(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.UserGroup {
	return predicate.UserGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.UserGroup {
	return predicate.UserGroup(func(s *sql.Selector) {
		step := newGroupStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserGroup) predicate.UserGroup {
	return predicate.UserGroup(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserGroup) predicate.UserGroup {
	return predicate.UserGroup(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserGroup) predicate.UserGroup {
	return predicate.UserGroup(sql.NotPredicates(p))
}
