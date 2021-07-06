package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/internal/model"
	"github.com/gocolly/colly"
)

// ScrapperUsecase ...
type ScrapperUsecase struct {
	ScrapperRepo model.MysqlRepositoryScrap
	Colly        *colly.Collector
}

// NewScrapperUsecase ...
func NewScrapperUsecase(scrapperRepo model.MysqlRepositoryScrap, c *colly.Collector) model.UsecaseScrap {
	return &ScrapperUsecase{
		ScrapperRepo: scrapperRepo,
		Colly:        c,
	}
}

// ScrapWebToon ...
func (su *ScrapperUsecase) ScrapWebToon(ctx context.Context) error {
	genres := make([]*model.Genre, 0)
	comics := make([]*model.Comic, 0)

	su.Colly.OnHTML(".sub_title", func(e *colly.HTMLElement) {
		genre := &model.Genre{
			Name: strings.ToLower(e.Text),
		}

		genres = append(genres, genre)
	})

	su.Colly.OnHTML(".card_lst li a .info", func(h *colly.HTMLElement) {
		comic := &model.Comic{
			Title:  h.ChildText(".subj"),
			Author: h.ChildText(".author"),
			Like:   h.ChildText(".grade_area > .grade_num"),
		}

		comics = append(comics, comic)
	})

	su.Colly.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	su.Colly.Visit("https://www.webtoons.com/id/genre")

	_, err := su.ScrapperRepo.InsertGenre(ctx, genres)
	if err != nil {
		return err
	}

	_, err = su.ScrapperRepo.InsertComic(ctx, comics)
	if err != nil {
		return err
	}

	return nil
}
