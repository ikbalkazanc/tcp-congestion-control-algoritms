package api

import (
	"encoding/json"
	"net/http"
)

type GenerateResponse struct {
	Length                   int     `json:"length"`
	Rate                     float64 `json:"staratetus"`
	Increment_decrement_rate float64 `json:"increment_decrement_rate"`
	Lost_packages            *[]int  `json:"lost_packages"`
}

// RespondWithJSON write json
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
