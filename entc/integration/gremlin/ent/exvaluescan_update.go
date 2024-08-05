// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"math/big"
	"net/url"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/exvaluescan"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// ExValueScanUpdate is the builder for updating ExValueScan entities.
type ExValueScanUpdate struct {
	config
	hooks    []Hook
	mutation *ExValueScanMutation
}

// Where appends a list predicates to the ExValueScanUpdate builder.
func (evsu *ExValueScanUpdate) Where(ps ...predicate.ExValueScan) *ExValueScanUpdate {
	evsu.mutation.Where(ps...)
	return evsu
}

// SetBinary sets the "binary" field.
func (evsu *ExValueScanUpdate) SetBinary(u *url.URL) *ExValueScanUpdate {
	evsu.mutation.SetBinary(u)
	return evsu
}

// SetBinaryBytes sets the "binary_bytes" field.
func (evsu *ExValueScanUpdate) SetBinaryBytes(u *url.URL) *ExValueScanUpdate {
	evsu.mutation.SetBinaryBytes(u)
	return evsu
}

// SetBinaryOptional sets the "binary_optional" field.
func (evsu *ExValueScanUpdate) SetBinaryOptional(u *url.URL) *ExValueScanUpdate {
	evsu.mutation.SetBinaryOptional(u)
	return evsu
}

// ClearBinaryOptional clears the value of the "binary_optional" field.
func (evsu *ExValueScanUpdate) ClearBinaryOptional() *ExValueScanUpdate {
	evsu.mutation.ClearBinaryOptional()
	return evsu
}

// SetText sets the "text" field.
func (evsu *ExValueScanUpdate) SetText(b *big.Int) *ExValueScanUpdate {
	evsu.mutation.SetText(b)
	return evsu
}

// SetTextOptional sets the "text_optional" field.
func (evsu *ExValueScanUpdate) SetTextOptional(b *big.Int) *ExValueScanUpdate {
	evsu.mutation.SetTextOptional(b)
	return evsu
}

// ClearTextOptional clears the value of the "text_optional" field.
func (evsu *ExValueScanUpdate) ClearTextOptional() *ExValueScanUpdate {
	evsu.mutation.ClearTextOptional()
	return evsu
}

// SetBase64 sets the "base64" field.
func (evsu *ExValueScanUpdate) SetBase64(s string) *ExValueScanUpdate {
	evsu.mutation.SetBase64(s)
	return evsu
}

// SetNillableBase64 sets the "base64" field if the given value is not nil.
func (evsu *ExValueScanUpdate) SetNillableBase64(s *string) *ExValueScanUpdate {
	if s != nil {
		evsu.SetBase64(*s)
	}
	return evsu
}

// SetCustom sets the "custom" field.
func (evsu *ExValueScanUpdate) SetCustom(s string) *ExValueScanUpdate {
	evsu.mutation.SetCustom(s)
	return evsu
}

// SetNillableCustom sets the "custom" field if the given value is not nil.
func (evsu *ExValueScanUpdate) SetNillableCustom(s *string) *ExValueScanUpdate {
	if s != nil {
		evsu.SetCustom(*s)
	}
	return evsu
}

// SetCustomOptional sets the "custom_optional" field.
func (evsu *ExValueScanUpdate) SetCustomOptional(s string) *ExValueScanUpdate {
	evsu.mutation.SetCustomOptional(s)
	return evsu
}

// SetNillableCustomOptional sets the "custom_optional" field if the given value is not nil.
func (evsu *ExValueScanUpdate) SetNillableCustomOptional(s *string) *ExValueScanUpdate {
	if s != nil {
		evsu.SetCustomOptional(*s)
	}
	return evsu
}

// ClearCustomOptional clears the value of the "custom_optional" field.
func (evsu *ExValueScanUpdate) ClearCustomOptional() *ExValueScanUpdate {
	evsu.mutation.ClearCustomOptional()
	return evsu
}

