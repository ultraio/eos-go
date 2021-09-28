package system

import (
	eos "github.com/eoscanada/eos-go"
)

func NewSetramsponsr(ramSponsor eos.AccountName) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setramsponsr"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      AN("ultra.eosio"),
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(Setramsponsr{
			RAMSponsor: ramSponsor,
		}),
	}
	return a
}

type Setramsponsr struct {
	RAMSponsor eos.AccountName `json:"ram_sponsor"`
}
