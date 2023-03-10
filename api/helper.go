package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type APIError struct {
	Error struct {
		// Message is a developer friendly representation of the error issue
		Message string `json:"message"`
	} `json:"error"`
}

// RespondJson converts a Go value to JSON and sends it to the client.
func respondJson(ctx context.Context, w http.ResponseWriter, data interface{}, statusCode int) error {

	// Check to see if the json has already been encoded.
	jsonData, ok := data.([]byte)
	if !ok {
		// Convert the response value to JSON.
		var err error
		jsonData, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", MIMEApplicationJSON)

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

func handleIfError(w http.ResponseWriter, err error, code int) bool {
	if err == nil {
		return false
	}
	apiErr := APIError{}
	apiErr.Error.Message = err.Error()

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(apiErr); err != nil {
		log.Println(`ERROR`, err.Error())
	}
	return true
}