// Mutation returns the ExValueScanMutation object of the builder.
func (evsu *ExValueScanUpdate) Mutation() *ExValueScanMutation {
	return evsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (evsu *ExValueScanUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, evsu.gremlinSave, evsu.mutation, evsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (evsu *ExValueScanUpdate) SaveX(ctx context.Context) int {
	affected, err := evsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (evsu *ExValueScanUpdate) Exec(ctx context.Context) error {
	_, err := evsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (evsu *ExValueScanUpdate) ExecX(ctx context.Context) {
	if err := evsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (evsu *ExValueScanUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := evsu.gremlin().Query()
	if err := evsu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	evsu.mutation.done = true
	return res.ReadInt()
}

func (evsu *ExValueScanUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(exvaluescan.Label)
	for _, p := range evsu.mutation.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	if value, ok := evsu.mutation.Binary(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBinary, value)
	}
	if value, ok := evsu.mutation.BinaryBytes(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBinaryBytes, value)
	}
	if value, ok := evsu.mutation.BinaryOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBinaryOptional, value)
	}
	if value, ok := evsu.mutation.Text(); ok {
		v.Property(dsl.Single, exvaluescan.FieldText, value)
	}
	if value, ok := evsu.mutation.TextOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldTextOptional, value)
	}
	if value, ok := evsu.mutation.Base64(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBase64, value)
	}
	if value, ok := evsu.mutation.Custom(); ok {
		v.Property(dsl.Single, exvaluescan.FieldCustom, value)
	}
	if value, ok := evsu.mutation.CustomOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldCustomOptional, value)
	}
	var properties []any
	if evsu.mutation.BinaryOptionalCleared() {
		properties = append(properties, exvaluescan.FieldBinaryOptional)
	}
	if evsu.mutation.TextOptionalCleared() {
		properties = append(properties, exvaluescan.FieldTextOptional)
	}
	if evsu.mutation.CustomOptionalCleared() {
		properties = append(properties, exvaluescan.FieldCustomOptional)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// ExValueScanUpdateOne is the builder for updating a single ExValueScan entity.
type ExValueScanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExValueScanMutation
}

// SetBinary sets the "binary" field.
func (evsuo *ExValueScanUpdateOne) SetBinary(u *url.URL) *ExValueScanUpdateOne {
	evsuo.mutation.SetBinary(u)
	return evsuo
}

// SetBinaryBytes sets the "binary_bytes" field.
func (evsuo *ExValueScanUpdateOne) SetBinaryBytes(u *url.URL) *ExValueScanUpdateOne {
	evsuo.mutation.SetBinaryBytes(u)
	return evsuo
}

// SetBinaryOptional sets the "binary_optional" field.
func (evsuo *ExValueScanUpdateOne) SetBinaryOptional(u *url.URL) *ExValueScanUpdateOne {
	evsuo.mutation.SetBinaryOptional(u)
	return evsuo
}

// ClearBinaryOptional clears the value of the "binary_optional" field.
func (evsuo *ExValueScanUpdateOne) ClearBinaryOptional() *ExValueScanUpdateOne {
	evsuo.mutation.ClearBinaryOptional()
	return evsuo
}

// SetText sets the "text" field.
func (evsuo *ExValueScanUpdateOne) SetText(b *big.Int) *ExValueScanUpdateOne {
	evsuo.mutation.SetText(b)
	return evsuo
}

// SetTextOptional sets the "text_optional" field.
func (evsuo *ExValueScanUpdateOne) SetTextOptional(b *big.Int) *ExValueScanUpdateOne {
	evsuo.mutation.SetTextOptional(b)
	return evsuo
}

// ClearTextOptional clears the value of the "text_optional" field.
func (evsuo *ExValueScanUpdateOne) ClearTextOptional() *ExValueScanUpdateOne {
	evsuo.mutation.ClearTextOptional()
	return evsuo
}

// SetBase64 sets the "base64" field.
func (evsuo *ExValueScanUpdateOne) SetBase64(s string) *ExValueScanUpdateOne {
	evsuo.mutation.SetBase64(s)
	return evsuo
}

// SetNillableBase64 sets the "base64" field if the given value is not nil.
func (evsuo *ExValueScanUpdateOne) SetNillableBase64(s *string) *ExValueScanUpdateOne {
	if s != nil {
		evsuo.SetBase64(*s)
	}
	return evsuo
}

// SetCustom sets the "custom" field.
func (evsuo *ExValueScanUpdateOne) SetCustom(s string) *ExValueScanUpdateOne {
	evsuo.mutation.SetCustom(s)
	return evsuo
}

