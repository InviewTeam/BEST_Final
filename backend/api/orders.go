package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	db "github.com/InviewTeam/BEST_Final/database"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	h := r.Header.Get("Authorization")
	if _, err := db.Auth(h); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ord := []db.Order{}
	rBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(rBytes, &ord)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.MakeOrder(ord)
	w.WriteHeader(http.StatusOK)
}
