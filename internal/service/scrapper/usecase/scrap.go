package usecase

import (
	"fmt"
	"strings"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/internal/model"
	"github.com/gocolly/colly"
)

// ScrapWebToon ...
func ScrapWebToon(c *colly.Collector) {
	genres := make([]*model.Genre, 0)
	comics := make([]*model.Comic, 0)

	c.OnHTML(".sub_title", func(e *colly.HTMLElement) {
		genre := &model.Genre{
			Name: strings.ToLower(e.Text),
		}

		genres = append(genres, genre)
	})

	c.OnHTML(".card_lst li a .info", func(h *colly.HTMLElement) {
		comic := &model.Comic{
			Title:  h.ChildText(".subj"),
			Author: h.ChildText(".author"),
			Like:   h.ChildText(".grade_area > .grade_num"),
		}

		comics = append(comics, comic)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.webtoons.com/id/genre")
}
