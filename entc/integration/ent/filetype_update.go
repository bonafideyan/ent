// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/filetype"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// FileTypeUpdate is the builder for updating FileType entities.
type FileTypeUpdate struct {
	config
	hooks     []Hook
	mutation  *FileTypeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FileTypeUpdate builder.
func (ftu *FileTypeUpdate) Where(ps ...predicate.FileType) *FileTypeUpdate {
	ftu.mutation.Where(ps...)
	return ftu
}

// SetName sets the "name" field.
func (ftu *FileTypeUpdate) SetName(s string) *FileTypeUpdate {
	ftu.mutation.SetName(s)
	return ftu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ftu *FileTypeUpdate) SetNillableName(s *string) *FileTypeUpdate {
	if s != nil {
		ftu.SetName(*s)
	}
	return ftu
}

// SetType sets the "type" field.
func (ftu *FileTypeUpdate) SetType(f filetype.Type) *FileTypeUpdate {
	ftu.mutation.SetType(f)
	return ftu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ftu *FileTypeUpdate) SetNillableType(f *filetype.Type) *FileTypeUpdate {
	if f != nil {
		ftu.SetType(*f)
	}
	return ftu
}

// SetState sets the "state" field.
func (ftu *FileTypeUpdate) SetState(f filetype.State) *FileTypeUpdate {
	ftu.mutation.SetState(f)
	return ftu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ftu *FileTypeUpdate) SetNillableState(f *filetype.State) *FileTypeUpdate {
	if f != nil {
		ftu.SetState(*f)
	}
	return ftu
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (ftu *FileTypeUpdate) AddFileIDs(ids ...int) *FileTypeUpdate {
	ftu.mutation.AddFileIDs(ids...)
	return ftu
}

// AddFiles adds the "files" edges to the File entity.
func (ftu *FileTypeUpdate) AddFiles(f ...*File) *FileTypeUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.AddFileIDs(ids...)
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftu *FileTypeUpdate) Mutation() *FileTypeMutation {
	return ftu.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (ftu *FileTypeUpdate) ClearFiles() *FileTypeUpdate {
	ftu.mutation.ClearFiles()
	return ftu
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (ftu *FileTypeUpdate) RemoveFileIDs(ids ...int) *FileTypeUpdate {
	ftu.mutation.RemoveFileIDs(ids...)
	return ftu
}

// RemoveFiles removes "files" edges to File entities.
func (ftu *FileTypeUpdate) RemoveFiles(f ...*File) *FileTypeUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ftu *FileTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ftu.sqlSave, ftu.mutation, ftu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FileTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FileTypeUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FileTypeUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftu *FileTypeUpdate) check() error {
	if v, ok := ftu.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	if v, ok := ftu.mutation.State(); ok {
		if err := filetype.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "FileType.state": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ftu *FileTypeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FileTypeUpdate {
	ftu.modifiers = append(ftu.modifiers, modifiers...)
	return ftu
}

func (ftu *FileTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ftu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(filetype.Table, filetype.Columns, sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeInt))
	if ps := ftu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftu.mutation.Name(); ok {
		_spec.SetField(filetype.FieldName, field.TypeString, value)
	}
	if value, ok := ftu.mutation.GetType(); ok {
		_spec.SetField(filetype.FieldType, field.TypeEnum, value)
	}
	if value, ok := ftu.mutation.State(); ok {
		_spec.SetField(filetype.FieldState, field.TypeEnum, value)
	}
	if ftu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.RemovedFilesIDs(); len(nodes) > 0 && !ftu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ftu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ftu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ftu.mutation.done = true
	return n, nil
}

// FileTypeUpdateOne is the builder for updating a single FileType entity.
type FileTypeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FileTypeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (ftuo *FileTypeUpdateOne) SetName(s string) *FileTypeUpdateOne {
	ftuo.mutation.SetName(s)
	return ftuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ftuo *FileTypeUpdateOne) SetNillableName(s *string) *FileTypeUpdateOne {
	if s != nil {
		ftuo.SetName(*s)
	}
	return ftuo
}

// SetType sets the "type" field.
func (ftuo *FileTypeUpdateOne) SetType(f filetype.Type) *FileTypeUpdateOne {
	ftuo.mutation.SetType(f)
	return ftuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ftuo *FileTypeUpdateOne) SetNillableType(f *filetype.Type) *FileTypeUpdateOne {
	if f != nil {
		ftuo.SetType(*f)
	}
	return ftuo
}

// SetState sets the "state" field.
func (ftuo *FileTypeUpdateOne) SetState(f filetype.State) *FileTypeUpdateOne {
	ftuo.mutation.SetState(f)
	return ftuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ftuo *FileTypeUpdateOne) SetNillableState(f *filetype.State) *FileTypeUpdateOne {
	if f != nil {
		ftuo.SetState(*f)
	}
	return ftuo
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (ftuo *FileTypeUpdateOne) AddFileIDs(ids ...int) *FileTypeUpdateOne {
	ftuo.mutation.AddFileIDs(ids...)
	return ftuo
}

// AddFiles adds the "files" edges to the File entity.
func (ftuo *FileTypeUpdateOne) AddFiles(f ...*File) *FileTypeUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.AddFileIDs(ids...)
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftuo *FileTypeUpdateOne) Mutation() *FileTypeMutation {
	return ftuo.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (ftuo *FileTypeUpdateOne) ClearFiles() *FileTypeUpdateOne {
	ftuo.mutation.ClearFiles()
	return ftuo
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (ftuo *FileTypeUpdateOne) RemoveFileIDs(ids ...int) *FileTypeUpdateOne {
	ftuo.mutation.RemoveFileIDs(ids...)
	return ftuo
}

// RemoveFiles removes "files" edges to File entities.
func (ftuo *FileTypeUpdateOne) RemoveFiles(f ...*File) *FileTypeUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.RemoveFileIDs(ids...)
}

// Where appends a list predicates to the FileTypeUpdate builder.
func (ftuo *FileTypeUpdateOne) Where(ps ...predicate.FileType) *FileTypeUpdateOne {
	ftuo.mutation.Where(ps...)
	return ftuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ftuo *FileTypeUpdateOne) Select(field string, fields ...string) *FileTypeUpdateOne {
	ftuo.fields = append([]string{field}, fields...)
	return ftuo
}

// Save executes the query and returns the updated FileType entity.
func (ftuo *FileTypeUpdateOne) Save(ctx context.Context) (*FileType, error) {
	return withHooks(ctx, ftuo.sqlSave, ftuo.mutation, ftuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) SaveX(ctx context.Context) *FileType {
	node, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ftuo *FileTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftuo *FileTypeUpdateOne) check() error {
	if v, ok := ftuo.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	if v, ok := ftuo.mutation.State(); ok {
		if err := filetype.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "FileType.state": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ftuo *FileTypeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FileTypeUpdateOne {
	ftuo.modifiers = append(ftuo.modifiers, modifiers...)
	return ftuo
}

func (ftuo *FileTypeUpdateOne) sqlSave(ctx context.Context) (_node *FileType, err error) {
	if err := ftuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(filetype.Table, filetype.Columns, sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeInt))
	id, ok := ftuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FileType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ftuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filetype.FieldID)
		for _, f := range fields {
			if !filetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != filetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ftuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftuo.mutation.Name(); ok {
		_spec.SetField(filetype.FieldName, field.TypeString, value)
	}
	if value, ok := ftuo.mutation.GetType(); ok {
		_spec.SetField(filetype.FieldType, field.TypeEnum, value)
	}
	if value, ok := ftuo.mutation.State(); ok {
		_spec.SetField(filetype.FieldState, field.TypeEnum, value)
	}
	if ftuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.RemovedFilesIDs(); len(nodes) > 0 && !ftuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ftuo.modifiers...)
	_node = &FileType{config: ftuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ftuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ftuo.mutation.done = true
	return _node, nil
}
