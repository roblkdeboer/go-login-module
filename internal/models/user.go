package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TODO: add any necessary methods or functions related to the User model