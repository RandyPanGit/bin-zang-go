package main

import (
	"net/http"

	"github.com/RandyPanGit/bin-zang-go/internal/delivery"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	delivery.BuildRoute(mux)
	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}
