package token

import eos "github.com/eoscanada/eos-go"

func NewOpen(owner eos.AccountName, symbol eos.Symbol, ram_payer eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: AN("eosio.token"),
		Name:    ActN("open"),
		Authorization: []eos.PermissionLevel{
			{Actor: ram_payer, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(Open{
			Owner:		owner,
			Symbol: 	symbol,
			RamPayer:	ram_payer,
		}),
	}
}

// Open represents the `open` struct on the `eosio.token` contract.
type Open struct {
	Owner       eos.AccountName	`json:"owner"`
	Symbol		eos.Symbol		`json:"symbol"`
	RamPayer	eos.AccountName	`json:"ram_payer"`
}
