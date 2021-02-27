package httperror

import (
	"encoding/json"
	"net/http"
)

// WrapError wrap error response http
func WrapError(resp http.ResponseWriter, status int, err string) {
	resp.WriteHeader(status)
	json.NewEncoder(resp).Encode(map[string]string{
		"error": err,
	})
}
