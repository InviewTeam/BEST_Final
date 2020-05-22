package database

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var (
	bearerPrefix    = "Bearer "
	ErrUnauthorized = errors.New("Unauthorized")
)

func Auth(h string) (string, error) {
	n := len(bearerPrefix)
	if len(h) <= n || h[:n] != bearerPrefix {
		return "", ErrUnauthorized
	}
	tk := strings.TrimRight(h[n:], " \t\n")
	temp := &User{}
	err := GetDB().Table("users").Where("token = ?", tk).First(temp).Error
	if err != nil {
		return "", ErrUnauthorized
	}
	return tk, nil
}

func (c Credentials) Validate() bool {
	temp := &User{}

	err := GetDB().Table("users").Where("username = ?", c.Username).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}
	if temp.Username != "" {
		return false
	}
	return true
}

func LoginUser(c Credentials) (string, error) {
	user := &User{}
	err := GetDB().Table("users").Where("username = ?", c.Username).First(user).Error
	if err != nil {
		return "", ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.Password))
	if err != nil {
		return "", ErrUnauthorized
	}

	return user.Token, nil
}

func CreateUser(c Credentials) error {
	if !c.Validate() {
		return fmt.Errorf("Username already exists")
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	c.Password = string(hashedPassword)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = c.Username
	claims["ts"] = time.Now().Unix()

	t, err := token.SignedString([]byte(c.Username))
	if err != nil {
		return err
	}

	GetDB().Create(&User{Username: c.Username, Password: c.Password, Token: t, Role: "user"})

	return nil
}

func IsAdmin(token string) bool {
	temp := &User{}
	GetDB().Table("users").Where("token = ?", token).First(temp)
	if temp.Role == "admin" {
		return true
	}
	return false
}

func GetAllUsers(token string) ([]UserData, error) {
	if !IsAdmin(token) {
		return nil, ErrUnauthorized
	}
	res := []User{}
	GetDB().Find(&res)

	resConverted := make([]UserData, len(res))
	for i, u := range res {
		resConverted[i].Username = u.Username
		resConverted[i].Role = u.Role
	}
	return resConverted, nil
}

func GetMyself(token string) (UserData, error) {
	temp := &User{}
	res := UserData{}
	err := GetDB().Table("users").Where("token = ?", token).First(temp).Error
	if err != nil {
		return res, err
	}

	res.Username = temp.Username
	res.Role = temp.Role
	return res, nil
}
