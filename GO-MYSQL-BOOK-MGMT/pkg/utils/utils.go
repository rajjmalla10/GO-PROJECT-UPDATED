// utils.go

package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody parses the JSON body of an HTTP request into the provided 'x' interface.
func ParseBody(r *http.Request, x interface{}) error {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}
