package controller

import (
	"encoding/json"
	"net/http"

	"github.com/guspanc/go-crud-pets-api/cmd/api/models"
)

func encodeResponse(w http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}

func decodeRequest(r *http.Request, i interface{}) error {
	dec := json.NewDecoder(r.Body)
	return dec.Decode(&i)
}

func handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case *models.APIError:
		w.WriteHeader(http.StatusBadRequest)
		encodeResponse(w, e)
	default:
		apiError := models.NewAPIError("InternalServerError", e.Error())
		w.WriteHeader(http.StatusInternalServerError)
		encodeResponse(w, apiError)
	}
}
