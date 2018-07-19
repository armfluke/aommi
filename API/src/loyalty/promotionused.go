package promotionused

import (
	"encoding/json"
	"net/http"
)

type Promotion struct {
	Prmt string `json:"promotionID"`
	Rewc string `json:"rewardcode"`
	Name string `json:"accountID"`
}

type Status struct {
	Stsc string `json:"statusCode"`
	Sts  string `json:"status"`
}

func Promotion_Is_Use_Post(w http.ResponseWriter, r *http.Request) {
	var u Status
	body := json.NewDecoder(r.Body)
	body.Decode(&u)

	msg := Status{Stsc: "200", Sts: "ok"}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}
