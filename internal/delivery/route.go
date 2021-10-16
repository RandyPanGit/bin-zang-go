package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RandyPanGit/bin-zang-go/internal/repository"
	"github.com/RandyPanGit/bin-zang-go/internal/usecase"
	"github.com/gorilla/mux"
)

func BuildRoute(mux *mux.Router) {
	mux.HandleFunc("/user/info/{id}", getUser).Methods(http.MethodGet)
	mux.HandleFunc("/user", handleUsers).Methods(http.MethodPost, http.MethodPut)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user repository.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userData, _ := usecase.AddNewUser(user)
		w.WriteHeader(http.StatusOK)
		w.Write(userData)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\": \"not found\"}")
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	vars := mux.Vars(r)
	id := vars["id"]
	if method == "GET" {
		userData, _ := usecase.GetUser(id)
		w.WriteHeader(http.StatusOK)
		w.Write(userData)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\": \"not found\"}")
	}
}
