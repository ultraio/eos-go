package system

import (
	eos "github.com/eoscanada/eos-go"
)

func NewSetramcurve(core_reserve eos.Asset, ram_supply int64, connector_weight float64) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setramcurve"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      AN("ultra.eosio"),
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(Setramcurve{
			CoreReserve:     core_reserve,
			RamSupply:       ram_supply,
			ConnectorWeight: connector_weight,
		}),
	}
	return a
}

type Setramcurve struct {
	CoreReserve     eos.Asset `json:"core_reserve"`
	RamSupply       int64     `json:"ram_supply"`
	ConnectorWeight float64   `json:"connector_weight"`
}
