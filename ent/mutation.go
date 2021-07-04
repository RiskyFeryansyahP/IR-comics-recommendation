// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/comic"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/genre"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeComic = "Comic"
	TypeGenre = "Genre"
)

// ComicMutation represents an operation that mutates the Comic nodes in the graph.
type ComicMutation struct {
	config
	op            Op
	typ           string
	id            *int
	_Title        *string
	_Author       *string
	_Like         *string
	clearedFields map[string]struct{}
	genres        map[int]struct{}
	removedgenres map[int]struct{}
	clearedgenres bool
	done          bool
	oldValue      func(context.Context) (*Comic, error)
	predicates    []predicate.Comic
}

var _ ent.Mutation = (*ComicMutation)(nil)

// comicOption allows management of the mutation configuration using functional options.
type comicOption func(*ComicMutation)

// newComicMutation creates new mutation for the Comic entity.
func newComicMutation(c config, op Op, opts ...comicOption) *ComicMutation {
	m := &ComicMutation{
		config:        c,
		op:            op,
		typ:           TypeComic,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withComicID sets the ID field of the mutation.
func withComicID(id int) comicOption {
	return func(m *ComicMutation) {
		var (
			err   error
			once  sync.Once
			value *Comic
		)
		m.oldValue = func(ctx context.Context) (*Comic, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Comic.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withComic sets the old Comic of the mutation.
func withComic(node *Comic) comicOption {
	return func(m *ComicMutation) {
		m.oldValue = func(context.Context) (*Comic, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ComicMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ComicMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *ComicMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetTitle sets the "Title" field.
func (m *ComicMutation) SetTitle(s string) {
	m._Title = &s
}

// Title returns the value of the "Title" field in the mutation.
func (m *ComicMutation) Title() (r string, exists bool) {
	v := m._Title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "Title" field's value of the Comic entity.
// If the Comic object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ComicMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "Title" field.
func (m *ComicMutation) ResetTitle() {
	m._Title = nil
}

// SetAuthor sets the "Author" field.
func (m *ComicMutation) SetAuthor(s string) {
	m._Author = &s
}

// Author returns the value of the "Author" field in the mutation.
func (m *ComicMutation) Author() (r string, exists bool) {
	v := m._Author
	if v == nil {
		return
	}
	return *v, true
}

// OldAuthor returns the old "Author" field's value of the Comic entity.
// If the Comic object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ComicMutation) OldAuthor(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAuthor is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAuthor requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAuthor: %w", err)
	}
	return oldValue.Author, nil
}

// ResetAuthor resets all changes to the "Author" field.
func (m *ComicMutation) ResetAuthor() {
	m._Author = nil
}

// SetLike sets the "Like" field.
func (m *ComicMutation) SetLike(s string) {
	m._Like = &s
}

// Like returns the value of the "Like" field in the mutation.
func (m *ComicMutation) Like() (r string, exists bool) {
	v := m._Like
	if v == nil {
		return
	}
	return *v, true
}

// OldLike returns the old "Like" field's value of the Comic entity.
// If the Comic object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ComicMutation) OldLike(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLike is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLike requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLike: %w", err)
	}
	return oldValue.Like, nil
}

// ResetLike resets all changes to the "Like" field.
func (m *ComicMutation) ResetLike() {
	m._Like = nil
}

// AddGenreIDs adds the "genres" edge to the Genre entity by ids.
func (m *ComicMutation) AddGenreIDs(ids ...int) {
	if m.genres == nil {
		m.genres = make(map[int]struct{})
	}
	for i := range ids {
		m.genres[ids[i]] = struct{}{}
	}
}

// ClearGenres clears the "genres" edge to the Genre entity.
func (m *ComicMutation) ClearGenres() {
	m.clearedgenres = true
}

// GenresCleared reports if the "genres" edge to the Genre entity was cleared.
func (m *ComicMutation) GenresCleared() bool {
	return m.clearedgenres
}

// RemoveGenreIDs removes the "genres" edge to the Genre entity by IDs.
func (m *ComicMutation) RemoveGenreIDs(ids ...int) {
	if m.removedgenres == nil {
		m.removedgenres = make(map[int]struct{})
	}
	for i := range ids {
		m.removedgenres[ids[i]] = struct{}{}
	}
}

// RemovedGenres returns the removed IDs of the "genres" edge to the Genre entity.
func (m *ComicMutation) RemovedGenresIDs() (ids []int) {
	for id := range m.removedgenres {
		ids = append(ids, id)
	}
	return
}

// GenresIDs returns the "genres" edge IDs in the mutation.
func (m *ComicMutation) GenresIDs() (ids []int) {
	for id := range m.genres {
		ids = append(ids, id)
	}
	return
}

// ResetGenres resets all changes to the "genres" edge.
func (m *ComicMutation) ResetGenres() {
	m.genres = nil
	m.clearedgenres = false
	m.removedgenres = nil
}

// Op returns the operation name.
func (m *ComicMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Comic).
func (m *ComicMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ComicMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m._Title != nil {
		fields = append(fields, comic.FieldTitle)
	}
	if m._Author != nil {
		fields = append(fields, comic.FieldAuthor)
	}
	if m._Like != nil {
		fields = append(fields, comic.FieldLike)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ComicMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case comic.FieldTitle:
		return m.Title()
	case comic.FieldAuthor:
		return m.Author()
	case comic.FieldLike:
		return m.Like()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ComicMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case comic.FieldTitle:
		return m.OldTitle(ctx)
	case comic.FieldAuthor:
		return m.OldAuthor(ctx)
	case comic.FieldLike:
		return m.OldLike(ctx)
	}
	return nil, fmt.Errorf("unknown Comic field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ComicMutation) SetField(name string, value ent.Value) error {
	switch name {
	case comic.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case comic.FieldAuthor:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAuthor(v)
		return nil
	case comic.FieldLike:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLike(v)
		return nil
	}
	return fmt.Errorf("unknown Comic field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ComicMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ComicMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ComicMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Comic numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ComicMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ComicMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ComicMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Comic nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ComicMutation) ResetField(name string) error {
	switch name {
	case comic.FieldTitle:
		m.ResetTitle()
		return nil
	case comic.FieldAuthor:
		m.ResetAuthor()
		return nil
	case comic.FieldLike:
		m.ResetLike()
		return nil
	}
	return fmt.Errorf("unknown Comic field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ComicMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.genres != nil {
		edges = append(edges, comic.EdgeGenres)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ComicMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case comic.EdgeGenres:
		ids := make([]ent.Value, 0, len(m.genres))
		for id := range m.genres {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ComicMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedgenres != nil {
		edges = append(edges, comic.EdgeGenres)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ComicMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case comic.EdgeGenres:
		ids := make([]ent.Value, 0, len(m.removedgenres))
		for id := range m.removedgenres {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ComicMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedgenres {
		edges = append(edges, comic.EdgeGenres)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ComicMutation) EdgeCleared(name string) bool {
	switch name {
	case comic.EdgeGenres:
		return m.clearedgenres
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ComicMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Comic unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ComicMutation) ResetEdge(name string) error {
	switch name {
	case comic.EdgeGenres:
		m.ResetGenres()
		return nil
	}
	return fmt.Errorf("unknown Comic edge %s", name)
}

// GenreMutation represents an operation that mutates the Genre nodes in the graph.
type GenreMutation struct {
	config
	op            Op
	typ           string
	id            *int
	_Name         *string
	clearedFields map[string]struct{}
	comics        map[int]struct{}
	removedcomics map[int]struct{}
	clearedcomics bool
	done          bool
	oldValue      func(context.Context) (*Genre, error)
	predicates    []predicate.Genre
}

var _ ent.Mutation = (*GenreMutation)(nil)

// genreOption allows management of the mutation configuration using functional options.
type genreOption func(*GenreMutation)

// newGenreMutation creates new mutation for the Genre entity.
func newGenreMutation(c config, op Op, opts ...genreOption) *GenreMutation {
	m := &GenreMutation{
		config:        c,
		op:            op,
		typ:           TypeGenre,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGenreID sets the ID field of the mutation.
func withGenreID(id int) genreOption {
	return func(m *GenreMutation) {
		var (
			err   error
			once  sync.Once
			value *Genre
		)
		m.oldValue = func(ctx context.Context) (*Genre, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Genre.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGenre sets the old Genre of the mutation.
func withGenre(node *Genre) genreOption {
	return func(m *GenreMutation) {
		m.oldValue = func(context.Context) (*Genre, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GenreMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GenreMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *GenreMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "Name" field.
func (m *GenreMutation) SetName(s string) {
	m._Name = &s
}

// Name returns the value of the "Name" field in the mutation.
func (m *GenreMutation) Name() (r string, exists bool) {
	v := m._Name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "Name" field's value of the Genre entity.
// If the Genre object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GenreMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "Name" field.
func (m *GenreMutation) ResetName() {
	m._Name = nil
}

// AddComicIDs adds the "comics" edge to the Comic entity by ids.
func (m *GenreMutation) AddComicIDs(ids ...int) {
	if m.comics == nil {
		m.comics = make(map[int]struct{})
	}
	for i := range ids {
		m.comics[ids[i]] = struct{}{}
	}
}

// ClearComics clears the "comics" edge to the Comic entity.
func (m *GenreMutation) ClearComics() {
	m.clearedcomics = true
}

// ComicsCleared reports if the "comics" edge to the Comic entity was cleared.
func (m *GenreMutation) ComicsCleared() bool {
	return m.clearedcomics
}

// RemoveComicIDs removes the "comics" edge to the Comic entity by IDs.
func (m *GenreMutation) RemoveComicIDs(ids ...int) {
	if m.removedcomics == nil {
		m.removedcomics = make(map[int]struct{})
	}
	for i := range ids {
		m.removedcomics[ids[i]] = struct{}{}
	}
}

// RemovedComics returns the removed IDs of the "comics" edge to the Comic entity.
func (m *GenreMutation) RemovedComicsIDs() (ids []int) {
	for id := range m.removedcomics {
		ids = append(ids, id)
	}
	return
}

// ComicsIDs returns the "comics" edge IDs in the mutation.
func (m *GenreMutation) ComicsIDs() (ids []int) {
	for id := range m.comics {
		ids = append(ids, id)
	}
	return
}

// ResetComics resets all changes to the "comics" edge.
func (m *GenreMutation) ResetComics() {
	m.comics = nil
	m.clearedcomics = false
	m.removedcomics = nil
}

// Op returns the operation name.
func (m *GenreMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Genre).
func (m *GenreMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *GenreMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m._Name != nil {
		fields = append(fields, genre.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *GenreMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case genre.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *GenreMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case genre.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Genre field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GenreMutation) SetField(name string, value ent.Value) error {
	switch name {
	case genre.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Genre field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *GenreMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *GenreMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GenreMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Genre numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *GenreMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *GenreMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *GenreMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Genre nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *GenreMutation) ResetField(name string) error {
	switch name {
	case genre.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Genre field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *GenreMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.comics != nil {
		edges = append(edges, genre.EdgeComics)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *GenreMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case genre.EdgeComics:
		ids := make([]ent.Value, 0, len(m.comics))
		for id := range m.comics {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *GenreMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedcomics != nil {
		edges = append(edges, genre.EdgeComics)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *GenreMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case genre.EdgeComics:
		ids := make([]ent.Value, 0, len(m.removedcomics))
		for id := range m.removedcomics {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *GenreMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcomics {
		edges = append(edges, genre.EdgeComics)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *GenreMutation) EdgeCleared(name string) bool {
	switch name {
	case genre.EdgeComics:
		return m.clearedcomics
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *GenreMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Genre unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *GenreMutation) ResetEdge(name string) error {
	switch name {
	case genre.EdgeComics:
		m.ResetComics()
		return nil
	}
	return fmt.Errorf("unknown Genre edge %s", name)
}