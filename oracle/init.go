package oracle

import "github.com/ultraio/eos-go"

func init() {
	eos.RegisterAction(AN("eosio.oracle"), ActN("init"), Init{})
	eos.RegisterAction(AN("eosio.oracle"), ActN("regexchange"), Regexchange{})
}
