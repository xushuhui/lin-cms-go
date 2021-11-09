// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"lin-cms-go/internal/data/model/lingroup"
	"lin-cms-go/internal/data/model/linpermission"
	"lin-cms-go/internal/data/model/linuser"
	"lin-cms-go/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinGroupUpdate is the builder for updating LinGroup entities.
type LinGroupUpdate struct {
	config
	hooks    []Hook
	mutation *LinGroupMutation
}

// Where appends a list predicates to the LinGroupUpdate builder.
func (lgu *LinGroupUpdate) Where(ps ...predicate.LinGroup) *LinGroupUpdate {
	lgu.mutation.Where(ps...)
	return lgu
}

// SetUpdateTime sets the "update_time" field.
func (lgu *LinGroupUpdate) SetUpdateTime(t time.Time) *LinGroupUpdate {
	lgu.mutation.SetUpdateTime(t)
	return lgu
}

// SetDeleteTime sets the "delete_time" field.
func (lgu *LinGroupUpdate) SetDeleteTime(t time.Time) *LinGroupUpdate {
	lgu.mutation.SetDeleteTime(t)
	return lgu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (lgu *LinGroupUpdate) SetNillableDeleteTime(t *time.Time) *LinGroupUpdate {
	if t != nil {
		lgu.SetDeleteTime(*t)
	}
	return lgu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (lgu *LinGroupUpdate) ClearDeleteTime() *LinGroupUpdate {
	lgu.mutation.ClearDeleteTime()
	return lgu
}

// SetName sets the "name" field.
func (lgu *LinGroupUpdate) SetName(s string) *LinGroupUpdate {
	lgu.mutation.SetName(s)
	return lgu
}

// SetInfo sets the "info" field.
func (lgu *LinGroupUpdate) SetInfo(s string) *LinGroupUpdate {
	lgu.mutation.SetInfo(s)
	return lgu
}

// SetLevel sets the "level" field.
func (lgu *LinGroupUpdate) SetLevel(i int8) *LinGroupUpdate {
	lgu.mutation.ResetLevel()
	lgu.mutation.SetLevel(i)
	return lgu
}

// AddLevel adds i to the "level" field.
func (lgu *LinGroupUpdate) AddLevel(i int8) *LinGroupUpdate {
	lgu.mutation.AddLevel(i)
	return lgu
}

// AddLinUserIDs adds the "lin_user" edge to the LinUser entity by IDs.
func (lgu *LinGroupUpdate) AddLinUserIDs(ids ...int) *LinGroupUpdate {
	lgu.mutation.AddLinUserIDs(ids...)
	return lgu
}

// AddLinUser adds the "lin_user" edges to the LinUser entity.
func (lgu *LinGroupUpdate) AddLinUser(l ...*LinUser) *LinGroupUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lgu.AddLinUserIDs(ids...)
}

// AddLinPermissionIDs adds the "lin_permission" edge to the LinPermission entity by IDs.
func (lgu *LinGroupUpdate) AddLinPermissionIDs(ids ...int) *LinGroupUpdate {
	lgu.mutation.AddLinPermissionIDs(ids...)
	return lgu
}

// AddLinPermission adds the "lin_permission" edges to the LinPermission entity.
func (lgu *LinGroupUpdate) AddLinPermission(l ...*LinPermission) *LinGroupUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lgu.AddLinPermissionIDs(ids...)
}

// Mutation returns the LinGroupMutation object of the builder.
func (lgu *LinGroupUpdate) Mutation() *LinGroupMutation {
	return lgu.mutation
}

// ClearLinUser clears all "lin_user" edges to the LinUser entity.
func (lgu *LinGroupUpdate) ClearLinUser() *LinGroupUpdate {
	lgu.mutation.ClearLinUser()
	return lgu
}

// RemoveLinUserIDs removes the "lin_user" edge to LinUser entities by IDs.
func (lgu *LinGroupUpdate) RemoveLinUserIDs(ids ...int) *LinGroupUpdate {
	lgu.mutation.RemoveLinUserIDs(ids...)
	return lgu
}

// RemoveLinUser removes "lin_user" edges to LinUser entities.
func (lgu *LinGroupUpdate) RemoveLinUser(l ...*LinUser) *LinGroupUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lgu.RemoveLinUserIDs(ids...)
}

// ClearLinPermission clears all "lin_permission" edges to the LinPermission entity.
func (lgu *LinGroupUpdate) ClearLinPermission() *LinGroupUpdate {
	lgu.mutation.ClearLinPermission()
	return lgu
}

// RemoveLinPermissionIDs removes the "lin_permission" edge to LinPermission entities by IDs.
func (lgu *LinGroupUpdate) RemoveLinPermissionIDs(ids ...int) *LinGroupUpdate {
	lgu.mutation.RemoveLinPermissionIDs(ids...)
	return lgu
}

