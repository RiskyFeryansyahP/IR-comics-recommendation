package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/config"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/migrate"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Information Retrieval - Comics Recommendation")

	ctx := context.Background()

	cfg := config.NewMapConfig()
	client, err := cfg.DBConnection()
	if err != nil {
		log.Printf("no connection database: %+v \n", err)
	}

	err = client.Schema.Create(ctx, migrate.WithDropIndex(true))
	if err != nil {
		log.Printf("failed migration: %+v \n", err)
	}

	// c := colly.NewCollector()
}
