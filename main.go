package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users = []User{
	{ID: 1, Name: "张三"},
	{ID: 2, Name: "李四"},
	{ID: 3, Name: "王五"},
}

type User struct {
	ID   int
	Name string
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/user/info/{id}", getUser).Methods(http.MethodGet)
	// http.HandleFunc("/user", getUser)
	// http.HandleFunc("/user/{id}", handleUsers)
	http.Handle("/", mux)
	http.ListenAndServe(":8082", nil)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(users)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\": \"not found\"}")
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.Atoi(id)
	if method == "GET" {
		var user User
		for i := 0; i < len(users); i++ {
			userId := users[i].ID
			if userId == idInt {
				user = users[i]
			}
		}
		userData, _ := json.Marshal(user)
		w.WriteHeader(http.StatusOK)
		w.Write(userData)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\": \"not found\"}")
	}
}
