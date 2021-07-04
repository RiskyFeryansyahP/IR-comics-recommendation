// Code generated by entc, DO NOT EDIT.

package comic

const (
	// Label holds the string label denoting the comic type in the database.
	Label = "comic"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldLike holds the string denoting the like field in the database.
	FieldLike = "like"
	// EdgeGenres holds the string denoting the genres edge name in mutations.
	EdgeGenres = "genres"
	// Table holds the table name of the comic in the database.
	Table = "comics"
	// GenresTable is the table the holds the genres relation/edge. The primary key declared below.
	GenresTable = "comic_genres"
	// GenresInverseTable is the table name for the Genre entity.
	// It exists in this package in order to avoid circular dependency with the "genre" package.
	GenresInverseTable = "genres"
)

// Columns holds all SQL columns for comic fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldAuthor,
	FieldLike,
}

var (
	// GenresPrimaryKey and GenresColumn2 are the table columns denoting the
	// primary key for the genres relation (M2M).
	GenresPrimaryKey = []string{"comic_id", "genre_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}