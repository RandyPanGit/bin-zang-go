package main

import (
	"net/http"
	"os"

	"github.com/RandyPanGit/bin-zang-go/internal/delivery"
	"github.com/gorilla/mux"
)

func main() {
	var port string
	mux := mux.NewRouter()
	delivery.BuildRoute(mux)
	port, _ = os.LookupEnv("PORT")
	http.Handle("/", mux)
	http.ListenAndServe(port, nil)
}
