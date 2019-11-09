package codec

import (
	"io"

	"github.com/coinexchain/codon"
)

func ShowInfo() {
	codon.ShowInfoForVar(nil, RangeProof{})
	codon.ShowInfoForVar(nil, IAVLAbsenceOp{})
	codon.ShowInfoForVar(nil, IAVLValueOp{})
}

var TypeEntryList = []codon.TypeEntry{
	{Alias: "RangeProof", Value: RangeProof{}},
	{Alias: "ProofInnerNode", Value: ProofInnerNode{}},
	{Alias: "ProofLeafNode", Value: ProofLeafNode{}},
	{Alias: "PathToLeaf", Value: PathToLeaf{}},
	{Alias: "IAVLAbsenceOp", Value: IAVLAbsenceOp{}},
	{Alias: "IAVLValueOp", Value: IAVLValueOp{}},
}

func GenerateCodecFile(w io.Writer) {
	codon.GenerateCodecFile(w, nil, nil, TypeEntryList, codon.BridgeLogic, codon.ImportsForBridgeLogic)
}

const MaxSliceLength = 10
const MaxStringLength = 100

