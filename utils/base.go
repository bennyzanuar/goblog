package utils

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Message - generate a response to be sent with json
func Message(status bool, message string, data interface{}) *Result {
	var r Result
	r.Status = status
	r.Message = message
	r.Data = data

	return &r
}

// Respond - send a JSON response
func Respond(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
