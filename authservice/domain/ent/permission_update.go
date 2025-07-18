// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent/permission"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent/predicate"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetPermissionName sets the "permissionName" field.
func (pu *PermissionUpdate) SetPermissionName(s string) *PermissionUpdate {
	pu.mutation.SetPermissionName(s)
	return pu
}

// SetNillablePermissionName sets the "permissionName" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillablePermissionName(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetPermissionName(*s)
	}
	return pu
}

// SetURI sets the "uri" field.
func (pu *PermissionUpdate) SetURI(s string) *PermissionUpdate {
	pu.mutation.SetURI(s)
	return pu
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableURI(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetURI(*s)
	}
	return pu
}

// SetMethod sets the "method" field.
func (pu *PermissionUpdate) SetMethod(s string) *PermissionUpdate {
	pu.mutation.SetMethod(s)
	return pu
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableMethod(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetMethod(*s)
	}
	return pu
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (pu *PermissionUpdate) AddPermissionIDs(ids ...uuid.UUID) *PermissionUpdate {
	pu.mutation.AddPermissionIDs(ids...)
	return pu
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (pu *PermissionUpdate) AddPermissions(p ...*Permission) *PermissionUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPermissionIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearPermissions clears all "permissions" edges to the Permission entity.
func (pu *PermissionUpdate) ClearPermissions() *PermissionUpdate {
	pu.mutation.ClearPermissions()
	return pu
}

// RemovePermissionIDs removes the "permissions" edge to Permission entities by IDs.
func (pu *PermissionUpdate) RemovePermissionIDs(ids ...uuid.UUID) *PermissionUpdate {
	pu.mutation.RemovePermissionIDs(ids...)
	return pu
}

// RemovePermissions removes "permissions" edges to Permission entities.
func (pu *PermissionUpdate) RemovePermissions(p ...*Permission) *PermissionUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePermissionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PermissionUpdate) check() error {
	if v, ok := pu.mutation.PermissionName(); ok {
		if err := permission.PermissionNameValidator(v); err != nil {
			return &ValidationError{Name: "permissionName", err: fmt.Errorf(`ent: validator failed for field "Permission.permissionName": %w`, err)}
		}
	}
	if v, ok := pu.mutation.URI(); ok {
		if err := permission.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Permission.uri": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Method(); ok {
		if err := permission.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "Permission.method": %w`, err)}
		}
	}
	return nil
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.PermissionName(); ok {
		_spec.SetField(permission.FieldPermissionName, field.TypeString, value)
	}
	if value, ok := pu.mutation.URI(); ok {
		_spec.SetField(permission.FieldURI, field.TypeString, value)
	}
	if value, ok := pu.mutation.Method(); ok {
		_spec.SetField(permission.FieldMethod, field.TypeString, value)
	}
	if pu.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.PermissionsTable,
			Columns: permission.PermissionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !pu.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.PermissionsTable,
			Columns: permission.PermissionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.PermissionsTable,
			Columns: permission.PermissionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionMutation
}

// SetPermissionName sets the "permissionName" field.
func (puo *PermissionUpdateOne) SetPermissionName(s string) *PermissionUpdateOne {
	puo.mutation.SetPermissionName(s)
	return puo
}

// SetNillablePermissionName sets the "permissionName" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillablePermissionName(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetPermissionName(*s)
	}
	return puo
}

// SetURI sets the "uri" field.
func (puo *PermissionUpdateOne) SetURI(s string) *PermissionUpdateOne {
	puo.mutation.SetURI(s)
	return puo
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableURI(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetURI(*s)
	}
	return puo
}

// SetMethod sets the "method" field.
func (puo *PermissionUpdateOne) SetMethod(s string) *PermissionUpdateOne {
	puo.mutation.SetMethod(s)
	return puo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableMethod(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetMethod(*s)
	}
	return puo
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (puo *PermissionUpdateOne) AddPermissionIDs(ids ...uuid.UUID) *PermissionUpdateOne {
	puo.mutation.AddPermissionIDs(ids...)
	return puo
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (puo *PermissionUpdateOne) AddPermissions(p ...*Permission) *PermissionUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPermissionIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearPermissions clears all "permissions" edges to the Permission entity.
func (puo *PermissionUpdateOne) ClearPermissions() *PermissionUpdateOne {
	puo.mutation.ClearPermissions()
	return puo
}

// RemovePermissionIDs removes the "permissions" edge to Permission entities by IDs.
func (puo *PermissionUpdateOne) RemovePermissionIDs(ids ...uuid.UUID) *PermissionUpdateOne {
	puo.mutation.RemovePermissionIDs(ids...)
	return puo
}

// RemovePermissions removes "permissions" edges to Permission entities.
func (puo *PermissionUpdateOne) RemovePermissions(p ...*Permission) *PermissionUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePermissionIDs(ids...)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PermissionUpdateOne) check() error {
	if v, ok := puo.mutation.PermissionName(); ok {
		if err := permission.PermissionNameValidator(v); err != nil {
			return &ValidationError{Name: "permissionName", err: fmt.Errorf(`ent: validator failed for field "Permission.permissionName": %w`, err)}
		}
	}
	if v, ok := puo.mutation.URI(); ok {
		if err := permission.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Permission.uri": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Method(); ok {
		if err := permission.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "Permission.method": %w`, err)}
		}
	}
	return nil
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.PermissionName(); ok {
		_spec.SetField(permission.FieldPermissionName, field.TypeString, value)
	}
	if value, ok := puo.mutation.URI(); ok {
		_spec.SetField(permission.FieldURI, field.TypeString, value)
	}
	if value, ok := puo.mutation.Method(); ok {
		_spec.SetField(permission.FieldMethod, field.TypeString, value)
	}
	if puo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.PermissionsTable,
			Columns: permission.PermissionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !puo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.PermissionsTable,
			Columns: permission.PermissionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.PermissionsTable,
			Columns: permission.PermissionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
