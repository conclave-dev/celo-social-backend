package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ParseRequestParameters parse and return valid params and options from the request body
func ParseRequestParameters(w http.ResponseWriter, r *http.Request, param interface{}) error {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(param)

	// An io.EOF error is returned by Decode() if the body is empty.
	if err != nil && !errors.Is(err, io.EOF) {
		HandleJSONDecodeError(err, w)
		return err
	}

	return nil
}
