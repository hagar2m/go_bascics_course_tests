package api

import (
	"encoding/json"
	"net/http"

	"test.com/museum/data"
)

func PostApi(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodPost {
		var exhibition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&exhibition)
		if err != nil {
			http.Error(w,err.Error(), http.StatusBadRequest)
			return
		}
		data.Add(exhibition)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data.GetAll())
		// w.Write([]byte("OK"))
	}	else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}