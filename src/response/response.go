package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorAPI Struct represente a error response for API
type ErrorAPI struct {
	Error string `json:"error"`
}

// JSON return a respost in json format
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}

}

// TreatStatusCodeError treat requests with status code 400 or greater
func TreatStatusCodeError(w http.ResponseWriter, r *http.Response) {
	var err ErrorAPI

	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)

}
