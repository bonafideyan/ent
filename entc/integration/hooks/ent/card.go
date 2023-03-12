// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/hooks/ent/card"
	"entgo.io/ent/entc/integration/hooks/ent/user"
)

// Card is the model entity for the Card schema.
type Card struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// Exact name written on card
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// InHook is a mandatory field that is set by the hook.
	InHook string `json:"in_hook,omitempty"`
	// ExpiredAt holds the value of the "expired_at" field.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardQuery when eager-loading is set.
	Edges      CardEdges `json:"edges"`
	user_cards *int
}

// CardEdges holds the relations/edges for other nodes in the graph.
type CardEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Card) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case card.FieldID:
			values[i] = new(sql.NullInt64)
		case card.FieldNumber, card.FieldName, card.FieldInHook:
			values[i] = new(sql.NullString)
		case card.FieldCreatedAt, card.FieldExpiredAt:
			values[i] = new(sql.NullTime)
		case card.ForeignKeys[0]: // user_cards
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Card", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Card fields.
func (c *Card) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case card.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case card.FieldNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				c.Number = value.String
			}
		case card.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case card.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case card.FieldInHook:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field in_hook", values[i])
			} else if value.Valid {
				c.InHook = value.String
			}
		case card.FieldExpiredAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expired_at", values[i])
			} else if value.Valid {
				c.ExpiredAt = value.Time
			}
		case card.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_cards", value)
			} else if value.Valid {
				c.user_cards = new(int)
				*c.user_cards = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Card entity.
func (c *Card) QueryOwner() *UserQuery {
	return NewCardClient(c.config).QueryOwner(c)
}

// Update returns a builder for updating this Card.
// Note that you need to call Card.Unwrap() before calling this method if this Card
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Card) Update() *CardUpdateOne {
	return NewCardClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Card entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Card) Unwrap() *Card {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Card is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Card) String() string {
	var builder strings.Builder
	builder.WriteString("Card(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("number=")
	builder.WriteString(c.Number)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("in_hook=")
	builder.WriteString(c.InHook)
	builder.WriteString(", ")
	builder.WriteString("expired_at=")
	builder.WriteString(c.ExpiredAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Cards is a parsable slice of Card.
type Cards []*Card
