package handlers

import (
	"api/database"
	"encoding/json"
	"net/http"
)

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
