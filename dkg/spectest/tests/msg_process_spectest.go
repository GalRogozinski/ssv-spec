package tests

import (
	"fmt"
	"github.com/bloxapp/ssv-spec/dkg"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
	"github.com/stretchr/testify/require"
	"testing"
)

type MsgProcessingSpecTest struct {
	Name           string
	InputMessages  []*dkg.SignedMessage
	OutputMessages []*dkg.SignedMessage
	KeySet         *testingutils.TestKeySet
	ExpectedError  string
}

func (test *MsgProcessingSpecTest) TestName() string {
	return test.Name
}

func (test *MsgProcessingSpecTest) Run(t *testing.T) {
	node := testingutils.TestingDKGNode(test.KeySet)

	var lastErr error
	for _, msg := range test.InputMessages {
		byts, _ := msg.Encode()
		err := node.ProcessMessage(&types.SSVMessage{
			MsgType: types.DKGMsgType,
			Data:    byts,
		})

		if err != nil {
			lastErr = err
		}
	}

	if len(test.ExpectedError) > 0 {
		require.EqualError(t, lastErr, test.ExpectedError)
	} else {
		require.NoError(t, lastErr)
	}

	// test output message
	broadcastedMsgs := node.GetConfig().Network.(*testingutils.TestingNetwork).BroadcastedDKGMsgs
	if len(test.OutputMessages) > 0 || len(broadcastedMsgs) > 0 {
		require.Len(t, broadcastedMsgs, len(test.OutputMessages))

		for i, msg := range test.OutputMessages {
			r1, _ := msg.GetRoot()
			r2, _ := broadcastedMsgs[i].GetRoot()
			require.EqualValues(t, r1, r2, fmt.Sprintf("output msg %d roots not equal", i))
		}
	}

}
