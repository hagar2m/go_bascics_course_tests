package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"test.com/museum/data"
)

func GetApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// api/exhibitions?id=3
	idFromQuery := r.URL.Query()["id"]
	// (id) here is the index
	if idFromQuery != nil  {
		index, err := strconv.Atoi(idFromQuery[0])
		if err == nil && index < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[index])
		}	else {
			http.Error(w, "Invalid exhibition", http.StatusBadRequest)
		}
	}	else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}