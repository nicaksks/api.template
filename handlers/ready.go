package handlers

import (
	"encoding/json"
	"net/http"
)

func Ready(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusOK, "message": "Hello, World!"})
}