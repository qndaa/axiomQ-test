package utils

import (
	"encoding/json"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func BadRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(400)
	w.Write([]byte("Invalid ID!"))
}
