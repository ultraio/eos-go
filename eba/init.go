package eba

import "github.com/ultraio/eos-go"

func init() {
	eos.RegisterAction(AN("eosio.eba"), ActN("regidp"), Regidp{})
}

var AN = eos.AN
var PN = eos.PN
var ActN = eos.ActN
