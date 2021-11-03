// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"lin-cms-go/internal/data/model/lingroup"
	"lin-cms-go/internal/data/model/linpermission"
	"lin-cms-go/internal/data/model/linuser"
	"lin-cms-go/internal/data/model/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinGroupQuery is the builder for querying LinGroup entities.
type LinGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.LinGroup
	// eager-loading edges.
	withLinUser       *LinUserQuery
	withLinPermission *LinPermissionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LinGroupQuery builder.
func (lgq *LinGroupQuery) Where(ps ...predicate.LinGroup) *LinGroupQuery {
	lgq.predicates = append(lgq.predicates, ps...)
	return lgq
}

// Limit adds a limit step to the query.
func (lgq *LinGroupQuery) Limit(limit int) *LinGroupQuery {
	lgq.limit = &limit
	return lgq
}

// Offset adds an offset step to the query.
func (lgq *LinGroupQuery) Offset(offset int) *LinGroupQuery {
	lgq.offset = &offset
	return lgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lgq *LinGroupQuery) Unique(unique bool) *LinGroupQuery {
	lgq.unique = &unique
	return lgq
}

// Order adds an order step to the query.
func (lgq *LinGroupQuery) Order(o ...OrderFunc) *LinGroupQuery {
	lgq.order = append(lgq.order, o...)
	return lgq
}

