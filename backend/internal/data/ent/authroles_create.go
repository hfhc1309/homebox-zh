// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/authroles"
	"github.com/hay-kot/homebox/backend/internal/data/ent/authtokens"
)

// AuthRolesCreate is the builder for creating a AuthRoles entity.
type AuthRolesCreate struct {
	config
	mutation *AuthRolesMutation
	hooks    []Hook
}

// SetRole sets the "role" field.
func (arc *AuthRolesCreate) SetRole(a authroles.Role) *AuthRolesCreate {
	arc.mutation.SetRole(a)
	return arc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (arc *AuthRolesCreate) SetNillableRole(a *authroles.Role) *AuthRolesCreate {
	if a != nil {
		arc.SetRole(*a)
	}
	return arc
}

// SetTokenID sets the "token" edge to the AuthTokens entity by ID.
func (arc *AuthRolesCreate) SetTokenID(id uuid.UUID) *AuthRolesCreate {
	arc.mutation.SetTokenID(id)
	return arc
}

// SetNillableTokenID sets the "token" edge to the AuthTokens entity by ID if the given value is not nil.
func (arc *AuthRolesCreate) SetNillableTokenID(id *uuid.UUID) *AuthRolesCreate {
	if id != nil {
		arc = arc.SetTokenID(*id)
	}
	return arc
}

// SetToken sets the "token" edge to the AuthTokens entity.
func (arc *AuthRolesCreate) SetToken(a *AuthTokens) *AuthRolesCreate {
	return arc.SetTokenID(a.ID)
}

// Mutation returns the AuthRolesMutation object of the builder.
func (arc *AuthRolesCreate) Mutation() *AuthRolesMutation {
	return arc.mutation
}

// Save creates the AuthRoles in the database.
func (arc *AuthRolesCreate) Save(ctx context.Context) (*AuthRoles, error) {
	arc.defaults()
	return withHooks(ctx, arc.sqlSave, arc.mutation, arc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AuthRolesCreate) SaveX(ctx context.Context) *AuthRoles {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *AuthRolesCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *AuthRolesCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arc *AuthRolesCreate) defaults() {
	if _, ok := arc.mutation.Role(); !ok {
		v := authroles.DefaultRole
		arc.mutation.SetRole(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *AuthRolesCreate) check() error {
	if _, ok := arc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "AuthRoles.role"`)}
	}
	if v, ok := arc.mutation.Role(); ok {
		if err := authroles.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "AuthRoles.role": %w`, err)}
		}
	}
	return nil
}

func (arc *AuthRolesCreate) sqlSave(ctx context.Context) (*AuthRoles, error) {
	if err := arc.check(); err != nil {
		return nil, err
	}
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	arc.mutation.id = &_node.ID
	arc.mutation.done = true
	return _node, nil
}

func (arc *AuthRolesCreate) createSpec() (*AuthRoles, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthRoles{config: arc.config}
		_spec = sqlgraph.NewCreateSpec(authroles.Table, sqlgraph.NewFieldSpec(authroles.FieldID, field.TypeInt))
	)
	if value, ok := arc.mutation.Role(); ok {
		_spec.SetField(authroles.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if nodes := arc.mutation.TokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   authroles.TokenTable,
			Columns: []string{authroles.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(authtokens.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.auth_tokens_roles = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AuthRolesCreateBulk is the builder for creating many AuthRoles entities in bulk.
type AuthRolesCreateBulk struct {
	config
	err      error
	builders []*AuthRolesCreate
}

// Save creates the AuthRoles entities in the database.
func (arcb *AuthRolesCreateBulk) Save(ctx context.Context) ([]*AuthRoles, error) {
	if arcb.err != nil {
		return nil, arcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AuthRoles, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthRolesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AuthRolesCreateBulk) SaveX(ctx context.Context) []*AuthRoles {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *AuthRolesCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *AuthRolesCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}