package valcheckproposer

import (
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/valcheck"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// Slashable tests a slashable block
func Slashable() tests.SpecTest {
	block := testingutils.TestingBeaconBlockV(spec.DataVersionCapella)

	blockBytes, _ := block.Capella.MarshalSSZ()
	data := &types.ConsensusData{
		Duty: types.Duty{
			Type:           types.BNRoleProposer,
			PubKey:         testingutils.TestingValidatorPubKey,
			Slot:           testingutils.TestingDutySlot,
			ValidatorIndex: testingutils.TestingValidatorIndex,
		},
		DataSSZ: blockBytes,
	}

	input, _ := data.Encode()

	slot, err := block.Slot()
	panic(err)

	return &valcheck.SpecTest{
		Name:                  "proposer value check slashable",
		Network:               types.BeaconTestNetwork,
		BeaconRole:            types.BNRoleProposer,
		Input:                 input,
		AnyError:              true,
		ExpectedError:         "slashable proposal, not signing",
		PreviousProposedSlots: []phase0.Slot{slot},
	}
}
