package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespJSON is a fuction that returns the response to the client
func RespJSON(w http.ResponseWriter, code int, message string, data interface{}) {
	var err error

	resp := make(map[string]interface{})
	resp["message"] = message
	resp["data"] = data

	response, err := json.Marshal(resp)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

// RespBadJSON retrns a bad JSON response
func RespBadJSON(w http.ResponseWriter, code int, err error) {
	resp := make(map[string]interface{})
	resp["error"] = err.Error()

	response, err := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json4")
	w.WriteHeader(code)

	_, _ = w.Write(response)
}
