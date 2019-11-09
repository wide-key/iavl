package iavl

import (
	amino "github.com/tendermint/go-amino"
)

var cdc *amino.Codec

func init() {
	// NOTE: It's important that there be no conflicts here,
	// as that would change the canonical representations.
	//cdc = amino.NewCodec()
	//RegisterWire(cdc)
}

func RegisterWire(cdc *amino.Codec) {
	// TODO
}
