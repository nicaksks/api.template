package handlers

import (
	"api/database"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func Ready(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusOK, "message": "Hello, World!"})
}

func Err(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusNotFound, "message": "Not Found"})
}

func RegisterAnime(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var header = r.Header.Get("x-anime-token")
	header = strings.TrimSpace(header)

	if header != os.Getenv("TOKEN") {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusForbidden, "message": "Missing auth token"})
		return
	}

	timeNow := time.Now()
	err := database.Register(vars["name"], vars["description"], vars["genre"], timeNow, timeNow)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusNotFound, "message": "Not Found"})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusOK, "message": "Success"})
}

func DeleteAnime(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var header = r.Header.Get("x-anime-token")
	header = strings.TrimSpace(header)

	if header != os.Getenv("TOKEN") {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusForbidden, "message": "Missing auth token"})
		return
	}

	err := database.Delete(vars["name"])
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusNotFound, "message": "Not Found"})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusOK, "message": "Success"})
}

func FindAnime(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	request, err := database.Find(vars["name"])
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusNotFound, "message": "Not Found"})
	}

	if request == nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusNotFound, "message": "Anime not found"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusOK, "message": "Success", "user": request})
}

func FindAll(w http.ResponseWriter, r *http.Request) {

	request, err := database.FindAll()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusNotFound, "message": "Not Found"})
	}

	if request == nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusNotFound, "message": "Animes not found"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusOK, "message": "Success", "users": request})
}

func UpdateByTitle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var header = r.Header.Get("x-anime-token")
	header = strings.TrimSpace(header)

	if header != os.Getenv("TOKEN") {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusForbidden, "message": "Missing auth token"})
		return
	}

	timeNow := time.Now()
	err := database.Update(vars["name"], vars["description"], vars["genre"], timeNow)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusNotFound, "message": "Anime not found"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusOK, "message": "Success"})
}
