package util

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// ParseJSONBody parse and return valid params and options from the request body
func ParseJSONBody(jsonBody io.Reader, v interface{}) error {
	d := json.NewDecoder(jsonBody)
	d.DisallowUnknownFields()
	err := d.Decode(v)

	// An io.EOF error is returned by Decode() if the body is empty.
	if err != nil && !errors.Is(err, io.EOF) {
		return HandleJSONDecodeError(err)
	}

	return nil
}

func SendGET(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
