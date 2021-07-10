package repository

import (
	"context"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/comic"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/internal/model"
)

// CosineRepository ...
type CosineRepository struct {
	Client *ent.Client
}

// NewCosineRepository ...
func NewCosineRepository(client *ent.Client) model.MysqlRepositoryCosine {
	return &CosineRepository{
		Client: client,
	}
}

// GetAllComic ...
func (cr *CosineRepository) GetAllComic(ctx context.Context) ([]*ent.Comic, error) {
	comics, err := cr.Client.Comic.Query().WithGenres().All(ctx)

	return comics, err
}

// GetComicByID ...
func (cr *CosineRepository) GetComicByID(ctx context.Context, id int) (*ent.Comic, error) {
	comic, err := cr.Client.Comic.Query().
		WithGenres().
		Where(comic.IDEQ(id)).
		Only(ctx)

	return comic, err
}

// GetComicByName ...
func (cr *CosineRepository) GetComicByName(ctx context.Context, title string) (*ent.Comic, error) {
	comic, err := cr.Client.Comic.Query().
		WithGenres().
		Where(comic.TitleEQ(title)).
		Only(ctx)

	return comic, err
}
