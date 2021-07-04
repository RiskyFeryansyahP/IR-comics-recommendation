package model

import (
	"context"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent"
)

// Scrap ...
type Scrap struct{}

// UsecaseScrap ...
type UsecaseScrap interface{}

// MysqlRepositoryScrap ...
type MysqlRepositoryScrap struct{
  InsertGenre(ctx context.Context, data []*Genre) ([]*ent.Genre, error)
	InsertComic(ctx context.Context, data []*Comic) ([]*ent.Comic, error)
}
