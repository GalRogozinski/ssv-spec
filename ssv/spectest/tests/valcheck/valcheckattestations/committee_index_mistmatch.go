package valcheckattestations

import (
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/valcheck"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// CommitteeIndexMismatch tests Duty.CommitteeIndex != AttestationData.Index
func CommitteeIndexMismatch() *valcheck.SpecTest {
	data := &types.ConsensusData{
		Duty: &types.Duty{
			Type:                    types.BNRoleAttester,
			PubKey:                  testingutils.TestingValidatorPubKey,
			Slot:                    testingutils.TestingDutySlot,
			ValidatorIndex:          testingutils.TestingValidatorIndex,
			CommitteeIndex:          2,
			CommitteesAtSlot:        36,
			CommitteeLength:         128,
			ValidatorCommitteeIndex: 11,
		},
		AttestationData: &spec.AttestationData{
			Slot:            testingutils.TestingDutySlot,
			Index:           3,
			BeaconBlockRoot: spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
			Source: &spec.Checkpoint{
				Epoch: 0,
				Root:  spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
			},
			Target: &spec.Checkpoint{
				Epoch: 1,
				Root:  spec.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
			},
		},
	}

	input, _ := data.Encode()

	return &valcheck.SpecTest{
		Name:          "attestation value check committee index mismatch",
		Network:       types.PraterNetwork,
		BeaconRole:    types.BNRoleAttester,
		Input:         input,
		ExpectedError: "attestation data CommitteeIndex != duty CommitteeIndex",
	}
}
