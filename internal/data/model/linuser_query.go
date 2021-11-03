// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"lin-cms-go/internal/data/model/lingroup"
	"lin-cms-go/internal/data/model/linuser"
	"lin-cms-go/internal/data/model/linuseridentiy"
	"lin-cms-go/internal/data/model/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinUserQuery is the builder for querying LinUser entities.
type LinUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.LinUser
	// eager-loading edges.
	withLinUserIdentiy *LinUserIdentiyQuery
	withLinGroup       *LinGroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LinUserQuery builder.
func (luq *LinUserQuery) Where(ps ...predicate.LinUser) *LinUserQuery {
	luq.predicates = append(luq.predicates, ps...)
	return luq
}

// Limit adds a limit step to the query.
func (luq *LinUserQuery) Limit(limit int) *LinUserQuery {
	luq.limit = &limit
	return luq
}

// Offset adds an offset step to the query.
func (luq *LinUserQuery) Offset(offset int) *LinUserQuery {
	luq.offset = &offset
	return luq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (luq *LinUserQuery) Unique(unique bool) *LinUserQuery {
	luq.unique = &unique
	return luq
}

// Order adds an order step to the query.
func (luq *LinUserQuery) Order(o ...OrderFunc) *LinUserQuery {
	luq.order = append(luq.order, o...)
	return luq
}

// QueryLinUserIdentiy chains the current query on the "lin_user_identiy" edge.
func (luq *LinUserQuery) QueryLinUserIdentiy() *LinUserIdentiyQuery {
	query := &LinUserIdentiyQuery{config: luq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := luq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := luq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(linuser.Table, linuser.FieldID, selector),
			sqlgraph.To(linuseridentiy.Table, linuseridentiy.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, linuser.LinUserIdentiyTable, linuser.LinUserIdentiyColumn),
		)
		fromU = sqlgraph.SetNeighbors(luq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLinGroup chains the current query on the "lin_group" edge.
func (luq *LinUserQuery) QueryLinGroup() *LinGroupQuery {
	query := &LinGroupQuery{config: luq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := luq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := luq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(linuser.Table, linuser.FieldID, selector),
			sqlgraph.To(lingroup.Table, lingroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, linuser.LinGroupTable, linuser.LinGroupPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(luq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LinUser entity from the query.
// Returns a *NotFoundError when no LinUser was found.
func (luq *LinUserQuery) First(ctx context.Context) (*LinUser, error) {
	nodes, err := luq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{linuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (luq *LinUserQuery) FirstX(ctx context.Context) *LinUser {
	node, err := luq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LinUser ID from the query.
// Returns a *NotFoundError when no LinUser ID was found.
func (luq *LinUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = luq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{linuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (luq *LinUserQuery) FirstIDX(ctx context.Context) int {
	id, err := luq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Last returns the last LinUser entity from the query.
// Returns a *NotFoundError when no LinUser was found.
func (luq *LinUserQuery) Last(ctx context.Context) (*LinUser, error) {
	nodes, err := luq.All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{linuser.Label}
	}
	return nodes[len(nodes)-1], nil
}

// LastX is like Last, but panics if an error occurs.
func (luq *LinUserQuery) LastX(ctx context.Context) *LinUser {
	node, err := luq.Last(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// LastID returns the last LinUser ID from the query.
// Returns a *NotFoundError when no LinUser ID was found.
func (luq *LinUserQuery) LastID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = luq.IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{linuser.Label}
		return
	}
	return ids[len(ids)-1], nil
}

// LastIDX is like LastID, but panics if an error occurs.
func (luq *LinUserQuery) LastIDX(ctx context.Context) int {
	id, err := luq.LastID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LinUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one LinUser entity is not found.
// Returns a *NotFoundError when no LinUser entities are found.
func (luq *LinUserQuery) Only(ctx context.Context) (*LinUser, error) {
	nodes, err := luq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{linuser.Label}
	default:
		return nil, &NotSingularError{linuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (luq *LinUserQuery) OnlyX(ctx context.Context) *LinUser {
	node, err := luq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LinUser ID in the query.
// Returns a *NotSingularError when exactly one LinUser ID is not found.
// Returns a *NotFoundError when no entities are found.
func (luq *LinUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = luq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{linuser.Label}
	default:
		err = &NotSingularError{linuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (luq *LinUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := luq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LinUsers.
func (luq *LinUserQuery) All(ctx context.Context) ([]*LinUser, error) {
	if err := luq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return luq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (luq *LinUserQuery) AllX(ctx context.Context) []*LinUser {
	nodes, err := luq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LinUser IDs.
func (luq *LinUserQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := luq.Select(linuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (luq *LinUserQuery) IDsX(ctx context.Context) []int {
	ids, err := luq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (luq *LinUserQuery) Count(ctx context.Context) (int, error) {
	if err := luq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return luq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (luq *LinUserQuery) CountX(ctx context.Context) int {
	count, err := luq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (luq *LinUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := luq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return luq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (luq *LinUserQuery) ExistX(ctx context.Context) bool {
	exist, err := luq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LinUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (luq *LinUserQuery) Clone() *LinUserQuery {
	if luq == nil {
		return nil
	}
	return &LinUserQuery{
		config:             luq.config,
		limit:              luq.limit,
		offset:             luq.offset,
		order:              append([]OrderFunc{}, luq.order...),
		predicates:         append([]predicate.LinUser{}, luq.predicates...),
		withLinUserIdentiy: luq.withLinUserIdentiy.Clone(),
		withLinGroup:       luq.withLinGroup.Clone(),
		// clone intermediate query.
		sql:  luq.sql.Clone(),
		path: luq.path,
	}
}

// WithLinUserIdentiy tells the query-builder to eager-load the nodes that are connected to
// the "lin_user_identiy" edge. The optional arguments are used to configure the query builder of the edge.
func (luq *LinUserQuery) WithLinUserIdentiy(opts ...func(*LinUserIdentiyQuery)) *LinUserQuery {
	query := &LinUserIdentiyQuery{config: luq.config}
	for _, opt := range opts {
		opt(query)
	}
	luq.withLinUserIdentiy = query
	return luq
}

// WithLinGroup tells the query-builder to eager-load the nodes that are connected to
// the "lin_group" edge. The optional arguments are used to configure the query builder of the edge.
func (luq *LinUserQuery) WithLinGroup(opts ...func(*LinGroupQuery)) *LinUserQuery {
	query := &LinGroupQuery{config: luq.config}
	for _, opt := range opts {
		opt(query)
	}
	luq.withLinGroup = query
	return luq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LinUser.Query().
//		GroupBy(linuser.FieldCreateTime).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (luq *LinUserQuery) GroupBy(field string, fields ...string) *LinUserGroupBy {
	group := &LinUserGroupBy{config: luq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := luq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return luq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.LinUser.Query().
//		Select(linuser.FieldCreateTime).
//		Scan(ctx, &v)
//
func (luq *LinUserQuery) Select(fields ...string) *LinUserSelect {
	luq.fields = append(luq.fields, fields...)
	return &LinUserSelect{LinUserQuery: luq}
}

func (luq *LinUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range luq.fields {
		if !linuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if luq.path != nil {
		prev, err := luq.path(ctx)
		if err != nil {
			return err
		}
		luq.sql = prev
	}
	return nil
}

func (luq *LinUserQuery) sqlAll(ctx context.Context) ([]*LinUser, error) {
	var (
		nodes       = []*LinUser{}
		_spec       = luq.querySpec()
		loadedTypes = [2]bool{
			luq.withLinUserIdentiy != nil,
			luq.withLinGroup != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &LinUser{config: luq.config}
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
	if err := sqlgraph.QueryNodes(ctx, luq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := luq.withLinUserIdentiy; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*LinUser)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.LinUserIdentiy = []*LinUserIdentiy{}
		}
		query.withFKs = true
		query.Where(predicate.LinUserIdentiy(func(s *sql.Selector) {
			s.Where(sql.InValues(linuser.LinUserIdentiyColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.lin_user_lin_user_identiy
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "lin_user_lin_user_identiy" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "lin_user_lin_user_identiy" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.LinUserIdentiy = append(node.Edges.LinUserIdentiy, n)
		}
	}

	if query := luq.withLinGroup; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*LinUser, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.LinGroup = []*LinGroup{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*LinUser)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   linuser.LinGroupTable,
				Columns: linuser.LinGroupPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(linuser.LinGroupPrimaryKey[1], fks...))
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
		if err := sqlgraph.QueryEdges(ctx, luq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "lin_group": %w`, err)
		}
		query.Where(lingroup.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "lin_group" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.LinGroup = append(nodes[i].Edges.LinGroup, n)
			}
		}
	}

	return nodes, nil
}

func (luq *LinUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := luq.querySpec()
	return sqlgraph.CountNodes(ctx, luq.driver, _spec)
}

func (luq *LinUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := luq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (luq *LinUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   linuser.Table,
			Columns: linuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: linuser.FieldID,
			},
		},
		From:   luq.sql,
		Unique: true,
	}
	if unique := luq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := luq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, linuser.FieldID)
		for i := range fields {
			if fields[i] != linuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := luq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := luq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := luq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := luq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (luq *LinUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(luq.driver.Dialect())
	t1 := builder.Table(linuser.Table)
	columns := luq.fields
	if len(columns) == 0 {
		columns = linuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if luq.sql != nil {
		selector = luq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range luq.predicates {
		p(selector)
	}
	for _, p := range luq.order {
		p(selector)
	}
	if offset := luq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := luq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LinUserGroupBy is the group-by builder for LinUser entities.
type LinUserGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lugb *LinUserGroupBy) Aggregate(fns ...AggregateFunc) *LinUserGroupBy {
	lugb.fns = append(lugb.fns, fns...)
	return lugb
}

// Scan applies the group-by query and scans the result into the given value.
func (lugb *LinUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lugb.path(ctx)
	if err != nil {
		return err
	}
	lugb.sql = query
	return lugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lugb *LinUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := lugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (lugb *LinUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(lugb.fields) > 1 {
		return nil, errors.New("model: LinUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := lugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lugb *LinUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := lugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lugb *LinUserGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linuser.Label}
	default:
		err = fmt.Errorf("model: LinUserGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lugb *LinUserGroupBy) StringX(ctx context.Context) string {
	v, err := lugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (lugb *LinUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(lugb.fields) > 1 {
		return nil, errors.New("model: LinUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := lugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lugb *LinUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := lugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lugb *LinUserGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linuser.Label}
	default:
		err = fmt.Errorf("model: LinUserGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lugb *LinUserGroupBy) IntX(ctx context.Context) int {
	v, err := lugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (lugb *LinUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(lugb.fields) > 1 {
		return nil, errors.New("model: LinUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := lugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lugb *LinUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := lugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lugb *LinUserGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linuser.Label}
	default:
		err = fmt.Errorf("model: LinUserGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lugb *LinUserGroupBy) Float64X(ctx context.Context) float64 {
	v, err := lugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (lugb *LinUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(lugb.fields) > 1 {
		return nil, errors.New("model: LinUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := lugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lugb *LinUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := lugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (lugb *LinUserGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linuser.Label}
	default:
		err = fmt.Errorf("model: LinUserGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lugb *LinUserGroupBy) BoolX(ctx context.Context) bool {
	v, err := lugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lugb *LinUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lugb.fields {
		if !linuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lugb *LinUserGroupBy) sqlQuery() *sql.Selector {
	selector := lugb.sql.Select()
	aggregation := make([]string, 0, len(lugb.fns))
	for _, fn := range lugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lugb.fields)+len(lugb.fns))
		for _, f := range lugb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lugb.fields...)...)
}

// LinUserSelect is the builder for selecting fields of LinUser entities.
type LinUserSelect struct {
	*LinUserQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (lus *LinUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := lus.prepareQuery(ctx); err != nil {
		return err
	}
	lus.sql = lus.LinUserQuery.sqlQuery(ctx)
	return lus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lus *LinUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := lus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (lus *LinUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(lus.fields) > 1 {
		return nil, errors.New("model: LinUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := lus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lus *LinUserSelect) StringsX(ctx context.Context) []string {
	v, err := lus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (lus *LinUserSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linuser.Label}
	default:
		err = fmt.Errorf("model: LinUserSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lus *LinUserSelect) StringX(ctx context.Context) string {
	v, err := lus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (lus *LinUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(lus.fields) > 1 {
		return nil, errors.New("model: LinUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := lus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lus *LinUserSelect) IntsX(ctx context.Context) []int {
	v, err := lus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (lus *LinUserSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linuser.Label}
	default:
		err = fmt.Errorf("model: LinUserSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lus *LinUserSelect) IntX(ctx context.Context) int {
	v, err := lus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (lus *LinUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(lus.fields) > 1 {
		return nil, errors.New("model: LinUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := lus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lus *LinUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := lus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (lus *LinUserSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linuser.Label}
	default:
		err = fmt.Errorf("model: LinUserSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lus *LinUserSelect) Float64X(ctx context.Context) float64 {
	v, err := lus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (lus *LinUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(lus.fields) > 1 {
		return nil, errors.New("model: LinUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := lus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lus *LinUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := lus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (lus *LinUserSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{linuser.Label}
	default:
		err = fmt.Errorf("model: LinUserSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lus *LinUserSelect) BoolX(ctx context.Context) bool {
	v, err := lus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lus *LinUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := lus.sql.Query()
	if err := lus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}