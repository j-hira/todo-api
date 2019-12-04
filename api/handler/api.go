package handler

import (
	"api/handler/delete"
	"api/handler/get"
	"api/handler/post"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case http.MethodGet:
		get.Get(w, r)
	case http.MethodPost:
		post.Post(w, r)
	case http.MethodDelete:
		delete.Delete(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
