package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	db "github.com/InviewTeam/BEST_Final/database"
)

func CreateApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	h := r.Header.Get("Authorization")
	if _, err := db.Auth(h); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	app := db.Application{}
	rBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(rBytes, &app)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.CreateApp(app)
	w.WriteHeader(http.StatusOK)
}

func GetApps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var apps []db.ApplicationDB
	apps = db.GetApps()
	req, err := json.Marshal(apps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(req)
}

func UpdateApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	h := r.Header.Get("Authorization")
	tk, err := db.Auth(h)
	if (err != nil) || !db.IsAdmin(tk) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	app := db.ApplicationDB{}
	rBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(rBytes, &app)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	param := getParam(r)
	err = db.UpdateApp(app, param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	h := r.Header.Get("Authorization")
	if _, err := db.Auth(h); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	param := getParam(r)
	err := db.DeleteApp(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
