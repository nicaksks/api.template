package handlers

import (
	"encoding/json"
	"net/http"
)

func Err(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{"code": http.StatusNotFound, "message": "Not Found"})
}
