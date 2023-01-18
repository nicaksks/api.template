package handlers

import (
	"api/database"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

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
