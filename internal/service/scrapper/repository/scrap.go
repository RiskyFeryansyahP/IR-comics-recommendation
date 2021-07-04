package repository

import (
	"context"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/internal/model"
)

// ScrapperRepository ...
type ScrapperRepository struct {
	Client *ent.Client
}

// NewScrapperRepository ...
func NewScrapperRepository(client *ent.Client) model.MysqlRepositoryScrap {
	return &ScrapperRepository{
		Client: client,
	}
}

// InsertGenre ...
func (sr *ScrapperRepository) InsertGenre(ctx context.Context, data []*model.Genre) ([]*ent.Genre, error) {
	bulkGenre := make([]*ent.GenreCreate, len(data))

	for k, v := range data {
		bulkGenre[k] = sr.Client.Genre.Create().SetName(v.Name)
	}

	genres, err := sr.Client.Genre.CreateBulk(bulkGenre...).Save(ctx)
	return genres, err
}

// InsertComic ...
func (sr *ScrapperRepository) InsertComic(ctx context.Context, data []*model.Comic) ([]*ent.Comic, error) {
	bulkComic := make([]*ent.ComicCreate, len(data))

	for k, v := range data {
		bulkComic[k] = sr.Client.Comic.Create().SetTitle(v.Title).SetAuthor(v.Author).SetLike(v.Like)
	}

	comics, err := sr.Client.Comic.CreateBulk(bulkComic...).Save(ctx)
	return comics, err
}
