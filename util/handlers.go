package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// RespondWithError acts as a wrapper for writing out an erroneous response
func RespondWithError(err error, w http.ResponseWriter) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// HandleJSONDecodeError provides helpful responses to the client based on a series of error cases
// Credit goes to https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
func HandleJSONDecodeError(err error) error {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	// Catch any syntax errors in the JSON and send an error message
	// which interpolates the location of the problem to make it
	// easier for the client to fix.
	case errors.As(err, &syntaxError):
		return errors.New(fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset))

	// In some circumstances Decode() may also return an
	// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
	// is an open issue regarding this at
	// https://github.com/golang/go/issues/25956.
	case errors.Is(err, io.ErrUnexpectedEOF):
		return errors.New(fmt.Sprintf("Request body contains badly-formed JSON"))

	// Catch any type errors, like trying to assign a string in the
	// JSON request body to a int field in our Person struct. We can
	// interpolate the relevant field name and position into the error
	// message to make it easier for the client to fix.
	case errors.As(err, &unmarshalTypeError):
		return errors.New(fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset))

	// Catch the error caused by extra unexpected fields in the request
	// body. We extract the field name from the error message and
	// interpolate it in our custom error message. There is an open
	// issue at https://github.com/golang/go/issues/29035 regarding
	// turning this into a sentinel error.
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		return errors.New(strings.TrimPrefix(err.Error(), "json: unknown field "))

	// Catch the error caused by the request body being too large. Again
	// there is an open issue regarding turning this into a sentinel
	// error at https://github.com/golang/go/issues/30715.
	case err.Error() == "http: request body too large":
		return errors.New("Request body must not be larger than 1MB")

	// Otherwise default to logging the error and sending a 500 Internal
	// Server Error response.
	default:
		return errors.New(fmt.Sprintf(err.Error()))
	}
}
