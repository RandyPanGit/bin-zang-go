package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RandyPanGit/bin-zang-go/internal/delivery"
	"github.com/gorilla/mux"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8081"
	}
	mux := mux.NewRouter()
	delivery.BuildRoute(mux)
	fmt.Print(port)
	http.Handle("/", mux)
	http.ListenAndServe(":"+port, nil)
}
