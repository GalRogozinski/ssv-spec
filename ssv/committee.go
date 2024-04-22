package ssv

import (
	"fmt"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/pkg/errors"
)

type Committee struct {
	Runners                 map[spec.Slot]*CommitteeRunner
	Operator                types.Operator
	SignatureVerifier       types.SignatureVerifier
	CreateRunnerFn          func() *CommitteeRunner
	HighestAttestingSlotMap map[types.ValidatorPK]spec.Slot
}

// NewCommittee creates a new cluster
func NewCommittee(
	operator types.Operator,
	verifier types.SignatureVerifier,
	createRunnerFn func() *CommitteeRunner,
) *Committee {
	return &Committee{
		Runners:           make(map[spec.Slot]*CommitteeRunner),
		Operator:          operator,
		SignatureVerifier: verifier,
		CreateRunnerFn:    createRunnerFn,
	}

}

// StartDuty starts a new duty for the given slot
func (c *Committee) StartDuty(duty *types.CommitteeDuty) error {
	if _, exists := c.Runners[duty.Slot]; exists {
		return errors.New(fmt.Sprintf("CommitteeRunner for slot %d already exists", duty.Slot))
	}
	c.Runners[duty.Slot] = c.CreateRunnerFn()
	validatorToStopMap := make(map[spec.Slot]types.ValidatorPK)
	// Filter old duties based on highest attesting slot
	duty, validatorToStopMap, c.HighestAttestingSlotMap = FilterCommitteeDuty(duty, c.HighestAttestingSlotMap)
	// Stop validators with old duties
	c.stopDuties(validatorToStopMap)
	c.updateAttestingSlotMap(duty)
	return c.Runners[duty.Slot].StartNewDuty(duty)
}

func (c *Committee) stopDuties(validatorToStopMap map[spec.Slot]types.ValidatorPK) {
	for slot, validator := range validatorToStopMap {
		runner, exists := c.Runners[slot]
		if exists {
			runner.StopDuty(validator)
		}
	}
}

// FilterCommitteeDuty filters the committee duty. It returns the new duty, the validators to stop and the highest attesting slot map
func FilterCommitteeDuty(duty *types.CommitteeDuty, slotMap map[types.ValidatorPK]spec.Slot) (
	*types.CommitteeDuty,
	map[spec.Slot]types.ValidatorPK,
	map[types.ValidatorPK]spec.Slot) {
	validatorsToStop := make(map[spec.Slot]types.ValidatorPK)

	for i, beaconDuty := range duty.BeaconDuties {
		validatorPK := types.ValidatorPK(beaconDuty.PubKey)
		slot, exists := slotMap[validatorPK]
		if exists {
			if slot < beaconDuty.Slot {
				validatorsToStop[beaconDuty.Slot] = validatorPK
				slot = beaconDuty.Slot
			} else { // else don't run duty with old slot
				duty.BeaconDuties[i] = nil
			}
		}
	}
	return duty, validatorsToStop, slotMap
}

// ProcessMessage processes Network Message of all types
func (c *Committee) ProcessMessage(signedSSVMessage *types.SignedSSVMessage) error {
	// Validate message
	if err := signedSSVMessage.Validate(); err != nil {
		return errors.Wrap(err, "invalid SignedSSVMessage")
	}

	// Verify SignedSSVMessage's signature
	if err := c.SignatureVerifier.Verify(signedSSVMessage, c.Operator.Committee); err != nil {
		return errors.Wrap(err, "SignedSSVMessage has an invalid signature")
	}

	msg := signedSSVMessage.SSVMessage

	switch msg.GetType() {
	case types.SSVConsensusMsgType:
		qbftMsg := &qbft.Message{}
		if err := qbftMsg.Decode(msg.GetData()); err != nil {
			return errors.Wrap(err, "could not get consensus Message from network Message")
		}
		runner := c.Runners[spec.Slot(qbftMsg.Height)]
		// TODO: check if runner is nil
		return runner.ProcessConsensus(signedSSVMessage)
	case types.SSVPartialSignatureMsgType:
		pSigMessages := &types.PartialSignatureMessages{}
		if err := pSigMessages.Decode(msg.GetData()); err != nil {
			return errors.Wrap(err, "could not get post consensus Message from network Message")
		}
		if pSigMessages.Type == types.PostConsensusPartialSig {
			runner := c.Runners[pSigMessages.Slot]
			// TODO: check if runner is nil
			return runner.ProcessPostConsensus(pSigMessages)
		}
	default:
		return errors.New("unknown msg")
	}
	return nil

}

func (c *Committee) validateMessage(msg *types.SSVMessage) error {
	if !c.Operator.ClusterID.MessageIDBelongs(msg.GetID()) {
		return errors.New("Message ID does not match cluster IF")
	}
	if len(msg.GetData()) == 0 {
		return errors.New("msg data is invalid")
	}

	return nil
}

// updateAttestingSlotMap updates the highest attesting slot map from beacon duties
func (c *Committee) updateAttestingSlotMap(duty *types.CommitteeDuty) {
	for _, beaconDuty := range duty.BeaconDuties {
		if beaconDuty.Type == types.BNRoleAttester {
			validatorPK := types.ValidatorPK(beaconDuty.PubKey)
			if c.HighestAttestingSlotMap[validatorPK] < beaconDuty.Slot {
				c.HighestAttestingSlotMap[validatorPK] = beaconDuty.Slot
			}
		}
	}
}
