package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	db "github.com/InviewTeam/BEST_Final/database"
)

func getParam(r *http.Request) string {
	p := strings.Split(r.URL.Path, "/")
	return p[len(p)-1]
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	h := r.Header.Get("Authorization")
	if _, err := db.Auth(h); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	eq := db.Equipment{}
	rBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(rBytes, &eq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.Create(eq)
	w.WriteHeader(http.StatusOK)
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	h := r.Header.Get("Authorization")
	if _, err := db.Auth(h); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var eqs []db.Equipment
	eqs = db.Get()
	req, err := json.Marshal(eqs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(req)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	h := r.Header.Get("Authorization")
	if _, err := db.Auth(h); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	eq := db.Equipment{}
	rBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(rBytes, &eq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	param := getParam(r)
	err = db.Update(eq, param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	h := r.Header.Get("Authorization")
	if _, err := db.Auth(h); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	param := getParam(r)
	err := db.Delete(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
