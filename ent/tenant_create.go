// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wandile/smtp-provider/ent/configuration"
	"github.com/wandile/smtp-provider/ent/tenant"
)

// TenantCreate is the builder for creating a Tenant entity.
type TenantCreate struct {
	config
	mutation *TenantMutation
	hooks    []Hook
}

// SetTenantName sets the "tenant_name" field.
func (tc *TenantCreate) SetTenantName(s string) *TenantCreate {
	tc.mutation.SetTenantName(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TenantCreate) SetID(u uuid.UUID) *TenantCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TenantCreate) SetNillableID(u *uuid.UUID) *TenantCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// AddConfigurationIDs adds the "configurations" edge to the Configuration entity by IDs.
func (tc *TenantCreate) AddConfigurationIDs(ids ...uuid.UUID) *TenantCreate {
	tc.mutation.AddConfigurationIDs(ids...)
	return tc
}

// AddConfigurations adds the "configurations" edges to the Configuration entity.
func (tc *TenantCreate) AddConfigurations(c ...*Configuration) *TenantCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tc.AddConfigurationIDs(ids...)
}

// Mutation returns the TenantMutation object of the builder.
func (tc *TenantCreate) Mutation() *TenantMutation {
	return tc.mutation
}

// Save creates the Tenant in the database.
func (tc *TenantCreate) Save(ctx context.Context) (*Tenant, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TenantCreate) SaveX(ctx context.Context) *Tenant {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TenantCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TenantCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TenantCreate) defaults() {
	if _, ok := tc.mutation.ID(); !ok {
		v := tenant.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TenantCreate) check() error {
	if _, ok := tc.mutation.TenantName(); !ok {
		return &ValidationError{Name: "tenant_name", err: errors.New(`ent: missing required field "Tenant.tenant_name"`)}
	}
	return nil
}

func (tc *TenantCreate) sqlSave(ctx context.Context) (*Tenant, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TenantCreate) createSpec() (*Tenant, *sqlgraph.CreateSpec) {
	var (
		_node = &Tenant{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tenant.Table, sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeUUID))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.TenantName(); ok {
		_spec.SetField(tenant.FieldTenantName, field.TypeString, value)
		_node.TenantName = value
	}
	if nodes := tc.mutation.ConfigurationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tenant.ConfigurationsTable,
			Columns: tenant.ConfigurationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(configuration.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TenantCreateBulk is the builder for creating many Tenant entities in bulk.
type TenantCreateBulk struct {
	config
	err      error
	builders []*TenantCreate
}

// Save creates the Tenant entities in the database.
func (tcb *TenantCreateBulk) Save(ctx context.Context) ([]*Tenant, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tenant, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TenantMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TenantCreateBulk) SaveX(ctx context.Context) []*Tenant {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TenantCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TenantCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
