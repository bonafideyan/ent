// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package car

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Car {
	return predicate.Car(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Car {
	return predicate.Car(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Car {
	return predicate.Car(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Car {
	return predicate.Car(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Car {
	return predicate.Car(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Car {
	return predicate.Car(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Car {
	return predicate.Car(sql.FieldLTE(FieldID, id))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldNumber, v))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) predicate.Car {
	return predicate.Car(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) predicate.Car {
	return predicate.Car(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) predicate.Car {
	return predicate.Car(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) predicate.Car {
	return predicate.Car(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) predicate.Car {
	return predicate.Car(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) predicate.Car {
	return predicate.Car(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) predicate.Car {
	return predicate.Car(sql.FieldLTE(FieldNumber, v))
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) predicate.Car {
	return predicate.Car(sql.FieldContains(FieldNumber, v))
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) predicate.Car {
	return predicate.Car(sql.FieldHasPrefix(FieldNumber, v))
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) predicate.Car {
	return predicate.Car(sql.FieldHasSuffix(FieldNumber, v))
}

// NumberIsNil applies the IsNil predicate on the "number" field.
func NumberIsNil() predicate.Car {
	return predicate.Car(sql.FieldIsNull(FieldNumber))
}

// NumberNotNil applies the NotNil predicate on the "number" field.
func NumberNotNil() predicate.Car {
	return predicate.Car(sql.FieldNotNull(FieldNumber))
}

// NumberEqualFold applies the EqualFold predicate on the "number" field.
func NumberEqualFold(v string) predicate.Car {
	return predicate.Car(sql.FieldEqualFold(FieldNumber, v))
}

// NumberContainsFold applies the ContainsFold predicate on the "number" field.
func NumberContainsFold(v string) predicate.Car {
	return predicate.Car(sql.FieldContainsFold(FieldNumber, v))
}

// HasRentals applies the HasEdge predicate on the "rentals" edge.
func HasRentals() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RentalsTable, RentalsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRentalsWith applies the HasEdge predicate on the "rentals" edge with a given conditions (other predicates).
func HasRentalsWith(preds ...predicate.Rental) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := newRentalsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Car) predicate.Car {
	return predicate.Car(sql.NotPredicates(p))
}