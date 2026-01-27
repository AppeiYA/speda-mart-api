package shared

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors any `json:"errors,omitempty"`
	Token any `json:"token,omitempty"`
}

func ReqResponse(w http.ResponseWriter, statuscode int, payload Payload){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	resp := map[string]any {
		"message": payload.Message,
	}

	resp["success"] = statuscode >= 200 && statuscode < 300

	if payload.Data != nil {
		resp["data"] = payload.Data
	}
	if payload.Errors != nil {
		resp["errors"] = payload.Errors
	}
	if payload.Token != nil {
		resp["token"] = payload.Token
	}

	json.NewEncoder(w).Encode(resp)
}