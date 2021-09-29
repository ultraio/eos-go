package oracle

import "github.com/eoscanada/eos-go"

func NewRegexchange(exchange eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: AN("eosio.oracle"),
		Name:    ActN("init"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      AN("ultra.oracle"),
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(Regexchange{
			Exchange: exchange,
		}),
	}
}

// Open represents the `open` struct on the `eosio.token` contract.
type Regexchange struct {
	Exchange eos.AccountName `json:"exchange"`
}
