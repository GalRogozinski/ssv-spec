package prepare

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// AllPreparesSentLateCommit is a spec test that checks the case where all prepares are sent and quorum event is triggered more than once.
// A commit message was seen only after the last prepare
func AllPreparesSentLateCommit() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()
	pre := testingutils.BaseInstance()
	msgs := []*qbft.SignedMessage{
		testingutils.TestingProposalMessage(ks.Shares[1], 1),

		testingutils.TestingPrepareMessage(ks.Shares[1], 1),
		testingutils.TestingPrepareMessage(ks.Shares[2], 2),
		testingutils.TestingPrepareMessage(ks.Shares[3], 3),

		testingutils.TestingPrepareMessage(ks.Shares[4], 4),
		testingutils.TestingCommitMessage(ks.Shares[1], 1),
	}
	return &tests.MsgProcessingSpecTest{
		Name:          "all prepares sent",
		Pre:           pre,
		PostRoot:      "a3b1009cdc2ee22b439eab30cb89aa368171d1c87589c756300286393bd78631",
		InputMessages: msgs,
		OutputMessages: []*qbft.SignedMessage{
			testingutils.TestingPrepareMessage(ks.Shares[1], 1),
			testingutils.TestingCommitMessage(ks.Shares[1], 1),
			// ISSUE 214: we should have only commit broadcasted
			testingutils.TestingCommitMessage(ks.Shares[1], 1),
		},
	}
}
