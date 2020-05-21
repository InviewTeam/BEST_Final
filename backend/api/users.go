package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	db "github.com/InviewTeam/BEST_Final/database"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	creds := db.Credentials{}
	rBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(rBytes, &creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := db.LoginUser(creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		tokenString := db.Token{
			Token: res,
		}
		req, err := json.Marshal(tokenString)
		if err != nil {
			log.Printf("Failed to create token response: %s\n", res)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(req)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	creds := db.Credentials{}
	rBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(rBytes, &creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.CreateUser(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func Whoami(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	h := r.Header.Get("Authorization")
	tk, err := db.Auth(h)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	usr := db.UserData{}
	usr, err = db.GetMyself(tk)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req, err := json.Marshal(usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(req)
}
