package tests

import (
	"github.com/bloxapp/ssv-spec/types"
	"testing"
)

type SpecTest interface {
	TestName() string
	Run(t *testing.T) []types.Encoder
}

type TestF func() SpecTest
