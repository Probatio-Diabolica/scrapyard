package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey -> extracts the api_key from the http request
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	// err := nil
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")

	if len(vals) < 2 {
		if len(vals) == 1 {
			return "", errors.New("missing Api key")
		}
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("auth header's key is malformed")
	}

	return vals[1], nil
}
