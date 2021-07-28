package handler

import (
	"encoding/json"
	"net/http"

	"cat-feed-amount-api/models"
)

type catRequest struct {
	Weight     float64 `json:"weight"`
	Age        int     `json:"age"`
	IsNeutered bool    `json:"is_neutered"`
	Kcal       float64 `json:"kcal"`
	PerGram    int     `json:"per_gram"`
}

type catResponse struct {
	FeedAmount float64 `json:"feed_amount"`
	FeedKcal   float64 `json:"feed_kcal"`
}

func CatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req catRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		feed := models.Feed{
			Kcal:    req.Kcal,
			PerGram: req.PerGram,
		}
		cat := models.Cat{
			Weight:     req.Weight,
			Age:        req.Age,
			IsNeutered: req.IsNeutered,
			Feed:       feed,
		}

		feedAmount, feedKcal := cat.CalcFeedAmount()

		res, err := json.Marshal(catResponse{feedAmount, feedKcal})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}

}
