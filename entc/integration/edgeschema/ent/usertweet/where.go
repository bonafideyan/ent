// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package usertweet

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldEQ(FieldCreatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldEQ(FieldUserID, v))
}

// TweetID applies equality check predicate on the "tweet_id" field. It's identical to TweetIDEQ.
func TweetID(v int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldEQ(FieldTweetID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldLTE(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldNotIn(FieldUserID, vs...))
}

// TweetIDEQ applies the EQ predicate on the "tweet_id" field.
func TweetIDEQ(v int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldEQ(FieldTweetID, v))
}

// TweetIDNEQ applies the NEQ predicate on the "tweet_id" field.
func TweetIDNEQ(v int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldNEQ(FieldTweetID, v))
}

// TweetIDIn applies the In predicate on the "tweet_id" field.
func TweetIDIn(vs ...int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldIn(FieldTweetID, vs...))
}

// TweetIDNotIn applies the NotIn predicate on the "tweet_id" field.
func TweetIDNotIn(vs ...int) predicate.UserTweet {
	return predicate.UserTweet(sql.FieldNotIn(FieldTweetID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserTweet {
	return predicate.UserTweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserTweet {
	return predicate.UserTweet(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTweet applies the HasEdge predicate on the "tweet" edge.
func HasTweet() predicate.UserTweet {
	return predicate.UserTweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TweetTable, TweetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTweetWith applies the HasEdge predicate on the "tweet" edge with a given conditions (other predicates).
func HasTweetWith(preds ...predicate.Tweet) predicate.UserTweet {
	return predicate.UserTweet(func(s *sql.Selector) {
		step := newTweetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserTweet) predicate.UserTweet {
	return predicate.UserTweet(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserTweet) predicate.UserTweet {
	return predicate.UserTweet(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserTweet) predicate.UserTweet {
	return predicate.UserTweet(sql.NotPredicates(p))
}