package system

import (
	eos "github.com/eoscanada/eos-go"
)

func NewSetdeftier(tier eos.AccountName) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setdeftier"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      AN("ultra.eosio"),
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(Setdeftier{
			Tier: tier,
		}),
	}
	return a
}

type Setdeftier struct {
	Tier eos.AccountName `json:"tier"`
}
