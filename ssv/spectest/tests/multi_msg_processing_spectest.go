package tests

import (
	"github.com/bloxapp/ssv-spec/types"
	"testing"
)

type MultiMsgProcessingSpecTest struct {
	Name  string
	Tests []*MsgProcessingSpecTest
}

func (tests *MultiMsgProcessingSpecTest) TestName() string {
	return tests.Name
}

func (tests *MultiMsgProcessingSpecTest) Run(t *testing.T) []types.Encoder {
	var runners = make([]types.Encoder, len(tests.Tests))
	for i, test := range tests.Tests {
		t.Run(test.TestName(), func(t *testing.T) {
			runner := test.Run(t)
			runners[i] = runner[0]
		})
	}
	return runners
}