// RemoveLinPermission removes "lin_permission" edges to LinPermission entities.
func (lgu *LinGroupUpdate) RemoveLinPermission(l ...*LinPermission) *LinGroupUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lgu.RemoveLinPermissionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lgu *LinGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	lgu.defaults()
	if len(lgu.hooks) == 0 {
		affected, err = lgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lgu.mutation = mutation
			affected, err = lgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lgu.hooks) - 1; i >= 0; i-- {
			if lgu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = lgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lgu *LinGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := lgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lgu *LinGroupUpdate) Exec(ctx context.Context) error {
	_, err := lgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lgu *LinGroupUpdate) ExecX(ctx context.Context) {
	if err := lgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lgu *LinGroupUpdate) defaults() {
	if _, ok := lgu.mutation.UpdateTime(); !ok {
		v := lingroup.UpdateDefaultUpdateTime()
		lgu.mutation.SetUpdateTime(v)
	}
}

func (lgu *LinGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lingroup.Table,
			Columns: lingroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lingroup.FieldID,
			},
		},
	}
	if ps := lgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lgu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lingroup.FieldUpdateTime,
		})
	}
	if value, ok := lgu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lingroup.FieldDeleteTime,
		})
	}
	if lgu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: lingroup.FieldDeleteTime,
		})
	}
	if value, ok := lgu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lingroup.FieldName,
		})
	}
	if value, ok := lgu.mutation.Info(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lingroup.FieldInfo,
		})
	}
	if value, ok := lgu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: lingroup.FieldLevel,
		})
	}
	if value, ok := lgu.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: lingroup.FieldLevel,
		})
	}
	if lgu.mutation.LinUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   lingroup.LinUserTable,
			Columns: lingroup.LinUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lgu.mutation.RemovedLinUserIDs(); len(nodes) > 0 && !lgu.mutation.LinUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   lingroup.LinUserTable,
			Columns: lingroup.LinUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lgu.mutation.LinUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   lingroup.LinUserTable,
			Columns: lingroup.LinUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lgu.mutation.LinPermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   lingroup.LinPermissionTable,
			Columns: lingroup.LinPermissionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linpermission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lgu.mutation.RemovedLinPermissionIDs(); len(nodes) > 0 && !lgu.mutation.LinPermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   lingroup.LinPermissionTable,
			Columns: lingroup.LinPermissionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linpermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lgu.mutation.LinPermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   lingroup.LinPermissionTable,
			Columns: lingroup.LinPermissionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linpermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lingroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LinGroupUpdateOne is the builder for updating a single LinGroup entity.
type LinGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LinGroupMutation
}

// SetUpdateTime sets the "update_time" field.
func (lguo *LinGroupUpdateOne) SetUpdateTime(t time.Time) *LinGroupUpdateOne {
	lguo.mutation.SetUpdateTime(t)
	return lguo
}

// SetDeleteTime sets the "delete_time" field.
func (lguo *LinGroupUpdateOne) SetDeleteTime(t time.Time) *LinGroupUpdateOne {
	lguo.mutation.SetDeleteTime(t)
	return lguo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (lguo *LinGroupUpdateOne) SetNillableDeleteTime(t *time.Time) *LinGroupUpdateOne {
	if t != nil {
		lguo.SetDeleteTime(*t)
	}
	return lguo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (lguo *LinGroupUpdateOne) ClearDeleteTime() *LinGroupUpdateOne {
	lguo.mutation.ClearDeleteTime()
	return lguo
}

// SetName sets the "name" field.
func (lguo *LinGroupUpdateOne) SetName(s string) *LinGroupUpdateOne {
	lguo.mutation.SetName(s)
	return lguo
}

// SetInfo sets the "info" field.
func (lguo *LinGroupUpdateOne) SetInfo(s string) *LinGroupUpdateOne {
	lguo.mutation.SetInfo(s)
	return lguo
}

// SetLevel sets the "level" field.
func (lguo *LinGroupUpdateOne) SetLevel(i int8) *LinGroupUpdateOne {
	lguo.mutation.ResetLevel()
	lguo.mutation.SetLevel(i)
	return lguo
}

// AddLevel adds i to the "level" field.
func (lguo *LinGroupUpdateOne) AddLevel(i int8) *LinGroupUpdateOne {
	lguo.mutation.AddLevel(i)
	return lguo
}

// AddLinUserIDs adds the "lin_user" edge to the LinUser entity by IDs.
func (lguo *LinGroupUpdateOne) AddLinUserIDs(ids ...int) *LinGroupUpdateOne {
	lguo.mutation.AddLinUserIDs(ids...)
	return lguo
}

// AddLinUser adds the "lin_user" edges to the LinUser entity.
func (lguo *LinGroupUpdateOne) AddLinUser(l ...*LinUser) *LinGroupUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lguo.AddLinUserIDs(ids...)
}

// AddLinPermissionIDs adds the "lin_permission" edge to the LinPermission entity by IDs.
func (lguo *LinGroupUpdateOne) AddLinPermissionIDs(ids ...int) *LinGroupUpdateOne {
	lguo.mutation.AddLinPermissionIDs(ids...)
	return lguo
}

// AddLinPermission adds the "lin_permission" edges to the LinPermission entity.
func (lguo *LinGroupUpdateOne) AddLinPermission(l ...*LinPermission) *LinGroupUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lguo.AddLinPermissionIDs(ids...)
}

