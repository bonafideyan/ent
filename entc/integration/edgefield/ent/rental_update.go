// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"entgo.io/ent/entc/integration/edgefield/ent/rental"
	"entgo.io/ent/schema/field"
)

// RentalUpdate is the builder for updating Rental entities.
type RentalUpdate struct {
	config
	hooks    []Hook
	mutation *RentalMutation
}

// Where appends a list predicates to the RentalUpdate builder.
func (ru *RentalUpdate) Where(ps ...predicate.Rental) *RentalUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetDate sets the "date" field.
func (ru *RentalUpdate) SetDate(t time.Time) *RentalUpdate {
	ru.mutation.SetDate(t)
	return ru
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (ru *RentalUpdate) SetNillableDate(t *time.Time) *RentalUpdate {
	if t != nil {
		ru.SetDate(*t)
	}
	return ru
}

// Mutation returns the RentalMutation object of the builder.
func (ru *RentalUpdate) Mutation() *RentalMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RentalUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RentalUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RentalUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RentalUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RentalUpdate) check() error {
	if ru.mutation.UserCleared() && len(ru.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Rental.user"`)
	}
	if ru.mutation.CarCleared() && len(ru.mutation.CarIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Rental.car"`)
	}
	return nil
}

func (ru *RentalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(rental.Table, rental.Columns, sqlgraph.NewFieldSpec(rental.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Date(); ok {
		_spec.SetField(rental.FieldDate, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rental.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RentalUpdateOne is the builder for updating a single Rental entity.
type RentalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RentalMutation
}

// SetDate sets the "date" field.
func (ruo *RentalUpdateOne) SetDate(t time.Time) *RentalUpdateOne {
	ruo.mutation.SetDate(t)
	return ruo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (ruo *RentalUpdateOne) SetNillableDate(t *time.Time) *RentalUpdateOne {
	if t != nil {
		ruo.SetDate(*t)
	}
	return ruo
}

// Mutation returns the RentalMutation object of the builder.
func (ruo *RentalUpdateOne) Mutation() *RentalMutation {
	return ruo.mutation
}

// Where appends a list predicates to the RentalUpdate builder.
func (ruo *RentalUpdateOne) Where(ps ...predicate.Rental) *RentalUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RentalUpdateOne) Select(field string, fields ...string) *RentalUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Rental entity.
func (ruo *RentalUpdateOne) Save(ctx context.Context) (*Rental, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RentalUpdateOne) SaveX(ctx context.Context) *Rental {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RentalUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RentalUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RentalUpdateOne) check() error {
	if ruo.mutation.UserCleared() && len(ruo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Rental.user"`)
	}
	if ruo.mutation.CarCleared() && len(ruo.mutation.CarIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Rental.car"`)
	}
	return nil
}

func (ruo *RentalUpdateOne) sqlSave(ctx context.Context) (_node *Rental, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(rental.Table, rental.Columns, sqlgraph.NewFieldSpec(rental.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Rental.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rental.FieldID)
		for _, f := range fields {
			if !rental.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rental.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Date(); ok {
		_spec.SetField(rental.FieldDate, field.TypeTime, value)
	}
	_node = &Rental{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rental.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}