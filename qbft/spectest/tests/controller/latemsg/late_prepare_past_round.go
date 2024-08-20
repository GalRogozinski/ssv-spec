package latemsg

import (
	"crypto/rsa"

	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests"
	"github.com/ssvlabs/ssv-spec/types"
	"github.com/ssvlabs/ssv-spec/types/testingutils"
)

// LatePreparePastRound tests process late prepare msg for an instance which just decided for a round < decided round
func LatePreparePastRound() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	rcMsgs := []*types.SignedSSVMessage{
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[1], types.OperatorID(1), 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[2], types.OperatorID(2), 2),
		testingutils.TestingRoundChangeMessageWithRound(ks.OperatorKeys[3], types.OperatorID(3), 2),
	}

	msgs := []*types.SignedSSVMessage{
		testingutils.TestingProposalMessage(ks.OperatorKeys[1], types.OperatorID(1)),

		testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(1)),
		testingutils.TestingPrepareMessage(ks.OperatorKeys[2], types.OperatorID(2)),
	}
	msgs = append(msgs, rcMsgs...)
	msgs = append(msgs, []*types.SignedSSVMessage{
		testingutils.TestingProposalMessageWithRoundAndRC(ks.OperatorKeys[1], types.OperatorID(1), 2,
			testingutils.MarshalJustifications(rcMsgs)),

		testingutils.TestingPrepareMessageWithRound(ks.OperatorKeys[1], types.OperatorID(1), 2),
		testingutils.TestingPrepareMessageWithRound(ks.OperatorKeys[2], types.OperatorID(2), 2),
		testingutils.TestingPrepareMessageWithRound(ks.OperatorKeys[3], types.OperatorID(3), 2),

		testingutils.TestingCommitMessageWithRound(ks.OperatorKeys[1], types.OperatorID(1), 2),
		testingutils.TestingCommitMessageWithRound(ks.OperatorKeys[2], types.OperatorID(2), 2),
		testingutils.TestingCommitMessageWithRound(ks.OperatorKeys[3], types.OperatorID(3), 2),

		testingutils.TestingPrepareMessage(ks.OperatorKeys[1], types.OperatorID(1)),
	}...)

	return &tests.ControllerSpecTest{
		Name: "late prepare past round",
		RunInstanceData: []*tests.RunInstanceData{
			{
				InputValue:    []byte{1, 2, 3, 4},
				InputMessages: msgs,
				ExpectedDecidedState: tests.DecidedState{
					DecidedVal: testingutils.TestingQBFTFullData,
					DecidedCnt: 1,
					BroadcastedDecided: testingutils.TestingCommitMultiSignerMessageWithRound(
						[]*rsa.PrivateKey{ks.OperatorKeys[1], ks.OperatorKeys[2], ks.OperatorKeys[3]},
						[]types.OperatorID{1, 2, 3},
						2,
					),
				},
			},
		},
		ExpectedError: "not processing consensus message since instance is already decided",
	}
}
