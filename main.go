package main

import (
	"fmt"
	"github.com/RandyPanGit/bin-zang-go/internal/delivery"
	"github.com/RandyPanGit/bin-zang-go/internal/logger"
	"github.com/gorilla/mux"
	"net/http"
	"os"
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
	logger.Log.Init()
}
