package timeoutduration

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/timeout"
	"github.com/bloxapp/ssv-spec/types"
)

// Round9Duration tests timeout duration for round 9 where the current time is the duty start time
func Round9DurationOnDutyStartTime() *tests.MultiSpecTest {
	testingNetwork := types.HoleskyNetwork
	height := qbft.Height(40)
	var round qbft.Round = 9
	dutyStartTime := uint64(testingNetwork.EstimatedTimeAtSlot(phase0.Slot(height)))

	return &tests.MultiSpecTest{
		Name: "round 9 duration duty start time",
		Tests: []tests.SpecTest{
			&timeout.TimeoutDurationTest{
				Name:             "sync committee",
				Role:             types.BNRoleSyncCommittee,
				Height:           height,
				Round:            round,
				Network:          testingNetwork,
				CurrentTime:      dutyStartTime,
				ExpectedDuration: 140,
			},
			&timeout.TimeoutDurationTest{
				Name:             "sync committee contribution",
				Role:             types.BNRoleSyncCommitteeContribution,
				Height:           height,
				Round:            round,
				Network:          testingNetwork,
				CurrentTime:      dutyStartTime,
				ExpectedDuration: 144,
			},
			&timeout.TimeoutDurationTest{
				Name:             "attester",
				Role:             types.BNRoleAttester,
				Height:           height,
				Round:            round,
				Network:          testingNetwork,
				CurrentTime:      dutyStartTime,
				ExpectedDuration: 140,
			},
			&timeout.TimeoutDurationTest{
				Name:             "aggregator",
				Role:             types.BNRoleAggregator,
				Height:           height,
				Round:            round,
				Network:          testingNetwork,
				CurrentTime:      dutyStartTime,
				ExpectedDuration: 144,
			},
			&timeout.TimeoutDurationTest{
				Name:             "block proposer",
				Role:             types.BNRoleProposer,
				Height:           height,
				Round:            round,
				Network:          testingNetwork,
				CurrentTime:      dutyStartTime,
				ExpectedDuration: 120,
			},
		},
	}

}
