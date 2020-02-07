package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// ParseJSONBody parse and return valid params and options from the request body
func ParseJSONBody(w http.ResponseWriter, jsonBody io.Reader, param interface{}) error {
	d := json.NewDecoder(jsonBody)
	d.DisallowUnknownFields()
	err := d.Decode(param)

	// An io.EOF error is returned by Decode() if the body is empty.
	if err != nil && !errors.Is(err, io.EOF) {
		HandleJSONDecodeError(err, w)
		return err
	}

	return nil
}

func SendPOST(url string, data []byte, param interface{}, w http.ResponseWriter) error {
	reqJSON := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, reqJSON)
	if err != nil {
		RespondWithError(err, w)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = ParseJSONBody(w, resp.Body, &param)
	if err != nil {
		return err
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
