package valcheckproposer

import (
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/valcheck"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// BlockDataInvalid tests invalid block data
func BlockDataInvalid() tests.SpecTest {
	invalidBlock := &spec.VersionedBeaconBlock{
		// Missing fields and incorrect size
		Version: spec.DataVersionCapella,
		Capella: &capella.BeaconBlock{
			// missing data
			Slot: testingutils.TestingDutySlotCapella,
		},
	}
	invalidBlockBytes, _ := invalidBlock.Capella.MarshalSSZ()
	data := &types.ConsensusData{
		Duty:    *testingutils.TestingProposerDutyV(spec.DataVersionCapella),
		Version: spec.DataVersionCapella,
		DataSSZ: invalidBlockBytes,
	}

	input, _ := data.Encode()

	return &valcheck.SpecTest{
		Name:          "proposer value check block data invalid",
		Network:       types.BeaconTestNetwork,
		BeaconRole:    types.BNRoleProposer,
		Input:         input,
		ExpectedError: "invalid value: could not unmarshal ssz: incorrect size",
	}
}
