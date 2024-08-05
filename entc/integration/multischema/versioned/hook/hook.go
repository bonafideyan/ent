// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"entgo.io/ent/entc/integration/multischema/versioned"
)

// The FriendshipFunc type is an adapter to allow the use of ordinary
// function as Friendship mutator.
type FriendshipFunc func(context.Context, *versioned.FriendshipMutation) (versioned.Value, error)

// Mutate calls f(ctx, m).
func (f FriendshipFunc) Mutate(ctx context.Context, m versioned.Mutation) (versioned.Value, error) {
	if mv, ok := m.(*versioned.FriendshipMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *versioned.FriendshipMutation", m)
}

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *versioned.GroupMutation) (versioned.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m versioned.Mutation) (versioned.Value, error) {
	if mv, ok := m.(*versioned.GroupMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *versioned.GroupMutation", m)
}

// The PetFunc type is an adapter to allow the use of ordinary
// function as Pet mutator.
type PetFunc func(context.Context, *versioned.PetMutation) (versioned.Value, error)

// Mutate calls f(ctx, m).
func (f PetFunc) Mutate(ctx context.Context, m versioned.Mutation) (versioned.Value, error) {
	if mv, ok := m.(*versioned.PetMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *versioned.PetMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *versioned.UserMutation) (versioned.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m versioned.Mutation) (versioned.Value, error) {
	if mv, ok := m.(*versioned.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *versioned.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, versioned.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m versioned.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m versioned.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m versioned.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op versioned.Op) Condition {
	return func(_ context.Context, m versioned.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m versioned.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m versioned.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m versioned.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk versioned.Hook, cond Condition) versioned.Hook {
	return func(next versioned.Mutator) versioned.Mutator {
		return versioned.MutateFunc(func(ctx context.Context, m versioned.Mutation) (versioned.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, versioned.Delete|versioned.Create)
func On(hk versioned.Hook, op versioned.Op) versioned.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, versioned.Update|versioned.UpdateOne)
func Unless(hk versioned.Hook, op versioned.Op) versioned.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) versioned.Hook {
	return func(versioned.Mutator) versioned.Mutator {
		return versioned.MutateFunc(func(context.Context, versioned.Mutation) (versioned.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []versioned.Hook {
//		return []versioned.Hook{
//			Reject(versioned.Delete|versioned.Update),
//		}
//	}
func Reject(op versioned.Op) versioned.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []versioned.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...versioned.Hook) Chain {
	return Chain{append([]versioned.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() versioned.Hook {
	return func(mutator versioned.Mutator) versioned.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...versioned.Hook) Chain {
	newHooks := make([]versioned.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
