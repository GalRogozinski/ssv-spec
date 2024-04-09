package testingutils

import (
	"github.com/bloxapp/ssv-spec/ssv"
	"github.com/bloxapp/ssv-spec/types"
)

var BaseValidator = func(keySet *TestKeySet) *ssv.Validator {
	return ssv.NewValidator(
		NewTestingNetwork(1, keySet.OperatorKeys[1]),
		NewTestingBeaconNode(),
		TestingShare(keySet),
		NewTestingKeyManager(),
		NewTestingOperatorSigner(),
		map[types.BeaconRole]ssv.Runner{
			types.BNRoleAttester:                  AttesterRunner(keySet),
			types.BNRoleProposer:                  ProposerRunner(keySet),
			types.BNRoleAggregator:                AggregatorRunner(keySet),
			types.BNRoleSyncCommittee:             SyncCommitteeRunner(keySet),
			types.BNRoleSyncCommitteeContribution: SyncCommitteeContributionRunner(keySet),
		},
		NewTestingVerifier(),
	)
}
