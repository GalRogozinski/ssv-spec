package latemsg

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
	"github.com/herumi/bls-eth-go-binary/bls"
)

// LatePreparePastInstance tests process prepare msg for a previously decided instance
func LatePreparePastInstance() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	allMsgs := testingutils.DecidingMsgsForHeightWithRoot(testingutils.TestingQBFTRootData,
		testingutils.TestingQBFTFullData, testingutils.TestingIdentifier, 5, ks)

	msgPerHeight := make(map[qbft.Height][]*qbft.SignedMessage)
	msgPerHeight[qbft.FirstHeight] = allMsgs[0:7]
	msgPerHeight[1] = allMsgs[7:14]
	msgPerHeight[2] = allMsgs[14:21]
	msgPerHeight[3] = allMsgs[21:28]
	msgPerHeight[4] = allMsgs[28:35]
	msgPerHeight[5] = allMsgs[35:42]

	instanceData := func(height qbft.Height, postRoot string) *tests.RunInstanceData {
		return &tests.RunInstanceData{
			InputValue:    []byte{1, 2, 3, 4},
			InputMessages: msgPerHeight[height],
			ExpectedDecidedState: tests.DecidedState{
				BroadcastedDecided: testingutils.TestingCommitMultiSignerMessageWithHeight(
					[]*bls.SecretKey{ks.Shares[1], ks.Shares[2], ks.Shares[3]},
					[]types.OperatorID{1, 2, 3},
					height,
				),
				DecidedVal: testingutils.TestingQBFTFullData,
				DecidedCnt: 1,
			},
			ControllerPostRoot: postRoot,
		}
	}

	return &tests.ControllerSpecTest{
		Name: "late prepare past instance",
		RunInstanceData: []*tests.RunInstanceData{
			instanceData(qbft.FirstHeight, "24cf697092529cfab3ab06b969d8696692c8bcbb9f41a954f71dc74c3b1d7e97"),
			instanceData(1, "676a681d7e66740832676ed2a7a34d153a64ae06d39872acef4bf0730464da4b"),
			instanceData(2, "20ec3a034efa8b7cebe91e40f56038bb5756750ae619ee090da563ac5049c829"),
			instanceData(3, "ea9ba94292a0ad2a60a4e57ce5c358cdd4ea27c4352eb98e1f6c9205043c3891"),
			instanceData(4, "70e9293510baa12e4861c6557d43d6b1f06c69cc3cc6b9fc7bb610e26de92575"),
			instanceData(5, "8144eb206920903da31bd7a0231cc4a5d93d195669e91deccc96123cfd04e0d5"),
			{
				InputValue: []byte{1, 2, 3, 4},
				InputMessages: []*qbft.SignedMessage{
					testingutils.TestingPrepareMultiSignerMessageWithHeight(
						[]*bls.SecretKey{ks.Shares[4]},
						[]types.OperatorID{4},
						4,
					),
				},
				ControllerPostRoot: "49347722798a798de5088297e0600ca1a2b577e3a5f52a5ea05fb701b9deb0f4",
			},
		},
	}
}
