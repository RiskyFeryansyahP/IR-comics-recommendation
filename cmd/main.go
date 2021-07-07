package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/config"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent/migrate"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/internal/service/cosine"
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

	// fmt.Printf("docs %v", docs)

	// d := map[string]string{
	// 	"d1": "naruto shippuuden menceritakan tentang sudаh duа ѕеtеngаh tahun ѕејаk naruto uzumaki meninggalkan konohagakure desa konoha untuk pelatihan intensif ѕеtеlаh acara yang memicu keinginannya untuk mеnјаdі lеbіh kuat sеkаrаng akatsuki organisasi misterius ninja nakal elit mendekati rencana besar mеrеkа yang dapat mengancam keselamatan ѕеluruh dunia shinobi",
	// 	"d2": "setelah peperangan di naruto shippuuden sekarang giliran menceritakan seorang anak dari naruto uzumaki dimana namanya boruto uzumaki di cerita kali ini boruto di hadapkan oleh musuh akatsuki juga yang merupakan musuh dari ayahnya uzumaki naruto disini boruto akan belajar bagaimana menjadi seorang ninja dan akan belajar langsung ke ahlinya",
	// 	"d3": "it has survived not only five centuries but also the leap into electronic typesetting remaining essentially unchanged it was popularised in the 1960s",
	// }

	docsIdx := cosine.DocsIndex(docs)
	reversedDocs := cosine.ReversedDocs(docsIdx)
	// fmt.Println("revers", reversedDocs)
	// fmt.Println(getTF(reversedDocs, "it", "d3"))
	// fmt.Println(getDF(reversedDocs, "it"))
	// fmt.Println(getIDF(reversedDocs, "it", len(d)))
	// fmt.Println(getTFIDF(reversedDocs, "it", "d3", 3))
	myTerms := cosine.MyTerms(reversedDocs)
	// v1 := vectorize(reversedDocs, "Naruto", myTerms, len(docs))
	// fmt.Println(v1)
	// v2 := vectorize(reversedDocs, "Boruto", myTerms, len(docs))
	// fmt.Println(v2)
	// v3 := vectorize(reversedDocs, "Nanatsu", myTerms, len(docs))
	// fmt.Println(v3)

	// fmt.Println(getDotProduct(v1, v2))
	// fmt.Println(getVectorLength(v1))
	// fmt.Println(getVectorLength(v3))

	// fmt.Println("cosine simolarity")
	// fmt.Println(cosineSimilarity(v1, v2))
	cosine.GetCosineSimilarity(docs, reversedDocs, "Tokyo Ghoul", myTerms)
}
