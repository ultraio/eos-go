package system

import (
	eos "github.com/eoscanada/eos-go"
)

func NewCreatetier(tier eos.AccountName, max_free_permission_objects uint64, max_free_shared_keys uint64, max_free_permission_levels uint64, max_free_waits uint64, max_free_permission_link_objects uint64) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("createtier"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      AN("ultra.eosio"),
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(Setramsponsr{
			Tier:                         tier,
			MaxFreePermissionObjects:     max_free_permission_objects,
			MaxFreeSharedKeys:            max_free_shared_keys,
			MaxFreePermissionLevels:      max_free_permission_levels,
			MaxFreeWaits:                 max_free_waits,
			MaxFreePermissionLinkObjects: max_free_permission_link_objects,
		}),
	}
	return a
}

type Createtier struct {
	Tier                         eos.AccountName `json:"tier"`
	MaxFreePermissionObjects     uint64          `json:"max_free_permission_objects"`
	MaxFreeSharedKeys            uint64          `json:"max_free_shared_keys"`
	MaxFreePermissionLevels      uint64          `json:"max_free_permission_levels"`
	MaxFreeWaits                 uint64          `json:"max_free_waits"`
	MaxFreePermissionLinkObjects uint64          `json:"max_free_permission_link_objects"`
}
