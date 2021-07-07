// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/comic"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/genre"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/predicate"
)

// ComicUpdate is the builder for updating Comic entities.
type ComicUpdate struct {
	config
	hooks    []Hook
	mutation *ComicMutation
}

// Where adds a new predicate for the ComicUpdate builder.
func (cu *ComicUpdate) Where(ps ...predicate.Comic) *ComicUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetTitle sets the "Title" field.
func (cu *ComicUpdate) SetTitle(s string) *ComicUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetAuthor sets the "Author" field.
func (cu *ComicUpdate) SetAuthor(s string) *ComicUpdate {
	cu.mutation.SetAuthor(s)
	return cu
}

// SetDescription sets the "Description" field.
func (cu *ComicUpdate) SetDescription(s string) *ComicUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// AddGenreIDs adds the "genres" edge to the Genre entity by IDs.
func (cu *ComicUpdate) AddGenreIDs(ids ...int) *ComicUpdate {
	cu.mutation.AddGenreIDs(ids...)
	return cu
}

// AddGenres adds the "genres" edges to the Genre entity.
func (cu *ComicUpdate) AddGenres(g ...*Genre) *ComicUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cu.AddGenreIDs(ids...)
}

// Mutation returns the ComicMutation object of the builder.
func (cu *ComicUpdate) Mutation() *ComicMutation {
	return cu.mutation
}

// ClearGenres clears all "genres" edges to the Genre entity.
func (cu *ComicUpdate) ClearGenres() *ComicUpdate {
	cu.mutation.ClearGenres()
	return cu
}

// RemoveGenreIDs removes the "genres" edge to Genre entities by IDs.
func (cu *ComicUpdate) RemoveGenreIDs(ids ...int) *ComicUpdate {
	cu.mutation.RemoveGenreIDs(ids...)
	return cu
}

// RemoveGenres removes "genres" edges to Genre entities.
func (cu *ComicUpdate) RemoveGenres(g ...*Genre) *ComicUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cu.RemoveGenreIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ComicUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ComicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ComicUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ComicUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ComicUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ComicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comic.Table,
			Columns: comic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comic.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comic.FieldTitle,
		})
	}
	if value, ok := cu.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comic.FieldAuthor,
		})
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comic.FieldDescription,
		})
	}
	if cu.mutation.GenresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   comic.GenresTable,
			Columns: comic.GenresPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: genre.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedGenresIDs(); len(nodes) > 0 && !cu.mutation.GenresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   comic.GenresTable,
			Columns: comic.GenresPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: genre.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.GenresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   comic.GenresTable,
			Columns: comic.GenresPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: genre.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comic.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ComicUpdateOne is the builder for updating a single Comic entity.
type ComicUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ComicMutation
}

// SetTitle sets the "Title" field.
func (cuo *ComicUpdateOne) SetTitle(s string) *ComicUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetAuthor sets the "Author" field.
func (cuo *ComicUpdateOne) SetAuthor(s string) *ComicUpdateOne {
	cuo.mutation.SetAuthor(s)
	return cuo
}

// SetDescription sets the "Description" field.
func (cuo *ComicUpdateOne) SetDescription(s string) *ComicUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// AddGenreIDs adds the "genres" edge to the Genre entity by IDs.
func (cuo *ComicUpdateOne) AddGenreIDs(ids ...int) *ComicUpdateOne {
	cuo.mutation.AddGenreIDs(ids...)
	return cuo
}

// AddGenres adds the "genres" edges to the Genre entity.
func (cuo *ComicUpdateOne) AddGenres(g ...*Genre) *ComicUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cuo.AddGenreIDs(ids...)
}

// Mutation returns the ComicMutation object of the builder.
func (cuo *ComicUpdateOne) Mutation() *ComicMutation {
	return cuo.mutation
}

// ClearGenres clears all "genres" edges to the Genre entity.
func (cuo *ComicUpdateOne) ClearGenres() *ComicUpdateOne {
	cuo.mutation.ClearGenres()
	return cuo
}

// RemoveGenreIDs removes the "genres" edge to Genre entities by IDs.
func (cuo *ComicUpdateOne) RemoveGenreIDs(ids ...int) *ComicUpdateOne {
	cuo.mutation.RemoveGenreIDs(ids...)
	return cuo
}

// RemoveGenres removes "genres" edges to Genre entities.
func (cuo *ComicUpdateOne) RemoveGenres(g ...*Genre) *ComicUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return cuo.RemoveGenreIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ComicUpdateOne) Select(field string, fields ...string) *ComicUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comic entity.
func (cuo *ComicUpdateOne) Save(ctx context.Context) (*Comic, error) {
	var (
		err  error
		node *Comic
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ComicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ComicUpdateOne) SaveX(ctx context.Context) *Comic {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ComicUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ComicUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ComicUpdateOne) sqlSave(ctx context.Context) (_node *Comic, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comic.Table,
			Columns: comic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comic.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Comic.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comic.FieldID)
		for _, f := range fields {
			if !comic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comic.FieldID {
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
	if value, ok := cuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comic.FieldTitle,
		})
	}
	if value, ok := cuo.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comic.FieldAuthor,
		})
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comic.FieldDescription,
		})
	}
	if cuo.mutation.GenresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   comic.GenresTable,
			Columns: comic.GenresPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: genre.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedGenresIDs(); len(nodes) > 0 && !cuo.mutation.GenresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   comic.GenresTable,
			Columns: comic.GenresPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: genre.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.GenresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   comic.GenresTable,
			Columns: comic.GenresPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: genre.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Comic{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comic.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
