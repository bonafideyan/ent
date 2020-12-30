// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/facebook/ent/examples/privacytenant/ent/schema","Package":"github.com/facebook/ent/examples/privacytenant/ent","Schemas":[{"name":"Group","config":{"Table":""},"edges":[{"name":"tenant","type":"Tenant","unique":true,"required":true},{"name":"users","type":"User","ref_name":"groups","inverse":true}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"default":true,"default_value":"Unknown","default_kind":24,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":true,"MixinIndex":1},{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"Tenant","config":{"Table":""},"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"User","config":{"Table":""},"edges":[{"name":"tenant","type":"Tenant","unique":true,"required":true},{"name":"groups","type":"Group"}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"default":true,"default_value":"Unknown","default_kind":24,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":true,"MixinIndex":1}]}],"Features":["privacy","entql","schema/snapshot"]}`
