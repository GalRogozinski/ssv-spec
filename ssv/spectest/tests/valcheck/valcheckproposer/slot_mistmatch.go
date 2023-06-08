package valcheckproposer

import (
	spec "github.com/attestantio/go-eth2-client/spec"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/valcheck"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// SlotMismatch tests Duty.Slot != BeaconBlock.Slot
func SlotMismatch() tests.SpecTest {
	block := testingutils.TestingBeaconBlockV(spec.DataVersionCapella)
	blockBytes, _ := block.Capella.MarshalSSZ()
	data := &types.ConsensusData{
		Duty: types.Duty{
			Type:           types.BNRoleProposer,
			PubKey:         testingutils.TestingValidatorPubKey,
			Slot:           testingutils.TestingDutySlotCapella + 1,
			ValidatorIndex: testingutils.TestingValidatorIndex,
		},
		Version:                    spec.DataVersionCapella,
		PreConsensusJustifications: nil,
		DataSSZ:                    blockBytes,
	}

	input, _ := data.Encode()

	return &valcheck.SpecTest{
		Name:          "proposer value check slot mismatch",
		Network:       types.BeaconTestNetwork,
		BeaconRole:    types.BNRoleProposer,
		Input:         input,
		ExpectedError: "block data slot != duty slot",
	}
}
