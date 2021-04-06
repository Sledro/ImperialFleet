package utils

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// NewAPIError - Formats and returns an API error
func NewAPIError(filename string, function string, err error, w http.ResponseWriter) error {
	// log for devs
	log.Error("❌ " + filename + " -> " + function + "() -> " + err.Error())

	// return to user
	formatError := map[string]string{
		"error": err.Error(),
	}
	return json.NewEncoder(w).Encode(formatError)
}
