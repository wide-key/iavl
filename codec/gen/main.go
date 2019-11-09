package main

import (
	"os"

	"github.com/tendermint/iavl/codec"
)

func main() {
	//codec.ShowInfo()
	genCode()
}

func genCode() {
	codec.GenerateCodecFile(os.Stdout)
}
