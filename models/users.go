package models

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	email     string `json:"email"`
	password  string `json:"password"`
}

func Create(user User) (User, error) {
	return user, nil
}
