package api

import (
	"encoding/json"
	"net/http"
)

// Endpoint parms
type PlayerCounterParam struct {
	Player string
}

// Get counter response
type PlayerCounterGetResponse struct {
	Code    int
	Counter int
}

// Error response
type Error struct {
	Code     int
	Meassage string
}

func writeError(w http.ResponseWriter, code int, message string) {

	resp := Error{
		Code:     code,
		Meassage: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusBadRequest, err.Error())
	}
	InternalServerError = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusInternalServerError, "An unexpected error occurred")
	}
)
