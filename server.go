package main

import (
	"encoding/json"
	"net/http"
)

type previsionPayload struct {
	Question  string `json:"question"`
	Prevision string `json:"prevision"`
}

func previse(w http.ResponseWriter, req *http.Request) {
	var reqBody struct {
		Question string `json:"question"`
	}

	// Read body
	if req.Body == nil {
		http.Error(w, "No request body", 400)
		return
	}

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Failed parsing request body", 400)
		return
	}

	// Shaking prevision
	r := previsionPayload{
		Question:  reqBody.Question,
		Prevision: Shake(),
	}

	// Sending response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(r)
}

func main() {
	http.HandleFunc("/previse", previse)

	http.ListenAndServe(":8090", nil)
}
