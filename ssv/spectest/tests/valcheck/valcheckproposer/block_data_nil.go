package valcheckproposer

import (
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/valcheck"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// BlockDataNil tests nil block data
func BlockDataNil() tests.SpecTest {
	data := &types.ConsensusData{
		Duty:                       *testingutils.TestingProposerDutyV(spec.DataVersionCapella),
		Version:                    spec.DataVersionCapella,
		PreConsensusJustifications: nil,
		DataSSZ:                    nil,
	}

	input, _ := data.Encode()

	return &valcheck.SpecTest{
		Name:          "proposer value check block data nil",
		Network:       types.BeaconTestNetwork,
		BeaconRole:    types.BNRoleProposer,
		Input:         input,
		ExpectedError: "invalid value: could not unmarshal ssz: incorrect size",
	}
}
