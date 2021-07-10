package usecase

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/ent"
	"github.com/RiskyFeryansyahP/ir-comics-recommendation/internal/model"
)

// CosineUsecase ...
type CosineUsecase struct {
	CosineRepo model.MysqlRepositoryCosine
}

// NewCosineUsecase ...
func NewCosineUsecase(repo model.MysqlRepositoryCosine) model.UsecaseCosine {
	return &CosineUsecase{
		CosineRepo: repo,
	}
}

// GetComics ...
func (cu *CosineUsecase) GetComics(ctx context.Context) ([]*ent.Comic, error) {
	comics, err := cu.CosineRepo.GetAllComic(ctx)
	if err != nil {
		return nil, err
	}

	return comics, nil
}

// GetCosine ...
func (cu *CosineUsecase) GetCosine(ctx context.Context, id int) (*model.Cosine, error) {
	docs := make(map[string]*ent.Comic)

	listComics := make([]*ent.Comic, 0)

	comics, err := cu.CosineRepo.GetAllComic(ctx)
	if err != nil {
		return nil, err
	}

	comic, err := cu.CosineRepo.GetComicByID(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, v := range comics {
		docs[v.Title] = v
	}

	docsIdx := DocsIndex(docs)
	reversedDocs := ReversedDocs(docsIdx)
	myTerms := MyTerms(reversedDocs)
	list := GetCosineSimilarity(docs, reversedDocs, comic.Title, myTerms)

	for _, v := range list {
		c, err := cu.CosineRepo.GetComicByName(ctx, v)
		if err != nil {
			return nil, err
		}

		listComics = append(listComics, c)
	}

	res := &model.Cosine{
		Comic:          comic,
		Recommendation: listComics,
	}

	return res, nil
}

// GetCosineSimilarity ...
func GetCosineSimilarity(docs map[string]*ent.Comic, reversed map[string]map[string]int, term string, myTerms []string) []string {
	list := make([]string, 0)

	v1 := vectorize(reversed, term, myTerms, len(docs))
	for k := range docs {
		if k == term {
			continue
		}
		v2 := vectorize(reversed, k, myTerms, len(docs))
		res := cosineSimilarity(v1, v2)

		if res > 0.50 {
			list = append(list, k)
			fmt.Println(k)
		}
	}

	return list
}

// DocsIndex ...
func DocsIndex(docs map[string]*ent.Comic) map[string]map[string]int {
	newDocs := make(map[string]map[string]int)

	for _, val := range docs {
		words := strings.Split(val.Description, " ")

		doc := make(map[string]int)
		for _, v := range words {
			if _, ok := doc[v]; !ok {
				doc[v] = 1
			} else {
				doc[v]++
			}
		}

		newDocs[val.Title] = doc
	}

	return newDocs
}

// DocsIndex2 ...
func DocsIndex2(docs map[string]string) map[string]map[string]int {
	newDocs := make(map[string]map[string]int)

	for key, val := range docs {
		words := strings.Split(val, " ")

		doc := make(map[string]int)
		for _, v := range words {
			if _, ok := doc[v]; !ok {
				doc[v] = 1
			} else {
				doc[v]++
			}
		}

		newDocs[key] = doc
	}

	return newDocs
}

// ReversedDocs ...
func ReversedDocs(documents map[string]map[string]int) map[string]map[string]int {
	newDocs := make(map[string]map[string]int)

	for key, tf := range documents {
		for tKey, frequency := range tf {
			if _, ok := newDocs[tKey]; !ok {
				newDocs[tKey] = make(map[string]int)
			}
			newDocs[tKey][key] = frequency
		}
	}

	return newDocs
}

// GetTF ...
func GetTF(reversedDocs map[string]map[string]int, term, doc string) int {
	if v, ok := reversedDocs[term][doc]; ok {
		return v
	}

	return 0
}

// GetDF ...
func GetDF(reversedDocs map[string]map[string]int, term string) int {
	if v, ok := reversedDocs[term]; ok {
		return len(v)
	}

	return 0
}

// GetIDF ...
func GetIDF(reversedDocs map[string]map[string]int, term string, docCount int) float64 {
	df := GetDF(reversedDocs, term)

	return math.Log(float64(docCount) / float64(df))
}

func getTFIDF(reversedDocs map[string]map[string]int, term, doc string, docCount int) float64 {
	idf := GetIDF(reversedDocs, term, docCount)
	tf := GetTF(reversedDocs, term, doc)

	return float64(tf) * float64(idf)
}

// MyTerms ...
func MyTerms(reversedDocs map[string]map[string]int) []string {
	terms := make([]string, 0)

	for k := range reversedDocs {
		terms = append(terms, k)
	}

	return terms
}

func vectorize(reversedDocs map[string]map[string]int, doc string, terms []string, docCount int) []float64 {
	vector := make([]float64, 0)

	for _, v := range terms {
		tfidf := getTFIDF(reversedDocs, v, doc, docCount)

		vector = append(vector, tfidf)
	}

	return vector
}

func getDotProduct(v1 []float64, v2 []float64) float64 {
	result := 0.0

	for k := range v1 {
		result += v1[k] * v2[k]
	}

	return result
}

func getVectorLength(v []float64) float64 {
	result := 0.0

	for k := range v {
		result = result + math.Pow(v[k], 2)
	}

	result = math.Pow(result, 0.5)

	return result
}

func cosineSimilarity(v1 []float64, v2 []float64) float64 {
	dotProduct := getDotProduct(v1, v2)
	v1Length := getVectorLength(v1)
	v2Length := getVectorLength(v2)

	return dotProduct / (v1Length * v2Length)
}
