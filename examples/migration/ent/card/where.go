// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package card

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/migration/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldType, v))
}

// NumberHash applies equality check predicate on the "number_hash" field. It's identical to NumberHashEQ.
func NumberHash(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldNumberHash, v))
}

// CvvHash applies equality check predicate on the "cvv_hash" field. It's identical to CvvHashEQ.
func CvvHash(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCvvHash, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldExpiresAt, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldOwnerID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldType, v))
}

// NumberHashEQ applies the EQ predicate on the "number_hash" field.
func NumberHashEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldNumberHash, v))
}

// NumberHashNEQ applies the NEQ predicate on the "number_hash" field.
func NumberHashNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldNumberHash, v))
}

// NumberHashIn applies the In predicate on the "number_hash" field.
func NumberHashIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldNumberHash, vs...))
}

// NumberHashNotIn applies the NotIn predicate on the "number_hash" field.
func NumberHashNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldNumberHash, vs...))
}

// NumberHashGT applies the GT predicate on the "number_hash" field.
func NumberHashGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldNumberHash, v))
}

// NumberHashGTE applies the GTE predicate on the "number_hash" field.
func NumberHashGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldNumberHash, v))
}

// NumberHashLT applies the LT predicate on the "number_hash" field.
func NumberHashLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldNumberHash, v))
}

// NumberHashLTE applies the LTE predicate on the "number_hash" field.
func NumberHashLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldNumberHash, v))
}

// NumberHashContains applies the Contains predicate on the "number_hash" field.
func NumberHashContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldNumberHash, v))
}

// NumberHashHasPrefix applies the HasPrefix predicate on the "number_hash" field.
func NumberHashHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldNumberHash, v))
}

// NumberHashHasSuffix applies the HasSuffix predicate on the "number_hash" field.
func NumberHashHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldNumberHash, v))
}

// NumberHashEqualFold applies the EqualFold predicate on the "number_hash" field.
func NumberHashEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldNumberHash, v))
}

// NumberHashContainsFold applies the ContainsFold predicate on the "number_hash" field.
func NumberHashContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldNumberHash, v))
}

// CvvHashEQ applies the EQ predicate on the "cvv_hash" field.
func CvvHashEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCvvHash, v))
}

// CvvHashNEQ applies the NEQ predicate on the "cvv_hash" field.
func CvvHashNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldCvvHash, v))
}

// CvvHashIn applies the In predicate on the "cvv_hash" field.
func CvvHashIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldCvvHash, vs...))
}

// CvvHashNotIn applies the NotIn predicate on the "cvv_hash" field.
func CvvHashNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldCvvHash, vs...))
}

// CvvHashGT applies the GT predicate on the "cvv_hash" field.
func CvvHashGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldCvvHash, v))
}

// CvvHashGTE applies the GTE predicate on the "cvv_hash" field.
func CvvHashGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldCvvHash, v))
}

// CvvHashLT applies the LT predicate on the "cvv_hash" field.
func CvvHashLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldCvvHash, v))
}

// CvvHashLTE applies the LTE predicate on the "cvv_hash" field.
func CvvHashLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldCvvHash, v))
}

// CvvHashContains applies the Contains predicate on the "cvv_hash" field.
func CvvHashContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldCvvHash, v))
}

// CvvHashHasPrefix applies the HasPrefix predicate on the "cvv_hash" field.
func CvvHashHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldCvvHash, v))
}

// CvvHashHasSuffix applies the HasSuffix predicate on the "cvv_hash" field.
func CvvHashHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldCvvHash, v))
}

// CvvHashEqualFold applies the EqualFold predicate on the "cvv_hash" field.
func CvvHashEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldCvvHash, v))
}

// CvvHashContainsFold applies the ContainsFold predicate on the "cvv_hash" field.
func CvvHashContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldCvvHash, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldExpiresAt, v))
}

// ExpiresAtIsNil applies the IsNil predicate on the "expires_at" field.
func ExpiresAtIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldExpiresAt))
}

// ExpiresAtNotNil applies the NotNil predicate on the "expires_at" field.
func ExpiresAtNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldExpiresAt))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v int) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...int) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...int) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldOwnerID, vs...))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPayments applies the HasEdge predicate on the "payments" edge.
func HasPayments() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentsTable, PaymentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentsWith applies the HasEdge predicate on the "payments" edge with a given conditions (other predicates).
func HasPaymentsWith(preds ...predicate.Payment) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := newPaymentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Card) predicate.Card {
	return predicate.Card(sql.NotPredicates(p))
}