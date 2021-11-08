// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"lin-cms-go/internal/data/model/linlog"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinLogCreate is the builder for creating a LinLog entity.
type LinLogCreate struct {
	config
	mutation *LinLogMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (llc *LinLogCreate) SetCreateTime(t time.Time) *LinLogCreate {
	llc.mutation.SetCreateTime(t)
	return llc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (llc *LinLogCreate) SetNillableCreateTime(t *time.Time) *LinLogCreate {
	if t != nil {
		llc.SetCreateTime(*t)
	}
	return llc
}

// SetUpdateTime sets the "update_time" field.
func (llc *LinLogCreate) SetUpdateTime(t time.Time) *LinLogCreate {
	llc.mutation.SetUpdateTime(t)
	return llc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (llc *LinLogCreate) SetNillableUpdateTime(t *time.Time) *LinLogCreate {
	if t != nil {
		llc.SetUpdateTime(*t)
	}
	return llc
}

// SetDeleteTime sets the "delete_time" field.
func (llc *LinLogCreate) SetDeleteTime(t time.Time) *LinLogCreate {
	llc.mutation.SetDeleteTime(t)
	return llc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (llc *LinLogCreate) SetNillableDeleteTime(t *time.Time) *LinLogCreate {
	if t != nil {
		llc.SetDeleteTime(*t)
	}
	return llc
}

// SetMessage sets the "message" field.
func (llc *LinLogCreate) SetMessage(s string) *LinLogCreate {
	llc.mutation.SetMessage(s)
	return llc
}

// SetUserID sets the "user_id" field.
func (llc *LinLogCreate) SetUserID(i int) *LinLogCreate {
	llc.mutation.SetUserID(i)
	return llc
}

// SetUsername sets the "username" field.
func (llc *LinLogCreate) SetUsername(s string) *LinLogCreate {
	llc.mutation.SetUsername(s)
	return llc
}

// SetStatusCode sets the "status_code" field.
func (llc *LinLogCreate) SetStatusCode(i int) *LinLogCreate {
	llc.mutation.SetStatusCode(i)
	return llc
}

// SetMethod sets the "method" field.
func (llc *LinLogCreate) SetMethod(s string) *LinLogCreate {
	llc.mutation.SetMethod(s)
	return llc
}

// SetPath sets the "path" field.
func (llc *LinLogCreate) SetPath(s string) *LinLogCreate {
	llc.mutation.SetPath(s)
	return llc
}

// SetPermission sets the "permission" field.
func (llc *LinLogCreate) SetPermission(s string) *LinLogCreate {
	llc.mutation.SetPermission(s)
	return llc
}

// Mutation returns the LinLogMutation object of the builder.
func (llc *LinLogCreate) Mutation() *LinLogMutation {
	return llc.mutation
}

// Save creates the LinLog in the database.
func (llc *LinLogCreate) Save(ctx context.Context) (*LinLog, error) {
	var (
		err  error
		node *LinLog
	)
	llc.defaults()
	if len(llc.hooks) == 0 {
		if err = llc.check(); err != nil {
			return nil, err
		}
		node, err = llc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = llc.check(); err != nil {
				return nil, err
			}
			llc.mutation = mutation
			if node, err = llc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(llc.hooks) - 1; i >= 0; i-- {
			if llc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = llc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, llc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (llc *LinLogCreate) SaveX(ctx context.Context) *LinLog {
	v, err := llc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (llc *LinLogCreate) Exec(ctx context.Context) error {
	_, err := llc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (llc *LinLogCreate) ExecX(ctx context.Context) {
	if err := llc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (llc *LinLogCreate) defaults() {
	if _, ok := llc.mutation.CreateTime(); !ok {
		v := linlog.DefaultCreateTime()
		llc.mutation.SetCreateTime(v)
	}
	if _, ok := llc.mutation.UpdateTime(); !ok {
		v := linlog.DefaultUpdateTime()
		llc.mutation.SetUpdateTime(v)
	}
	if _, ok := llc.mutation.DeleteTime(); !ok {
		v := linlog.DefaultDeleteTime()
		llc.mutation.SetDeleteTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (llc *LinLogCreate) check() error {
	if _, ok := llc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "create_time"`)}
	}
	if _, ok := llc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "update_time"`)}
	}
	if _, ok := llc.mutation.DeleteTime(); !ok {
		return &ValidationError{Name: "delete_time", err: errors.New(`model: missing required field "delete_time"`)}
	}
	if _, ok := llc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`model: missing required field "message"`)}
	}
	if _, ok := llc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`model: missing required field "user_id"`)}
	}
	if _, ok := llc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`model: missing required field "username"`)}
	}
	if _, ok := llc.mutation.StatusCode(); !ok {
		return &ValidationError{Name: "status_code", err: errors.New(`model: missing required field "status_code"`)}
	}
	if _, ok := llc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`model: missing required field "method"`)}
	}
	if _, ok := llc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`model: missing required field "path"`)}
	}
	if _, ok := llc.mutation.Permission(); !ok {
		return &ValidationError{Name: "permission", err: errors.New(`model: missing required field "permission"`)}
	}
	return nil
}

func (llc *LinLogCreate) sqlSave(ctx context.Context) (*LinLog, error) {
	_node, _spec := llc.createSpec()
	if err := sqlgraph.CreateNode(ctx, llc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (llc *LinLogCreate) createSpec() (*LinLog, *sqlgraph.CreateSpec) {
	var (
		_node = &LinLog{config: llc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: linlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: linlog.FieldID,
			},
		}
	)
	if value, ok := llc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: linlog.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := llc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: linlog.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := llc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: linlog.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := llc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: linlog.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := llc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: linlog.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := llc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: linlog.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := llc.mutation.StatusCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: linlog.FieldStatusCode,
		})
		_node.StatusCode = value
	}
	if value, ok := llc.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: linlog.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := llc.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: linlog.FieldPath,
		})
		_node.Path = value
	}
	if value, ok := llc.mutation.Permission(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: linlog.FieldPermission,
		})
		_node.Permission = value
	}
	return _node, _spec
}

// LinLogCreateBulk is the builder for creating many LinLog entities in bulk.
type LinLogCreateBulk struct {
	config
	builders []*LinLogCreate
}

// Save creates the LinLog entities in the database.
func (llcb *LinLogCreateBulk) Save(ctx context.Context) ([]*LinLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(llcb.builders))
	nodes := make([]*LinLog, len(llcb.builders))
	mutators := make([]Mutator, len(llcb.builders))
	for i := range llcb.builders {
		func(i int, root context.Context) {
			builder := llcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, llcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, llcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, llcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (llcb *LinLogCreateBulk) SaveX(ctx context.Context) []*LinLog {
	v, err := llcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (llcb *LinLogCreateBulk) Exec(ctx context.Context) error {
	_, err := llcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (llcb *LinLogCreateBulk) ExecX(ctx context.Context) {
	if err := llcb.Exec(ctx); err != nil {
		panic(err)
	}
}
