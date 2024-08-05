// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package comment

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUniqueInt holds the string denoting the unique_int field in the database.
	FieldUniqueInt = "unique_int"
	// FieldUniqueFloat holds the string denoting the unique_float field in the database.
	FieldUniqueFloat = "unique_float"
	// FieldNillableInt holds the string denoting the nillable_int field in the database.
	FieldNillableInt = "nillable_int"
	// FieldTable holds the string denoting the table field in the database.
	FieldTable = "table"
	// FieldDir holds the string denoting the dir field in the database.
	FieldDir = "dir"
	// FieldClient holds the string denoting the client field in the database.
	FieldClient = "client"
)

// OrderOption defines the ordering options for the Comment queries.
type OrderOption func(*dsl.Traversal)

// comment from another template.
