package profile

import "github.com/ethereum/go-ethereum/common"

type Profile struct {
	Address    common.Address `json:"address"`
	AddressDos common.Address `json:"addressDos"`
	PhotoURL   string         `json:"photoURL"`
}
