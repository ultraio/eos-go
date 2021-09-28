package oracle

import "github.com/ultraio/eos-go"

func NewInit(interval uint8, cache_window uint32, final_price_table_size []uint32, final_moving_average_settings []eos.Asset, ultra_comprehensive_rate_weight uint32) *eos.Action {
	return &eos.Action{
		Account: AN("eosio.oracle"),
		Name:    ActN("init"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      AN("ultra.oracle"),
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(Init{
			Interval:                     interval,
			CacheWindow:                  cache_window,
			FinalPriceTableSize:          final_price_table_size,
			FinalMovingAverageSettings:   final_moving_average_settings,
			UltraComprehensiveRateWeight: ultra_comprehensive_rate_weight,
		}),
	}
}

// Open represents the `open` struct on the `eosio.token` contract.
type Init struct {
	Interval                     uint8       `json:"interval"`
	CacheWindow                  uint32      `json:"cache_window"`
	FinalPriceTableSize          []uint32    `json:"final_price_table_size,omitempty"`
	FinalMovingAverageSettings   []eos.Asset `json:"final_moving_average_settings,omitempty"`
	UltraComprehensiveRateWeight uint32      `json:"ultra_comprehensive_rate_weight"`
}
