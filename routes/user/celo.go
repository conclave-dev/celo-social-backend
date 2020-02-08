package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/stella-zone/celo-social-backend/types"
	"github.com/stella-zone/celo-social-backend/util"
)

func fetchAccount(address string, w http.ResponseWriter) (types.User, error) {
	var user types.User
	accountSummary, err := fetchAccountSummary(address, w)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	metadata, err := fetchMetadata(accountSummary.Data.MetadataURL, w)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	user = types.User{
		AccountSummary: accountSummary.Data,
		Metadata:       metadata,
	}
	return user, nil
}

func fetchAccountSummary(address string, w http.ResponseWriter) (types.AccountSummaryResponse, error) {
	var accountSummaryResponse types.AccountSummaryResponse
	url := fmt.Sprintf("%s/accounts/getAccountSummary", CeloAPI)
	data := []byte(fmt.Sprintf(`{"address":"%s"}`, address))
	reqJSON := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, reqJSON)
	if err != nil {
		util.RespondWithError(err, w)
		return accountSummaryResponse, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return accountSummaryResponse, err
	}
	defer resp.Body.Close()

	d := json.NewDecoder(resp.Body)
	d.DisallowUnknownFields()

	err = d.Decode(&accountSummaryResponse)

	// An io.EOF error is returned by Decode() if the body is empty.
	if err != nil && !errors.Is(err, io.EOF) {
		util.HandleJSONDecodeError(err, w)
		return accountSummaryResponse, err
	}

	return accountSummaryResponse, nil
}

func fetchMetadata(metadataURL string, w http.ResponseWriter) (types.Metadata, error) {
	var metadata types.Metadata

	m, err := util.SendGET(metadataURL)
	if err != nil {
		util.RespondWithError(err, w)
		return metadata, err
	}

	err = json.Unmarshal(m, &metadata)
	if err != nil {
		util.RespondWithError(err, w)
		return metadata, err
	}

	return metadata, nil
}

func parseMetadataClaims(claims types.Claims) []string {
	http := fmt.Sprintf("http://%s", ClaimDomain)
	https := fmt.Sprintf("https://%s", ClaimDomain)
	var params []string

	for _, claim := range claims {
		if claim.Domain == "" {
			continue
		}

		switch true {
		case strings.Contains(claim.Domain, http):
			params = append(params, strings.ReplaceAll(claim.Domain, http, ""))
		case strings.Contains(claim.Domain, https):
			params = append(params, strings.ReplaceAll(claim.Domain, https, ""))
		default:
			break
		}
	}

	return params
}
