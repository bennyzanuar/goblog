package utils

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Params  interface{} `json:"params,omitempty" `
	Data    interface{} `json:"data,omitempty" `
}

type ErrorResponse struct {
	Success bool      `json:"success" example:"false"`
	Error   HttpError `json:"error,omitempty"`
}

type HttpError struct {
	Code    uint32 `json:"code" example:"40001"`
	Message string `json:"message" example:"status bad request"`
}

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
func Respond(w http.ResponseWriter, data interface{}, params interface{}) {
	w.Header().Set("Content-Type", "application/json")

	resp := SuccessResponse{
		Success: true,
		Params:  params,
		Data:    data,
	}

	json.NewEncoder(w).Encode(resp)
}
