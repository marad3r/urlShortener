package main

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	jsData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsData)
}
