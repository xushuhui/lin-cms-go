// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"lin-cms-go/internal/data/model/linpermission"
	"lin-cms-go/internal/data/model/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinPermissionDelete is the builder for deleting a LinPermission entity.
type LinPermissionDelete struct {
	config
	hooks    []Hook
	mutation *LinPermissionMutation
}

// Where appends a list predicates to the LinPermissionDelete builder.
func (lpd *LinPermissionDelete) Where(ps ...predicate.LinPermission) *LinPermissionDelete {
	lpd.mutation.Where(ps...)
	return lpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lpd *LinPermissionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lpd.hooks) == 0 {
		affected, err = lpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinPermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lpd.mutation = mutation
			affected, err = lpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lpd.hooks) - 1; i >= 0; i-- {
			if lpd.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = lpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (lpd *LinPermissionDelete) ExecX(ctx context.Context) int {
	n, err := lpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lpd *LinPermissionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: linpermission.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: linpermission.FieldID,
			},
		},
	}
	if ps := lpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, lpd.driver, _spec)
}

// LinPermissionDeleteOne is the builder for deleting a single LinPermission entity.
type LinPermissionDeleteOne struct {
	lpd *LinPermissionDelete
}

// Exec executes the deletion query.
func (lpdo *LinPermissionDeleteOne) Exec(ctx context.Context) error {
	n, err := lpdo.lpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{linpermission.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lpdo *LinPermissionDeleteOne) ExecX(ctx context.Context) {
	lpdo.lpd.ExecX(ctx)
}
