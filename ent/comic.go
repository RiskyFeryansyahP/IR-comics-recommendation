// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/comic"
)

// Comic is the model entity for the Comic schema.
type Comic struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "Title" field.
	Title string `json:"Title,omitempty"`
	// Author holds the value of the "Author" field.
	Author string `json:"Author,omitempty"`
	// Like holds the value of the "Like" field.
	Like string `json:"Like,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ComicQuery when eager-loading is set.
	Edges ComicEdges `json:"edges"`
}

// ComicEdges holds the relations/edges for other nodes in the graph.
type ComicEdges struct {
	// Genres holds the value of the genres edge.
	Genres []*Genre `json:"genres,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GenresOrErr returns the Genres value or an error if the edge
// was not loaded in eager-loading.
func (e ComicEdges) GenresOrErr() ([]*Genre, error) {
	if e.loadedTypes[0] {
		return e.Genres, nil
	}
	return nil, &NotLoadedError{edge: "genres"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comic) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case comic.FieldID:
			values[i] = new(sql.NullInt64)
		case comic.FieldTitle, comic.FieldAuthor, comic.FieldLike:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Comic", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comic fields.
func (c *Comic) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comic.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case comic.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case comic.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Author", values[i])
			} else if value.Valid {
				c.Author = value.String
			}
		case comic.FieldLike:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Like", values[i])
			} else if value.Valid {
				c.Like = value.String
			}
		}
	}
	return nil
}

// QueryGenres queries the "genres" edge of the Comic entity.
func (c *Comic) QueryGenres() *GenreQuery {
	return (&ComicClient{config: c.config}).QueryGenres(c)
}

// Update returns a builder for updating this Comic.
// Note that you need to call Comic.Unwrap() before calling this method if this Comic
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comic) Update() *ComicUpdateOne {
	return (&ComicClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Comic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comic) Unwrap() *Comic {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comic is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comic) String() string {
	var builder strings.Builder
	builder.WriteString("Comic(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", Title=")
	builder.WriteString(c.Title)
	builder.WriteString(", Author=")
	builder.WriteString(c.Author)
	builder.WriteString(", Like=")
	builder.WriteString(c.Like)
	builder.WriteByte(')')
	return builder.String()
}

// Comics is a parsable slice of Comic.
type Comics []*Comic

func (c Comics) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
