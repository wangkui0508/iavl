package iavl

import (
	iavlOrig "github.com/tendermint/iavl/iavl"
)

var (
	NewMutableTree = iavlOrig.NewMutableTree
	NewIAVLValueOp = iavlOrig.NewIAVLValueOp
	NewIAVLAbsenceOp = iavlOrig.NewIAVLAbsenceOp
	IAVLValueOpDecoder = iavlOrig.IAVLValueOpDecoder
	IAVLAbsenceOpDecoder = iavlOrig.IAVLAbsenceOpDecoder

	ErrVersionDoesNotExist = iavlOrig.ErrVersionDoesNotExist
)


const (
	ProofOpIAVLValue = iavlOrig.ProofOpIAVLValue
	ProofOpIAVLAbsence = iavlOrig.ProofOpIAVLAbsence
)

type (
	RangeProof = iavlOrig.RangeProof
	MutableTree = iavlOrig.MutableTree
	ImmutableTree = iavlOrig.ImmutableTree
)

