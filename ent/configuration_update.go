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
	"github.com/wandile/smtp-provider/ent/configuration"
	"github.com/wandile/smtp-provider/ent/predicate"
	"github.com/wandile/smtp-provider/ent/tenant"
)

// ConfigurationUpdate is the builder for updating Configuration entities.
type ConfigurationUpdate struct {
	config
	hooks    []Hook
	mutation *ConfigurationMutation
}

// Where appends a list predicates to the ConfigurationUpdate builder.
func (cu *ConfigurationUpdate) Where(ps ...predicate.Configuration) *ConfigurationUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetHost sets the "host" field.
func (cu *ConfigurationUpdate) SetHost(s string) *ConfigurationUpdate {
	cu.mutation.SetHost(s)
	return cu
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillableHost(s *string) *ConfigurationUpdate {
	if s != nil {
		cu.SetHost(*s)
	}
	return cu
}

// SetPort sets the "port" field.
func (cu *ConfigurationUpdate) SetPort(i int32) *ConfigurationUpdate {
	cu.mutation.ResetPort()
	cu.mutation.SetPort(i)
	return cu
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillablePort(i *int32) *ConfigurationUpdate {
	if i != nil {
		cu.SetPort(*i)
	}
	return cu
}

// AddPort adds i to the "port" field.
func (cu *ConfigurationUpdate) AddPort(i int32) *ConfigurationUpdate {
	cu.mutation.AddPort(i)
	return cu
}

// SetUsername sets the "username" field.
func (cu *ConfigurationUpdate) SetUsername(s string) *ConfigurationUpdate {
	cu.mutation.SetUsername(s)
	return cu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillableUsername(s *string) *ConfigurationUpdate {
	if s != nil {
		cu.SetUsername(*s)
	}
	return cu
}

// SetPassword sets the "password" field.
func (cu *ConfigurationUpdate) SetPassword(s string) *ConfigurationUpdate {
	cu.mutation.SetPassword(s)
	return cu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillablePassword(s *string) *ConfigurationUpdate {
	if s != nil {
		cu.SetPassword(*s)
	}
	return cu
}

// SetEnable sets the "enable" field.
func (cu *ConfigurationUpdate) SetEnable(b bool) *ConfigurationUpdate {
	cu.mutation.SetEnable(b)
	return cu
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillableEnable(b *bool) *ConfigurationUpdate {
	if b != nil {
		cu.SetEnable(*b)
	}
	return cu
}

// AddTenantIDs adds the "tenant" edge to the Tenant entity by IDs.
func (cu *ConfigurationUpdate) AddTenantIDs(ids ...uuid.UUID) *ConfigurationUpdate {
	cu.mutation.AddTenantIDs(ids...)
	return cu
}

// AddTenant adds the "tenant" edges to the Tenant entity.
func (cu *ConfigurationUpdate) AddTenant(t ...*Tenant) *ConfigurationUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddTenantIDs(ids...)
}

// Mutation returns the ConfigurationMutation object of the builder.
func (cu *ConfigurationUpdate) Mutation() *ConfigurationMutation {
	return cu.mutation
}

// ClearTenant clears all "tenant" edges to the Tenant entity.
func (cu *ConfigurationUpdate) ClearTenant() *ConfigurationUpdate {
	cu.mutation.ClearTenant()
	return cu
}

// RemoveTenantIDs removes the "tenant" edge to Tenant entities by IDs.
func (cu *ConfigurationUpdate) RemoveTenantIDs(ids ...uuid.UUID) *ConfigurationUpdate {
	cu.mutation.RemoveTenantIDs(ids...)
	return cu
}

// RemoveTenant removes "tenant" edges to Tenant entities.
func (cu *ConfigurationUpdate) RemoveTenant(t ...*Tenant) *ConfigurationUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveTenantIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConfigurationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConfigurationUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConfigurationUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConfigurationUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ConfigurationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(configuration.Table, configuration.Columns, sqlgraph.NewFieldSpec(configuration.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Host(); ok {
		_spec.SetField(configuration.FieldHost, field.TypeString, value)
	}
	if value, ok := cu.mutation.Port(); ok {
		_spec.SetField(configuration.FieldPort, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.AddedPort(); ok {
		_spec.AddField(configuration.FieldPort, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.Username(); ok {
		_spec.SetField(configuration.FieldUsername, field.TypeString, value)
	}
	if value, ok := cu.mutation.Password(); ok {
		_spec.SetField(configuration.FieldPassword, field.TypeString, value)
	}
	if value, ok := cu.mutation.Enable(); ok {
		_spec.SetField(configuration.FieldEnable, field.TypeBool, value)
	}
	if cu.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   configuration.TenantTable,
			Columns: configuration.TenantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedTenantIDs(); len(nodes) > 0 && !cu.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   configuration.TenantTable,
			Columns: configuration.TenantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   configuration.TenantTable,
			Columns: configuration.TenantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{configuration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConfigurationUpdateOne is the builder for updating a single Configuration entity.
type ConfigurationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConfigurationMutation
}

// SetHost sets the "host" field.
func (cuo *ConfigurationUpdateOne) SetHost(s string) *ConfigurationUpdateOne {
	cuo.mutation.SetHost(s)
	return cuo
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillableHost(s *string) *ConfigurationUpdateOne {
	if s != nil {
		cuo.SetHost(*s)
	}
	return cuo
}

// SetPort sets the "port" field.
func (cuo *ConfigurationUpdateOne) SetPort(i int32) *ConfigurationUpdateOne {
	cuo.mutation.ResetPort()
	cuo.mutation.SetPort(i)
	return cuo
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillablePort(i *int32) *ConfigurationUpdateOne {
	if i != nil {
		cuo.SetPort(*i)
	}
	return cuo
}

// AddPort adds i to the "port" field.
func (cuo *ConfigurationUpdateOne) AddPort(i int32) *ConfigurationUpdateOne {
	cuo.mutation.AddPort(i)
	return cuo
}

// SetUsername sets the "username" field.
func (cuo *ConfigurationUpdateOne) SetUsername(s string) *ConfigurationUpdateOne {
	cuo.mutation.SetUsername(s)
	return cuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillableUsername(s *string) *ConfigurationUpdateOne {
	if s != nil {
		cuo.SetUsername(*s)
	}
	return cuo
}

// SetPassword sets the "password" field.
func (cuo *ConfigurationUpdateOne) SetPassword(s string) *ConfigurationUpdateOne {
	cuo.mutation.SetPassword(s)
	return cuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillablePassword(s *string) *ConfigurationUpdateOne {
	if s != nil {
		cuo.SetPassword(*s)
	}
	return cuo
}

// SetEnable sets the "enable" field.
func (cuo *ConfigurationUpdateOne) SetEnable(b bool) *ConfigurationUpdateOne {
	cuo.mutation.SetEnable(b)
	return cuo
}

// SetNillableEnable sets the "enable" field if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillableEnable(b *bool) *ConfigurationUpdateOne {
	if b != nil {
		cuo.SetEnable(*b)
	}
	return cuo
}

// AddTenantIDs adds the "tenant" edge to the Tenant entity by IDs.
func (cuo *ConfigurationUpdateOne) AddTenantIDs(ids ...uuid.UUID) *ConfigurationUpdateOne {
	cuo.mutation.AddTenantIDs(ids...)
	return cuo
}

// AddTenant adds the "tenant" edges to the Tenant entity.
func (cuo *ConfigurationUpdateOne) AddTenant(t ...*Tenant) *ConfigurationUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddTenantIDs(ids...)
}

// Mutation returns the ConfigurationMutation object of the builder.
func (cuo *ConfigurationUpdateOne) Mutation() *ConfigurationMutation {
	return cuo.mutation
}

// ClearTenant clears all "tenant" edges to the Tenant entity.
func (cuo *ConfigurationUpdateOne) ClearTenant() *ConfigurationUpdateOne {
	cuo.mutation.ClearTenant()
	return cuo
}

// RemoveTenantIDs removes the "tenant" edge to Tenant entities by IDs.
func (cuo *ConfigurationUpdateOne) RemoveTenantIDs(ids ...uuid.UUID) *ConfigurationUpdateOne {
	cuo.mutation.RemoveTenantIDs(ids...)
	return cuo
}

// RemoveTenant removes "tenant" edges to Tenant entities.
func (cuo *ConfigurationUpdateOne) RemoveTenant(t ...*Tenant) *ConfigurationUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveTenantIDs(ids...)
}

// Where appends a list predicates to the ConfigurationUpdate builder.
func (cuo *ConfigurationUpdateOne) Where(ps ...predicate.Configuration) *ConfigurationUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConfigurationUpdateOne) Select(field string, fields ...string) *ConfigurationUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Configuration entity.
func (cuo *ConfigurationUpdateOne) Save(ctx context.Context) (*Configuration, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConfigurationUpdateOne) SaveX(ctx context.Context) *Configuration {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConfigurationUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConfigurationUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ConfigurationUpdateOne) sqlSave(ctx context.Context) (_node *Configuration, err error) {
	_spec := sqlgraph.NewUpdateSpec(configuration.Table, configuration.Columns, sqlgraph.NewFieldSpec(configuration.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Configuration.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, configuration.FieldID)
		for _, f := range fields {
			if !configuration.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != configuration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Host(); ok {
		_spec.SetField(configuration.FieldHost, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Port(); ok {
		_spec.SetField(configuration.FieldPort, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.AddedPort(); ok {
		_spec.AddField(configuration.FieldPort, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.Username(); ok {
		_spec.SetField(configuration.FieldUsername, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Password(); ok {
		_spec.SetField(configuration.FieldPassword, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Enable(); ok {
		_spec.SetField(configuration.FieldEnable, field.TypeBool, value)
	}
	if cuo.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   configuration.TenantTable,
			Columns: configuration.TenantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedTenantIDs(); len(nodes) > 0 && !cuo.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   configuration.TenantTable,
			Columns: configuration.TenantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   configuration.TenantTable,
			Columns: configuration.TenantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Configuration{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{configuration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}