// SetNillableCustom sets the "custom" field if the given value is not nil.
func (evsuo *ExValueScanUpdateOne) SetNillableCustom(s *string) *ExValueScanUpdateOne {
	if s != nil {
		evsuo.SetCustom(*s)
	}
	return evsuo
}

// SetCustomOptional sets the "custom_optional" field.
func (evsuo *ExValueScanUpdateOne) SetCustomOptional(s string) *ExValueScanUpdateOne {
	evsuo.mutation.SetCustomOptional(s)
	return evsuo
}

// SetNillableCustomOptional sets the "custom_optional" field if the given value is not nil.
func (evsuo *ExValueScanUpdateOne) SetNillableCustomOptional(s *string) *ExValueScanUpdateOne {
	if s != nil {
		evsuo.SetCustomOptional(*s)
	}
	return evsuo
}

// ClearCustomOptional clears the value of the "custom_optional" field.
func (evsuo *ExValueScanUpdateOne) ClearCustomOptional() *ExValueScanUpdateOne {
	evsuo.mutation.ClearCustomOptional()
	return evsuo
}

// Mutation returns the ExValueScanMutation object of the builder.
func (evsuo *ExValueScanUpdateOne) Mutation() *ExValueScanMutation {
	return evsuo.mutation
}

// Where appends a list predicates to the ExValueScanUpdate builder.
func (evsuo *ExValueScanUpdateOne) Where(ps ...predicate.ExValueScan) *ExValueScanUpdateOne {
	evsuo.mutation.Where(ps...)
	return evsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (evsuo *ExValueScanUpdateOne) Select(field string, fields ...string) *ExValueScanUpdateOne {
	evsuo.fields = append([]string{field}, fields...)
	return evsuo
}

// Save executes the query and returns the updated ExValueScan entity.
func (evsuo *ExValueScanUpdateOne) Save(ctx context.Context) (*ExValueScan, error) {
	return withHooks(ctx, evsuo.gremlinSave, evsuo.mutation, evsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (evsuo *ExValueScanUpdateOne) SaveX(ctx context.Context) *ExValueScan {
	node, err := evsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (evsuo *ExValueScanUpdateOne) Exec(ctx context.Context) error {
	_, err := evsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (evsuo *ExValueScanUpdateOne) ExecX(ctx context.Context) {
	if err := evsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (evsuo *ExValueScanUpdateOne) gremlinSave(ctx context.Context) (*ExValueScan, error) {
	res := &gremlin.Response{}
	id, ok := evsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ExValueScan.id" for update`)}
	}
	query, bindings := evsuo.gremlin(id).Query()
	if err := evsuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	evsuo.mutation.done = true
	evs := &ExValueScan{config: evsuo.config}
	if err := evs.FromResponse(res); err != nil {
		return nil, err
	}
	return evs, nil
}

func (evsuo *ExValueScanUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if value, ok := evsuo.mutation.Binary(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBinary, value)
	}
	if value, ok := evsuo.mutation.BinaryBytes(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBinaryBytes, value)
	}
	if value, ok := evsuo.mutation.BinaryOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBinaryOptional, value)
	}
	if value, ok := evsuo.mutation.Text(); ok {
		v.Property(dsl.Single, exvaluescan.FieldText, value)
	}
	if value, ok := evsuo.mutation.TextOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldTextOptional, value)
	}
	if value, ok := evsuo.mutation.Base64(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBase64, value)
	}
	if value, ok := evsuo.mutation.Custom(); ok {
		v.Property(dsl.Single, exvaluescan.FieldCustom, value)
	}
	if value, ok := evsuo.mutation.CustomOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldCustomOptional, value)
	}
	var properties []any
	if evsuo.mutation.BinaryOptionalCleared() {
		properties = append(properties, exvaluescan.FieldBinaryOptional)
	}
	if evsuo.mutation.TextOptionalCleared() {
		properties = append(properties, exvaluescan.FieldTextOptional)
	}
	if evsuo.mutation.CustomOptionalCleared() {
		properties = append(properties, exvaluescan.FieldCustomOptional)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if len(evsuo.fields) > 0 {
		fields := make([]any, 0, len(evsuo.fields)+1)
		fields = append(fields, true)
		for _, f := range evsuo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
