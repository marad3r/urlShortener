package main

import (
	"net/http"
)

type CreateResponse struct {
	ShortUrl string `json:"short"`
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	response := CreateResponse{
		ShortUrl: r.Host + "/" + r.URL.Query().Get("url"),
	}
	jsonResponse(w, response, http.StatusCreated)
}
