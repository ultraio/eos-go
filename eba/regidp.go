package eba

import (
	eos "github.com/ultraio/eos-go"
	"github.com/ultraio/eos-go/ecc"
)

func NewRegidp(id_providers []Provider) *eos.Action {
	return &eos.Action{
		Account: AN("eosio.eba"),
		Name:    ActN("regidp"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      AN("ultra.eba"),
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(Regidp{
			IdProviders: id_providers,
		}),
	}
}

type Provider struct {
	Account eos.AccountName `json:"account"`
	Key     ecc.PublicKey   `json:"key"`
}

// Open represents the `open` struct on the `eosio.token` contract.
type Regidp struct {
	IdProviders []Provider `json:"id_providers,omitempty"`
}
