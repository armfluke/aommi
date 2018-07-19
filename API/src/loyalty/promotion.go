package loyalty

import (
	"encoding/json"
	"net/http"
)

type Promotion struct {
	PromotionID   int    `json:"promotionID"`
	PromotionName string `json:"promotionName"`
	Reward        string `json:"reward"`
	Condition     string `json:"condition"`
	Point         int    `json:"point"`
	Image         string `json:"image"`
}

func GetPromotion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	msg := Promotion{PromotionID: 0, PromotionName: "", Reward: "", Condition: "", Point: 0, Image: ""}

	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}
