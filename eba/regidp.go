package eba

import (
	eos "github.com/eoscanada/eos-go"
)

func NewRegidp(id_providers []eos.Provider) *eos.Action {
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

// Open represents the `open` struct on the `eosio.token` contract.
type Regidp struct {
	IdProviders []eos.Provider `json:"id_providers,omitempty"`
}
