package main

import (
	"fmt"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/internal/service/scrapper/usecase"
	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Information Retrieval - Comics Recommendation")

	c := colly.NewCollector()

	usecase.ScrapWebToon(c)
}
