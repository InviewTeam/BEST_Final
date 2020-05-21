package database

import "github.com/jinzhu/gorm"

type Equipment struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type Order struct {
	gorm.Model
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Token struct {
	Token string `json:"token"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserData struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Role     string `json:"role"`
}
