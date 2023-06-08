package valcheckproposer

import (
	"github.com/attestantio/go-eth2-client/spec"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/valcheck"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// Valid tests valid data
func Valid() tests.SpecTest {
	return &valcheck.SpecTest{
		Name:       "proposed block value check valid",
		Network:    types.BeaconTestNetwork,
		BeaconRole: types.BNRoleProposer,
		Input:      testingutils.TestProposerConsensusDataBytsV(spec.DataVersionCapella),
		AnyError:   false,
	}
}
