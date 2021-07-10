package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/config"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/migrate"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/internal/service/cosine/handler"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/internal/service/cosine/repository"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/internal/service/cosine/usecase"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

var docs = make(map[string]*ent.Comic)

func main() {
	fmt.Println("Information Retrieval - Comics Recommendation")

	ctx := context.Background()

	cfg := config.NewMapConfig()
	client, err := cfg.DBConnection()
	if err != nil {
		log.Printf("no connection database: %+v \n", err)
	}

	err = client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true))
	if err != nil {
		log.Printf("failed migration: %+v \n", err)
	}

	c, _ := client.Comic.Query().All(ctx)

	for _, v := range c {
		docs[v.Title] = v
	}

	docsIdx := usecase.DocsIndex(docs)
	reversedDocs := usecase.ReversedDocs(docsIdx)
	myTerms := usecase.MyTerms(reversedDocs)
	usecase.GetCosineSimilarity(docs, reversedDocs, "Naruto", myTerms)

	repo := repository.NewCosineRepository(client)
	uc := usecase.NewCosineUsecase(repo)
	handler := handler.NewCosineHandler(uc)

	r := mux.NewRouter()

	r.HandleFunc("/comics", handler.Comics).Methods(http.MethodGet)
	r.HandleFunc("/comic/{id}", handler.Comic).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