// QueryLinUser chains the current query on the "lin_user" edge.
func (lgq *LinGroupQuery) QueryLinUser() *LinUserQuery {
	query := &LinUserQuery{config: lgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lingroup.Table, lingroup.FieldID, selector),
			sqlgraph.To(linuser.Table, linuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, lingroup.LinUserTable, lingroup.LinUserPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLinPermission chains the current query on the "lin_permission" edge.
func (lgq *LinGroupQuery) QueryLinPermission() *LinPermissionQuery {
	query := &LinPermissionQuery{config: lgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lingroup.Table, lingroup.FieldID, selector),
			sqlgraph.To(linpermission.Table, linpermission.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, lingroup.LinPermissionTable, lingroup.LinPermissionPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LinGroup entity from the query.
// Returns a *NotFoundError when no LinGroup was found.
func (lgq *LinGroupQuery) First(ctx context.Context) (*LinGroup, error) {
	nodes, err := lgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lingroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lgq *LinGroupQuery) FirstX(ctx context.Context) *LinGroup {
	node, err := lgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LinGroup ID from the query.
// Returns a *NotFoundError when no LinGroup ID was found.
func (lgq *LinGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lingroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lgq *LinGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := lgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Last returns the last LinGroup entity from the query.
// Returns a *NotFoundError when no LinGroup was found.
func (lgq *LinGroupQuery) Last(ctx context.Context) (*LinGroup, error) {
	nodes, err := lgq.All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lingroup.Label}
	}
	return nodes[len(nodes)-1], nil
}

// LastX is like Last, but panics if an error occurs.
func (lgq *LinGroupQuery) LastX(ctx context.Context) *LinGroup {
	node, err := lgq.Last(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// LastID returns the last LinGroup ID from the query.
// Returns a *NotFoundError when no LinGroup ID was found.
func (lgq *LinGroupQuery) LastID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lgq.IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lingroup.Label}
		return
	}
	return ids[len(ids)-1], nil
}

// LastIDX is like LastID, but panics if an error occurs.
func (lgq *LinGroupQuery) LastIDX(ctx context.Context) int {
	id, err := lgq.LastID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LinGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one LinGroup entity is not found.
// Returns a *NotFoundError when no LinGroup entities are found.
func (lgq *LinGroupQuery) Only(ctx context.Context) (*LinGroup, error) {
	nodes, err := lgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lingroup.Label}
	default:
		return nil, &NotSingularError{lingroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lgq *LinGroupQuery) OnlyX(ctx context.Context) *LinGroup {
	node, err := lgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LinGroup ID in the query.
// Returns a *NotSingularError when exactly one LinGroup ID is not found.
// Returns a *NotFoundError when no entities are found.
func (lgq *LinGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lingroup.Label}
	default:
		err = &NotSingularError{lingroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lgq *LinGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := lgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LinGroups.
func (lgq *LinGroupQuery) All(ctx context.Context) ([]*LinGroup, error) {
	if err := lgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lgq *LinGroupQuery) AllX(ctx context.Context) []*LinGroup {
	nodes, err := lgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LinGroup IDs.
func (lgq *LinGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := lgq.Select(lingroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lgq *LinGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := lgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lgq *LinGroupQuery) Count(ctx context.Context) (int, error) {
	if err := lgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lgq *LinGroupQuery) CountX(ctx context.Context) int {
	count, err := lgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lgq *LinGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := lgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lgq *LinGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := lgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LinGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lgq *LinGroupQuery) Clone() *LinGroupQuery {
	if lgq == nil {
		return nil
	}
	return &LinGroupQuery{
		config:            lgq.config,
		limit:             lgq.limit,
		offset:            lgq.offset,
		order:             append([]OrderFunc{}, lgq.order...),
		predicates:        append([]predicate.LinGroup{}, lgq.predicates...),
		withLinUser:       lgq.withLinUser.Clone(),
		withLinPermission: lgq.withLinPermission.Clone(),
		// clone intermediate query.
		sql:  lgq.sql.Clone(),
		path: lgq.path,
	}
}

// WithLinUser tells the query-builder to eager-load the nodes that are connected to
// the "lin_user" edge. The optional arguments are used to configure the query builder of the edge.
func (lgq *LinGroupQuery) WithLinUser(opts ...func(*LinUserQuery)) *LinGroupQuery {
	query := &LinUserQuery{config: lgq.config}
	for _, opt := range opts {
		opt(query)
	}
	lgq.withLinUser = query
	return lgq
}

// WithLinPermission tells the query-builder to eager-load the nodes that are connected to
// the "lin_permission" edge. The optional arguments are used to configure the query builder of the edge.
func (lgq *LinGroupQuery) WithLinPermission(opts ...func(*LinPermissionQuery)) *LinGroupQuery {
	query := &LinPermissionQuery{config: lgq.config}
	for _, opt := range opts {
		opt(query)
	}
	lgq.withLinPermission = query
	return lgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LinGroup.Query().
//		GroupBy(lingroup.FieldName).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (lgq *LinGroupQuery) GroupBy(field string, fields ...string) *LinGroupGroupBy {
	group := &LinGroupGroupBy{config: lgq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lgq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.LinGroup.Query().
//		Select(lingroup.FieldName).
//		Scan(ctx, &v)
//
func (lgq *LinGroupQuery) Select(fields ...string) *LinGroupSelect {
	lgq.fields = append(lgq.fields, fields...)
	return &LinGroupSelect{LinGroupQuery: lgq}
}

func (lgq *LinGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range lgq.fields {
		if !lingroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if lgq.path != nil {
		prev, err := lgq.path(ctx)
		if err != nil {
			return err
		}
		lgq.sql = prev
	}
	return nil
}

func (lgq *LinGroupQuery) sqlAll(ctx context.Context) ([]*LinGroup, error) {
	var (
		nodes       = []*LinGroup{}
		_spec       = lgq.querySpec()
		loadedTypes = [2]bool{
			lgq.withLinUser != nil,
			lgq.withLinPermission != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &LinGroup{config: lgq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("model: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, lgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := lgq.withLinUser; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*LinGroup, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.LinUser = []*LinUser{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*LinGroup)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   lingroup.LinUserTable,
				Columns: lingroup.LinUserPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(lingroup.LinUserPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, lgq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "lin_user": %w`, err)
		}
		query.Where(linuser.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "lin_user" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.LinUser = append(nodes[i].Edges.LinUser, n)
			}
		}
	}

	if query := lgq.withLinPermission; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*LinGroup, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.LinPermission = []*LinPermission{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*LinGroup)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   lingroup.LinPermissionTable,
				Columns: lingroup.LinPermissionPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(lingroup.LinPermissionPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, lgq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "lin_permission": %w`, err)
		}
		query.Where(linpermission.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "lin_permission" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.LinPermission = append(nodes[i].Edges.LinPermission, n)
			}
		}
	}

	return nodes, nil
}

func (lgq *LinGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lgq.querySpec()
	return sqlgraph.CountNodes(ctx, lgq.driver, _spec)
}

func (lgq *LinGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := lgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (lgq *LinGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lingroup.Table,
			Columns: lingroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lingroup.FieldID,
			},
		},
		From:   lgq.sql,
		Unique: true,
	}
	if unique := lgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lingroup.FieldID)
		for i := range fields {
			if fields[i] != lingroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lgq *LinGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lgq.driver.Dialect())
	t1 := builder.Table(lingroup.Table)
	columns := lgq.fields
	if len(columns) == 0 {
		columns = lingroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lgq.sql != nil {
		selector = lgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range lgq.predicates {
		p(selector)
	}
	for _, p := range lgq.order {
		p(selector)
	}
	if offset := lgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LinGroupGroupBy is the group-by builder for LinGroup entities.
type LinGroupGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lggb *LinGroupGroupBy) Aggregate(fns ...AggregateFunc) *LinGroupGroupBy {
	lggb.fns = append(lggb.fns, fns...)
	return lggb
}

// Scan applies the group-by query and scans the result into the given value.
func (lggb *LinGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lggb.path(ctx)
	if err != nil {
		return err
	}
	lggb.sql = query
	return lggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lggb *LinGroupGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := lggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (lggb *LinGroupGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(lggb.fields) > 1 {
		return nil, errors.New("model: LinGroupGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := lggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lggb *LinGroupGroupBy) StringsX(ctx context.Context) []string {
	v, err := lggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lggb *LinGroupGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lingroup.Label}
	default:
		err = fmt.Errorf("model: LinGroupGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lggb *LinGroupGroupBy) StringX(ctx context.Context) string {
	v, err := lggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (lggb *LinGroupGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(lggb.fields) > 1 {
		return nil, errors.New("model: LinGroupGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := lggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lggb *LinGroupGroupBy) IntsX(ctx context.Context) []int {
	v, err := lggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lggb *LinGroupGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lingroup.Label}
	default:
		err = fmt.Errorf("model: LinGroupGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lggb *LinGroupGroupBy) IntX(ctx context.Context) int {
	v, err := lggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (lggb *LinGroupGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(lggb.fields) > 1 {
		return nil, errors.New("model: LinGroupGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := lggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lggb *LinGroupGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := lggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lggb *LinGroupGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lingroup.Label}
	default:
		err = fmt.Errorf("model: LinGroupGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lggb *LinGroupGroupBy) Float64X(ctx context.Context) float64 {
	v, err := lggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (lggb *LinGroupGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(lggb.fields) > 1 {
		return nil, errors.New("model: LinGroupGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := lggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lggb *LinGroupGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := lggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lggb *LinGroupGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lingroup.Label}
	default:
		err = fmt.Errorf("model: LinGroupGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lggb *LinGroupGroupBy) BoolX(ctx context.Context) bool {
	v, err := lggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lggb *LinGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lggb.fields {
		if !lingroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lggb *LinGroupGroupBy) sqlQuery() *sql.Selector {
	selector := lggb.sql.Select()
	aggregation := make([]string, 0, len(lggb.fns))
	for _, fn := range lggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lggb.fields)+len(lggb.fns))
		for _, f := range lggb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lggb.fields...)...)
}

// LinGroupSelect is the builder for selecting fields of LinGroup entities.
type LinGroupSelect struct {
	*LinGroupQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (lgs *LinGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := lgs.prepareQuery(ctx); err != nil {
		return err
	}
	lgs.sql = lgs.LinGroupQuery.sqlQuery(ctx)
	return lgs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lgs *LinGroupSelect) ScanX(ctx context.Context, v interface{}) {
	if err := lgs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (lgs *LinGroupSelect) Strings(ctx context.Context) ([]string, error) {
	if len(lgs.fields) > 1 {
		return nil, errors.New("model: LinGroupSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := lgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lgs *LinGroupSelect) StringsX(ctx context.Context) []string {
	v, err := lgs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (lgs *LinGroupSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lgs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lingroup.Label}
	default:
		err = fmt.Errorf("model: LinGroupSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lgs *LinGroupSelect) StringX(ctx context.Context) string {
	v, err := lgs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (lgs *LinGroupSelect) Ints(ctx context.Context) ([]int, error) {
	if len(lgs.fields) > 1 {
		return nil, errors.New("model: LinGroupSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := lgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lgs *LinGroupSelect) IntsX(ctx context.Context) []int {
	v, err := lgs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (lgs *LinGroupSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lgs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lingroup.Label}
	default:
		err = fmt.Errorf("model: LinGroupSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lgs *LinGroupSelect) IntX(ctx context.Context) int {
	v, err := lgs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (lgs *LinGroupSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(lgs.fields) > 1 {
		return nil, errors.New("model: LinGroupSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := lgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lgs *LinGroupSelect) Float64sX(ctx context.Context) []float64 {
	v, err := lgs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (lgs *LinGroupSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lgs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lingroup.Label}
	default:
		err = fmt.Errorf("model: LinGroupSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lgs *LinGroupSelect) Float64X(ctx context.Context) float64 {
	v, err := lgs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (lgs *LinGroupSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(lgs.fields) > 1 {
		return nil, errors.New("model: LinGroupSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := lgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lgs *LinGroupSelect) BoolsX(ctx context.Context) []bool {
	v, err := lgs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (lgs *LinGroupSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lgs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{lingroup.Label}
	default:
		err = fmt.Errorf("model: LinGroupSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lgs *LinGroupSelect) BoolX(ctx context.Context) bool {
	v, err := lgs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lgs *LinGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := lgs.sql.Query()
	if err := lgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}