// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"queuetable/ent/predicate"
	"queuetable/ent/queue"
	"queuetable/ent/user"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeQueue = "Queue"
	TypeUser  = "User"
)

// QueueMutation represents an operation that mutates the Queue nodes in the graph.
type QueueMutation struct {
	config
	op            Op
	typ           string
	id            *int
	instruction   *string
	obj_table     *queue.ObjTable
	obj_id        *uuid.UUID
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Queue, error)
	predicates    []predicate.Queue
}

var _ ent.Mutation = (*QueueMutation)(nil)

// queueOption allows management of the mutation configuration using functional options.
type queueOption func(*QueueMutation)

// newQueueMutation creates new mutation for the Queue entity.
func newQueueMutation(c config, op Op, opts ...queueOption) *QueueMutation {
	m := &QueueMutation{
		config:        c,
		op:            op,
		typ:           TypeQueue,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withQueueID sets the ID field of the mutation.
func withQueueID(id int) queueOption {
	return func(m *QueueMutation) {
		var (
			err   error
			once  sync.Once
			value *Queue
		)
		m.oldValue = func(ctx context.Context) (*Queue, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Queue.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withQueue sets the old Queue of the mutation.
func withQueue(node *Queue) queueOption {
	return func(m *QueueMutation) {
		m.oldValue = func(context.Context) (*Queue, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m QueueMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m QueueMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *QueueMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *QueueMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Queue.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetInstruction sets the "instruction" field.
func (m *QueueMutation) SetInstruction(s string) {
	m.instruction = &s
}

// Instruction returns the value of the "instruction" field in the mutation.
func (m *QueueMutation) Instruction() (r string, exists bool) {
	v := m.instruction
	if v == nil {
		return
	}
	return *v, true
}

// OldInstruction returns the old "instruction" field's value of the Queue entity.
// If the Queue object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *QueueMutation) OldInstruction(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldInstruction is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldInstruction requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInstruction: %w", err)
	}
	return oldValue.Instruction, nil
}

// ResetInstruction resets all changes to the "instruction" field.
func (m *QueueMutation) ResetInstruction() {
	m.instruction = nil
}

// SetObjTable sets the "obj_table" field.
func (m *QueueMutation) SetObjTable(qt queue.ObjTable) {
	m.obj_table = &qt
}

// ObjTable returns the value of the "obj_table" field in the mutation.
func (m *QueueMutation) ObjTable() (r queue.ObjTable, exists bool) {
	v := m.obj_table
	if v == nil {
		return
	}
	return *v, true
}

// OldObjTable returns the old "obj_table" field's value of the Queue entity.
// If the Queue object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *QueueMutation) OldObjTable(ctx context.Context) (v queue.ObjTable, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldObjTable is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldObjTable requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldObjTable: %w", err)
	}
	return oldValue.ObjTable, nil
}

// ResetObjTable resets all changes to the "obj_table" field.
func (m *QueueMutation) ResetObjTable() {
	m.obj_table = nil
}

// SetObjID sets the "obj_id" field.
func (m *QueueMutation) SetObjID(u uuid.UUID) {
	m.obj_id = &u
}

// ObjID returns the value of the "obj_id" field in the mutation.
func (m *QueueMutation) ObjID() (r uuid.UUID, exists bool) {
	v := m.obj_id
	if v == nil {
		return
	}
	return *v, true
}

// OldObjID returns the old "obj_id" field's value of the Queue entity.
// If the Queue object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *QueueMutation) OldObjID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldObjID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldObjID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldObjID: %w", err)
	}
	return oldValue.ObjID, nil
}

// ResetObjID resets all changes to the "obj_id" field.
func (m *QueueMutation) ResetObjID() {
	m.obj_id = nil
}

// Where appends a list predicates to the QueueMutation builder.
func (m *QueueMutation) Where(ps ...predicate.Queue) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the QueueMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *QueueMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Queue, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *QueueMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *QueueMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Queue).
func (m *QueueMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *QueueMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.instruction != nil {
		fields = append(fields, queue.FieldInstruction)
	}
	if m.obj_table != nil {
		fields = append(fields, queue.FieldObjTable)
	}
	if m.obj_id != nil {
		fields = append(fields, queue.FieldObjID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *QueueMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case queue.FieldInstruction:
		return m.Instruction()
	case queue.FieldObjTable:
		return m.ObjTable()
	case queue.FieldObjID:
		return m.ObjID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *QueueMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case queue.FieldInstruction:
		return m.OldInstruction(ctx)
	case queue.FieldObjTable:
		return m.OldObjTable(ctx)
	case queue.FieldObjID:
		return m.OldObjID(ctx)
	}
	return nil, fmt.Errorf("unknown Queue field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *QueueMutation) SetField(name string, value ent.Value) error {
	switch name {
	case queue.FieldInstruction:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInstruction(v)
		return nil
	case queue.FieldObjTable:
		v, ok := value.(queue.ObjTable)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetObjTable(v)
		return nil
	case queue.FieldObjID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetObjID(v)
		return nil
	}
	return fmt.Errorf("unknown Queue field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *QueueMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *QueueMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *QueueMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Queue numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *QueueMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *QueueMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *QueueMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Queue nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *QueueMutation) ResetField(name string) error {
	switch name {
	case queue.FieldInstruction:
		m.ResetInstruction()
		return nil
	case queue.FieldObjTable:
		m.ResetObjTable()
		return nil
	case queue.FieldObjID:
		m.ResetObjID()
		return nil
	}
	return fmt.Errorf("unknown Queue field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *QueueMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *QueueMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *QueueMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *QueueMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *QueueMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *QueueMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *QueueMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Queue unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *QueueMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Queue edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	queue_obj_id  *uuid.UUID
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetQueueObjID sets the "queue_obj_id" field.
func (m *UserMutation) SetQueueObjID(u uuid.UUID) {
	m.queue_obj_id = &u
}

// QueueObjID returns the value of the "queue_obj_id" field in the mutation.
func (m *UserMutation) QueueObjID() (r uuid.UUID, exists bool) {
	v := m.queue_obj_id
	if v == nil {
		return
	}
	return *v, true
}

// OldQueueObjID returns the old "queue_obj_id" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldQueueObjID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldQueueObjID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldQueueObjID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldQueueObjID: %w", err)
	}
	return oldValue.QueueObjID, nil
}

// ResetQueueObjID resets all changes to the "queue_obj_id" field.
func (m *UserMutation) ResetQueueObjID() {
	m.queue_obj_id = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.queue_obj_id != nil {
		fields = append(fields, user.FieldQueueObjID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldQueueObjID:
		return m.QueueObjID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldQueueObjID:
		return m.OldQueueObjID(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldQueueObjID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetQueueObjID(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldQueueObjID:
		m.ResetQueueObjID()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
