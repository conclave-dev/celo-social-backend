package types

import (
	"github.com/ethereum/go-ethereum/common"
	kvstore "github.com/stella-zone/celo-social-backend/kvstore"
	"github.com/stella-zone/go-celo/types"
)

type JSONResponse struct {
	Data interface{} `json:"data"`
}

type User struct {
	Claimed        bool            `json:"claimed"`
	Hash           string          `json:"hash"`
	Profile        kvstore.Profile `json:"profile"`
	AccountSummary types.Account   `json:"accountSummary"`
	Metadata       Metadata        `json:"metadata"`
}

// Metadata is JSON fetched from a user account summary's metadata URL
type Metadata struct {
	Claims Claims `json:"claims"`
	Meta   Meta   `json:"meta"`
}

type Claims []Claim

// Claim is a information that a user has claimed in their Celo account
type Claim struct {
	Address   string `json:"address,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
	Name      string `json:"name,omitempty,omitempty"`
	Type      string `json:"type,omitempty,omitempty"`
	URL       string `json:"url,omitempty,omitempty"`
	Domain    string `json:"domain,omitempty,omitempty"`
}

// Meta is data related to the user making the claims
type Meta struct {
	Address   common.Address `json:"address"`
	Signature string         `json:"signature"`
}
