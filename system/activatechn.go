package system

import (
	eos "github.com/eoscanada/eos-go"
)

func NewActivatechn(tier eos.AccountName) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("activatechn"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      AN("ultra.eosio"),
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(nil),
	}
	return a
}
