package system

import (
	eos "github.com/eoscanada/eos-go"
)

func NewNewactconfig(cost eos.Asset, oracle eos.AccountName, candidate_moving_average []eos.Asset) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("newactconfig"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      AN("ultra.eosio"),
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(Newactconfig{
			Cost:                   cost,
			Oracle:                 oracle,
			CandidateMovingAverage: candidate_moving_average,
		}),
	}
	return a
}

type Newactconfig struct {
	Cost                   eos.Asset       `json:"cost"`
	Oracle                 eos.AccountName `json:"tier"`
	CandidateMovingAverage []eos.Asset     `json:"candidate_moving_average"`
}
