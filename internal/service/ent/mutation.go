// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-rpc-happysooner-file/internal/service/ent/file"
	"sync"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFile = "File"
)

// FileMutation represents an operation that mutate the Files
// nodes in the graph.
type FileMutation struct {
	config
	op            Op
	typ           string
	id            *int
	uid           *int
	adduid        *int
	url           *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*File, error)
}

var _ ent.Mutation = (*FileMutation)(nil)

// fileOption allows to manage the mutation configuration using functional options.
type fileOption func(*FileMutation)

// newFileMutation creates new mutation for $n.Name.
func newFileMutation(c config, op Op, opts ...fileOption) *FileMutation {
	m := &FileMutation{
		config:        c,
		op:            op,
		typ:           TypeFile,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFileID sets the id field of the mutation.
func withFileID(id int) fileOption {
	return func(m *FileMutation) {
		var (
			err   error
			once  sync.Once
			value *File
		)
		m.oldValue = func(ctx context.Context) (*File, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().File.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFile sets the old File of the mutation.
func withFile(node *File) fileOption {
	return func(m *FileMutation) {
		m.oldValue = func(context.Context) (*File, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FileMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FileMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *FileMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUID sets the uid field.
func (m *FileMutation) SetUID(i int) {
	m.uid = &i
	m.adduid = nil
}

// UID returns the uid value in the mutation.
func (m *FileMutation) UID() (r int, exists bool) {
	v := m.uid
	if v == nil {
		return
	}
	return *v, true
}

// OldUID returns the old uid value of the File.
// If the File object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *FileMutation) OldUID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUID: %w", err)
	}
	return oldValue.UID, nil
}

// AddUID adds i to uid.
func (m *FileMutation) AddUID(i int) {
	if m.adduid != nil {
		*m.adduid += i
	} else {
		m.adduid = &i
	}
}

// AddedUID returns the value that was added to the uid field in this mutation.
func (m *FileMutation) AddedUID() (r int, exists bool) {
	v := m.adduid
	if v == nil {
		return
	}
	return *v, true
}

// ResetUID reset all changes of the "uid" field.
func (m *FileMutation) ResetUID() {
	m.uid = nil
	m.adduid = nil
}

// SetURL sets the url field.
func (m *FileMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the url value in the mutation.
func (m *FileMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old url value of the File.
// If the File object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *FileMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldURL is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL reset all changes of the "url" field.
func (m *FileMutation) ResetURL() {
	m.url = nil
}

// Op returns the operation name.
func (m *FileMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (File).
func (m *FileMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *FileMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.uid != nil {
		fields = append(fields, file.FieldUID)
	}
	if m.url != nil {
		fields = append(fields, file.FieldURL)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *FileMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case file.FieldUID:
		return m.UID()
	case file.FieldURL:
		return m.URL()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *FileMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case file.FieldUID:
		return m.OldUID(ctx)
	case file.FieldURL:
		return m.OldURL(ctx)
	}
	return nil, fmt.Errorf("unknown File field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *FileMutation) SetField(name string, value ent.Value) error {
	switch name {
	case file.FieldUID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUID(v)
		return nil
	case file.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	}
	return fmt.Errorf("unknown File field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *FileMutation) AddedFields() []string {
	var fields []string
	if m.adduid != nil {
		fields = append(fields, file.FieldUID)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *FileMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case file.FieldUID:
		return m.AddedUID()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *FileMutation) AddField(name string, value ent.Value) error {
	switch name {
	case file.FieldUID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUID(v)
		return nil
	}
	return fmt.Errorf("unknown File numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *FileMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *FileMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *FileMutation) ClearField(name string) error {
	return fmt.Errorf("unknown File nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *FileMutation) ResetField(name string) error {
	switch name {
	case file.FieldUID:
		m.ResetUID()
		return nil
	case file.FieldURL:
		m.ResetURL()
		return nil
	}
	return fmt.Errorf("unknown File field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *FileMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *FileMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *FileMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *FileMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *FileMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *FileMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *FileMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown File unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *FileMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown File edge %s", name)
}
