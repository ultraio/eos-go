package system

import (
	eos "github.com/eoscanada/eos-go"
)

func NewSetramtrade(state bool) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setramtrade"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      AN("ultra.eosio"),
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(Setramtrade{
			State: state,
		}),
	}
	return a
}

type Setramtrade struct {
	State bool `json:"state"`
}
