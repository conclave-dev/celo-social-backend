package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stella-zone/go-celo/types"
)

type JSONResponse struct {
	Data interface{} `json:"data"`
}

type UserResponse struct {
	Hash           string        `json:"hash"`
	Profile        Profile       `json:"profile"`
	AccountSummary types.Account `json:"accountSummary"`
	Metadata       Metadata      `json:"metadata"`
}

type User struct {
	Hash string `json:"hash"`
}

// Profile is mutable user data
type Profile struct {
	Name        string   `json:"name"`
	PhotoURL    string   `json:"photoURL"`
	Email       string   `json:"email"`
	Description string   `json:"description"`
	Members     []Member `json:"members"`
}

type AccountSummaryResponse struct {
	Data types.Account `json:"data"`
}

type AddressData struct {
	AccountSummary types.Account `json:"accountSummary"`
	Metadata       Metadata      `json:"metadata"`
	ClaimParams    []string      `json:"claimParams"`
}

// Metadata is JSON fetched from a user account summary's metadata URL
type Metadata struct {
	Claims Claims `json:"claims"`
	Meta   Meta   `json:"meta"`
}

type Claims []Claim

// Claim is a information that a user has claimed in their Celo account
type Claim struct {
	Address   string `json:"address"`
	Timestamp int    `json:"timestamp"`
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	URL       string `json:"url,omitempty"`
	Domain    string `json:"domain,omitempty"`
}

// Meta is data related to the user making the claims
type Meta struct {
	Address   common.Address `json:"address"`
	Signature string         `json:"signature"`
}

// Member is a member that the user has added
type Member struct {
	Name    string `json:"name"`
	Role    string `json:"role"`
	Email   string `json:"email"`
	Website string `json:"website"`
}
