package model

import (
	"context"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent"
)

// Cosine ...
type Cosine struct {
	Comic          *ent.Comic
	Recommendation []*ent.Comic
}

// UsecaseCosine ...
type UsecaseCosine interface {
	GetComics(ctx context.Context) ([]*ent.Comic, error)
	GetCosine(ctx context.Context, id int) (*Cosine, error)
}

// MysqlRepositoryCosine ...
type MysqlRepositoryCosine interface {
	GetAllComic(ctx context.Context) ([]*ent.Comic, error)
	GetComicByID(ctx context.Context, id int) (*ent.Comic, error)
	GetComicByName(ctx context.Context, title string) (*ent.Comic, error)
}
