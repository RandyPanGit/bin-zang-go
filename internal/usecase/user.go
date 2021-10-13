package usecase

import (
	"encoding/json"
	"strconv"

	"github.com/RandyPanGit/bin-zang-go/internal/repository"
)

func GetUser(userId string) ([]byte, error) {
	var users = repository.GetAllUsers()

	var user repository.User
	idInt, _ := strconv.Atoi(userId)
	for i := 0; i < len(users); i++ {
		userId := users[i].ID
		if userId == idInt {
			user = users[i]
		}
	}
	userData, error := json.Marshal(user)
	return userData, error
}

func AddNewUser(user repository.User) ([]byte, error) {
	users := repository.AddNewUser(user)
	userData, error := json.Marshal(users)
	return userData, error
}
