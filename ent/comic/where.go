// Code generated by entc, DO NOT EDIT.

package comic

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "Title" field. It's identical to TitleEQ.
func Title(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Author applies equality check predicate on the "Author" field. It's identical to AuthorEQ.
func Author(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthor), v))
	})
}

// Description applies equality check predicate on the "Description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// TitleEQ applies the EQ predicate on the "Title" field.
func TitleEQ(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "Title" field.
func TitleNEQ(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "Title" field.
func TitleIn(vs ...string) predicate.Comic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "Title" field.
func TitleNotIn(vs ...string) predicate.Comic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "Title" field.
func TitleGT(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "Title" field.
func TitleGTE(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "Title" field.
func TitleLT(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "Title" field.
func TitleLTE(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "Title" field.
func TitleContains(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "Title" field.
func TitleHasPrefix(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "Title" field.
func TitleHasSuffix(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "Title" field.
func TitleEqualFold(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "Title" field.
func TitleContainsFold(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// AuthorEQ applies the EQ predicate on the "Author" field.
func AuthorEQ(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthor), v))
	})
}

// AuthorNEQ applies the NEQ predicate on the "Author" field.
func AuthorNEQ(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthor), v))
	})
}

// AuthorIn applies the In predicate on the "Author" field.
func AuthorIn(vs ...string) predicate.Comic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAuthor), v...))
	})
}

// AuthorNotIn applies the NotIn predicate on the "Author" field.
func AuthorNotIn(vs ...string) predicate.Comic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAuthor), v...))
	})
}

// AuthorGT applies the GT predicate on the "Author" field.
func AuthorGT(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthor), v))
	})
}

// AuthorGTE applies the GTE predicate on the "Author" field.
func AuthorGTE(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthor), v))
	})
}

// AuthorLT applies the LT predicate on the "Author" field.
func AuthorLT(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthor), v))
	})
}

// AuthorLTE applies the LTE predicate on the "Author" field.
func AuthorLTE(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthor), v))
	})
}

// AuthorContains applies the Contains predicate on the "Author" field.
func AuthorContains(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthor), v))
	})
}

// AuthorHasPrefix applies the HasPrefix predicate on the "Author" field.
func AuthorHasPrefix(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthor), v))
	})
}

// AuthorHasSuffix applies the HasSuffix predicate on the "Author" field.
func AuthorHasSuffix(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthor), v))
	})
}

// AuthorEqualFold applies the EqualFold predicate on the "Author" field.
func AuthorEqualFold(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthor), v))
	})
}

// AuthorContainsFold applies the ContainsFold predicate on the "Author" field.
func AuthorContainsFold(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthor), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "Description" field.
func DescriptionEQ(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "Description" field.
func DescriptionNEQ(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "Description" field.
func DescriptionIn(vs ...string) predicate.Comic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "Description" field.
func DescriptionNotIn(vs ...string) predicate.Comic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "Description" field.
func DescriptionGT(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "Description" field.
func DescriptionGTE(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "Description" field.
func DescriptionLT(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "Description" field.
func DescriptionLTE(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "Description" field.
func DescriptionContains(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "Description" field.
func DescriptionHasPrefix(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "Description" field.
func DescriptionHasSuffix(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "Description" field.
func DescriptionEqualFold(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "Description" field.
func DescriptionContainsFold(v string) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// HasGenres applies the HasEdge predicate on the "genres" edge.
func HasGenres() predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenresTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, GenresTable, GenresPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenresWith applies the HasEdge predicate on the "genres" edge with a given conditions (other predicates).
func HasGenresWith(preds ...predicate.Genre) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenresInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, GenresTable, GenresPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comic) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comic) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comic) predicate.Comic {
	return predicate.Comic(func(s *sql.Selector) {
		p(s.Not())
	})
}
