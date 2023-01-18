package handlers

import (
	"api/database"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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