// Mutation returns the LinGroupMutation object of the builder.
func (lguo *LinGroupUpdateOne) Mutation() *LinGroupMutation {
	return lguo.mutation
}

// ClearLinUser clears all "lin_user" edges to the LinUser entity.
func (lguo *LinGroupUpdateOne) ClearLinUser() *LinGroupUpdateOne {
	lguo.mutation.ClearLinUser()
	return lguo
}

// RemoveLinUserIDs removes the "lin_user" edge to LinUser entities by IDs.
func (lguo *LinGroupUpdateOne) RemoveLinUserIDs(ids ...int) *LinGroupUpdateOne {
	lguo.mutation.RemoveLinUserIDs(ids...)
	return lguo
}

// RemoveLinUser removes "lin_user" edges to LinUser entities.
func (lguo *LinGroupUpdateOne) RemoveLinUser(l ...*LinUser) *LinGroupUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lguo.RemoveLinUserIDs(ids...)
}

// ClearLinPermission clears all "lin_permission" edges to the LinPermission entity.
func (lguo *LinGroupUpdateOne) ClearLinPermission() *LinGroupUpdateOne {
	lguo.mutation.ClearLinPermission()
	return lguo
}

// RemoveLinPermissionIDs removes the "lin_permission" edge to LinPermission entities by IDs.
func (lguo *LinGroupUpdateOne) RemoveLinPermissionIDs(ids ...int) *LinGroupUpdateOne {
	lguo.mutation.RemoveLinPermissionIDs(ids...)
	return lguo
}

// RemoveLinPermission removes "lin_permission" edges to LinPermission entities.
func (lguo *LinGroupUpdateOne) RemoveLinPermission(l ...*LinPermission) *LinGroupUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lguo.RemoveLinPermissionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lguo *LinGroupUpdateOne) Select(field string, fields ...string) *LinGroupUpdateOne {
	lguo.fields = append([]string{field}, fields...)
	return lguo
}

// Save executes the query and returns the updated LinGroup entity.
func (lguo *LinGroupUpdateOne) Save(ctx context.Context) (*LinGroup, error) {
	var (
		err  error
		node *LinGroup
	)
	lguo.defaults()
	if len(lguo.hooks) == 0 {
		node, err = lguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lguo.mutation = mutation
			node, err = lguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lguo.hooks) - 1; i >= 0; i-- {
			if lguo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = lguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lguo *LinGroupUpdateOne) SaveX(ctx context.Context) *LinGroup {
	node, err := lguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lguo *LinGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := lguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lguo *LinGroupUpdateOne) ExecX(ctx context.Context) {
	if err := lguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lguo *LinGroupUpdateOne) defaults() {
	if _, ok := lguo.mutation.UpdateTime(); !ok {
		v := lingroup.UpdateDefaultUpdateTime()
		lguo.mutation.SetUpdateTime(v)
	}
}

func (lguo *LinGroupUpdateOne) sqlSave(ctx context.Context) (_node *LinGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lingroup.Table,
			Columns: lingroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lingroup.FieldID,
			},
		},
	}
	id, ok := lguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing LinGroup.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := lguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lingroup.FieldID)
		for _, f := range fields {
			if !lingroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != lingroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lguo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lingroup.FieldUpdateTime,
		})
	}
	if value, ok := lguo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lingroup.FieldDeleteTime,
		})
	}
	if lguo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: lingroup.FieldDeleteTime,
		})
	}
	if value, ok := lguo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lingroup.FieldName,
		})
	}
	if value, ok := lguo.mutation.Info(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lingroup.FieldInfo,
		})
	}
	if value, ok := lguo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: lingroup.FieldLevel,
		})
	}
	if value, ok := lguo.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: lingroup.FieldLevel,
		})
	}
	if lguo.mutation.LinUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   lingroup.LinUserTable,
			Columns: lingroup.LinUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lguo.mutation.RemovedLinUserIDs(); len(nodes) > 0 && !lguo.mutation.LinUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   lingroup.LinUserTable,
			Columns: lingroup.LinUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lguo.mutation.LinUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   lingroup.LinUserTable,
			Columns: lingroup.LinUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lguo.mutation.LinPermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   lingroup.LinPermissionTable,
			Columns: lingroup.LinPermissionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linpermission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lguo.mutation.RemovedLinPermissionIDs(); len(nodes) > 0 && !lguo.mutation.LinPermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   lingroup.LinPermissionTable,
			Columns: lingroup.LinPermissionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linpermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lguo.mutation.LinPermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   lingroup.LinPermissionTable,
			Columns: lingroup.LinPermissionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: linpermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LinGroup{config: lguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lingroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
