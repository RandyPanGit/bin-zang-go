package repository

var Users []User

func GetAllUsers() (users []User) {
	if Users == nil {
		Users = make([]User, 100)
	}
	return Users
}

func AddNewUser(user User) (users []User) {
	Users = append(Users, user)
	return Users
}

type User struct {
	ID   int
	Name string
}
