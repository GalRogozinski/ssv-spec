package testingutils

import (
	"bytes"
	"crypto/sha256"

	"github.com/pkg/errors"

	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types"
)

var TestingQBFTFullData = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var TestingQBFTRootData = func() [32]byte {
	return sha256.Sum256(TestingQBFTFullData)
}()

var TestingConfig = func(keySet *TestKeySet) *qbft.Config {
	return &qbft.Config{
		OperatorSigner: NewTestingOperatorSigner(keySet, 1),
		SigningPK:      keySet.Shares[1].GetPublicKey().Serialize(),
		Domain:         TestingSSVDomainType,
		ValueCheckF: func(data []byte) error {
			if bytes.Equal(data, TestingInvalidValueCheck) {
				return errors.New("invalid value")
			}

			// as a base validation we do not accept nil values
			if len(data) == 0 {
				return errors.New("invalid value")
			}
			return nil
		},
		ProposerF: func(state *qbft.State, round qbft.Round) types.OperatorID {
			return 1
		},
		Network:           NewTestingNetwork(1, keySet.OperatorKeys[1]),
		Timer:             NewTestingTimer(),
		SignatureVerifier: NewTestingVerifier(),
	}
}

var TestingInvalidValueCheck = []byte{1, 1, 1, 1}

var TestingShare = func(keysSet *TestKeySet) *types.Share {

	// Decode validator public key
	bytes := keysSet.ValidatorPK.Serialize()
	blsPubKeyBytes := [48]byte{}
	copy(blsPubKeyBytes[:], bytes)

	return &types.Share{
		ValidatorIndex:      TestingValidatorIndex,
		ValidatorPubKey:     blsPubKeyBytes,
		SharePubKey:         keysSet.Shares[1].GetPublicKey().Serialize(),
		Committee:           keysSet.Committee(),
		Quorum:              keysSet.Threshold,
		DomainType:          TestingSSVDomainType,
		FeeRecipientAddress: TestingFeeRecipient,
		Graffiti:            []byte{},
	}
}

var TestingOperator = func(keysSet *TestKeySet) *types.Operator {
	committeeMembers := []*types.CommitteeMember{}

	for _, key := range keysSet.Committee() {

		// Encode member's public key
		pkBytes, err := types.MarshalPublicKey(keysSet.OperatorKeys[key.Signer])
		if err != nil {
			panic(err)
		}

		committeeMembers = append(committeeMembers, &types.CommitteeMember{
			OperatorID:        key.Signer,
			SSVOperatorPubKey: pkBytes,
		})
	}

	opIds := []types.OperatorID{}
	for _, key := range keysSet.Committee() {
		opIds = append(opIds, key.Signer)
	}

	operatorPubKeyBytes, err := types.MarshalPublicKey(keysSet.OperatorKeys[1])
	if err != nil {
		panic(err)
	}

	return &types.Operator{
		OperatorID:        1,
		ClusterID:         types.GetClusterID(opIds),
		SSVOperatorPubKey: operatorPubKeyBytes,
		Quorum:            keysSet.Threshold,
		Committee:         committeeMembers,
	}
}

var BaseInstance = func() *qbft.Instance {
	return baseInstance(TestingOperator(Testing4SharesSet()), Testing4SharesSet(), []byte{1, 2, 3, 4})
}

var SevenOperatorsInstance = func() *qbft.Instance {
	return baseInstance(TestingOperator(Testing7SharesSet()), Testing7SharesSet(), []byte{1, 2, 3, 4})
}

var TenOperatorsInstance = func() *qbft.Instance {
	return baseInstance(TestingOperator(Testing10SharesSet()), Testing10SharesSet(), []byte{1, 2, 3, 4})
}

var ThirteenOperatorsInstance = func() *qbft.Instance {
	return baseInstance(TestingOperator(Testing13SharesSet()), Testing13SharesSet(), []byte{1, 2, 3, 4})
}

var baseInstance = func(operator *types.Operator, keySet *TestKeySet, identifier []byte) *qbft.Instance {
	ret := qbft.NewInstance(TestingConfig(keySet), operator, identifier, qbft.FirstHeight)
	ret.StartValue = TestingQBFTFullData
	return ret
}

func NewTestingQBFTController(
	identifier []byte,
	share *types.Operator,
	config qbft.IConfig,
) *qbft.Controller {
	return qbft.NewController(
		identifier,
		share,
		config,
	)
}
