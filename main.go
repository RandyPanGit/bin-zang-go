package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", handleUsers)
	http.ListenAndServe(":8080", nil)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ID:1,Name:Randy")
	fmt.Fprintln(w, "ID:2,Name:David")
	fmt.Fprintln(w, "ID:3,Name:shawn")
}
