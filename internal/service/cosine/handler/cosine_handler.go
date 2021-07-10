package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RiskyFeryansyahP/ir-comics-recommendation/internal/model"
	"github.com/gorilla/mux"
)

// CosineHandler ...
type CosineHandler struct {
	CosineUC model.UsecaseCosine
}

// NewCosineHandler ...
func NewCosineHandler(cosineUc model.UsecaseCosine) *CosineHandler {
	return &CosineHandler{
		CosineUC: cosineUc,
	}
}

// Comics ...
func (ch *CosineHandler) Comics(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	w.Header().Set("Content-Type", "application/json")

	res, err := ch.CosineUC.GetComics(ctx)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
	}

	json.NewEncoder(w).Encode(res)
}

// Comic ...
func (ch *CosineHandler) Comic(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	w.Header().Set("Content-Type", "application/json")

	res, err := ch.CosineUC.GetCosine(ctx, id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
	}

	json.NewEncoder(w).Encode(res)
}
