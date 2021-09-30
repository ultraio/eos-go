package eba

import (
	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
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

type Regidp struct {
	IdProviders []Provider `json:"id_providers,omitempty"`
}
