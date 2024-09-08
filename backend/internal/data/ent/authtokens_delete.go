// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hay-kot/homebox/backend/internal/data/ent/authtokens"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// AuthTokensDelete is the builder for deleting a AuthTokens entity.
type AuthTokensDelete struct {
	config
	hooks    []Hook
	mutation *AuthTokensMutation
}

// Where appends a list predicates to the AuthTokensDelete builder.
func (atd *AuthTokensDelete) Where(ps ...predicate.AuthTokens) *AuthTokensDelete {
	atd.mutation.Where(ps...)
	return atd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atd *AuthTokensDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, atd.sqlExec, atd.mutation, atd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (atd *AuthTokensDelete) ExecX(ctx context.Context) int {
	n, err := atd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atd *AuthTokensDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(authtokens.Table, sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID))
	if ps := atd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, atd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	atd.mutation.done = true
	return affected, err
}

// AuthTokensDeleteOne is the builder for deleting a single AuthTokens entity.
type AuthTokensDeleteOne struct {
	atd *AuthTokensDelete
}

// Where appends a list predicates to the AuthTokensDelete builder.
func (atdo *AuthTokensDeleteOne) Where(ps ...predicate.AuthTokens) *AuthTokensDeleteOne {
	atdo.atd.mutation.Where(ps...)
	return atdo
}

// Exec executes the deletion query.
func (atdo *AuthTokensDeleteOne) Exec(ctx context.Context) error {
	n, err := atdo.atd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authtokens.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atdo *AuthTokensDeleteOne) ExecX(ctx context.Context) {
	if err := atdo.Exec(ctx); err != nil {
		panic(err)
	}
}