package ssv

import (
	"bytes"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"

	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/types"
)

func dutyValueCheck(
	duty *types.BeaconDuty,
	network types.BeaconNetwork,
	expectedType types.BeaconRole,
	validatorPK types.ValidatorPK,
	validatorIndex phase0.ValidatorIndex,
) error {
	if network.EstimatedEpochAtSlot(duty.Slot) > network.EstimatedCurrentEpoch()+1 {
		return errors.New("duty epoch is into far future")
	}

	if expectedType != duty.Type {
		return errors.New("wrong beacon role type")
	}

	if !bytes.Equal(validatorPK[:], duty.PubKey[:]) {
		return errors.New("wrong validator pk")
	}

	if validatorIndex != duty.ValidatorIndex {
		return errors.New("wrong validator index")
	}

	return nil
}

// TODO add tests
func BeaconVoteValueCheckF(
	signer types.BeaconSigner,
	duty types.Duty,
	sharePublicKey []byte,
	estimatedCurrentEpoch phase0.Epoch,
) qbft.ProposedValueCheckF {
	return func(data []byte) error {
		bv := types.BeaconVote{}
		if err := bv.Decode(data); err != nil {
			return errors.Wrap(err, "failed decoding beacon vote")
		}

		if bv.Target.Epoch > estimatedCurrentEpoch+1 {
			return errors.New("attestation data target epoch is into far future")
		}

		if bv.Source.Epoch >= bv.Target.Epoch {
			return errors.New("attestation data source > target")
		}

		attestationData := &phase0.AttestationData{
			Slot: duty.DutySlot(),
			// CommitteeIndex doesn't matter for slashing checks
			Index:           0,
			BeaconBlockRoot: bv.BlockRoot,
			Source:          nil,
			Target:          nil,
		}

		return signer.IsAttestationSlashable(sharePublicKey, attestationData)
	}
}

func ProposerValueCheckF(
	signer types.BeaconSigner,
	network types.BeaconNetwork,
	validatorPK types.ValidatorPK,
	validatorIndex phase0.ValidatorIndex,
	sharePublicKey []byte,
) qbft.ProposedValueCheckF {
	return func(data []byte) error {
		cd := &types.ConsensusData{}
		if err := cd.Decode(data); err != nil {
			return errors.Wrap(err, "failed decoding consensus data")
		}
		if err := cd.Validate(); err != nil {
			return errors.Wrap(err, "invalid value")
		}

		if err := dutyValueCheck(&cd.Duty, network, types.BNRoleProposer, validatorPK, validatorIndex); err != nil {
			return errors.Wrap(err, "duty invalid")
		}

		if blockData, _, err := cd.GetBlindedBlockData(); err == nil {
			slot, err := blockData.Slot()
			if err != nil {
				return errors.Wrap(err, "failed to get slot from blinded block data")
			}
			return signer.IsBeaconBlockSlashable(sharePublicKey, slot)
		}
		if blockData, _, err := cd.GetBlockData(); err == nil {
			slot, err := blockData.Slot()
			if err != nil {
				return errors.Wrap(err, "failed to get slot from block data")
			}
			return signer.IsBeaconBlockSlashable(sharePublicKey, slot)
		}

		return errors.New("no block data")
	}
}

func AggregatorValueCheckF(
	signer types.BeaconSigner,
	network types.BeaconNetwork,
	validatorPK types.ValidatorPK,
	validatorIndex phase0.ValidatorIndex,
) qbft.ProposedValueCheckF {
	return func(data []byte) error {
		cd := &types.ConsensusData{}
		if err := cd.Decode(data); err != nil {
			return errors.Wrap(err, "failed decoding consensus data")
		}
		if err := cd.Validate(); err != nil {
			return errors.Wrap(err, "invalid value")
		}

		if err := dutyValueCheck(&cd.Duty, network, types.BNRoleAggregator, validatorPK, validatorIndex); err != nil {
			return errors.Wrap(err, "duty invalid")
		}
		return nil
	}
}

func SyncCommitteeContributionValueCheckF(
	signer types.BeaconSigner,
	network types.BeaconNetwork,
	validatorPK types.ValidatorPK,
	validatorIndex phase0.ValidatorIndex,
) qbft.ProposedValueCheckF {
	return func(data []byte) error {
		cd := &types.ConsensusData{}
		if err := cd.Decode(data); err != nil {
			return errors.Wrap(err, "failed decoding consensus data")
		}
		if err := cd.Validate(); err != nil {
			return errors.Wrap(err, "invalid value")
		}

		if err := dutyValueCheck(&cd.Duty, network, types.BNRoleSyncCommitteeContribution, validatorPK, validatorIndex); err != nil {
			return errors.Wrap(err, "duty invalid")
		}

		//contributions, _ := cd.GetSyncCommitteeContributions()
		//
		//for _, c := range contributions {
		//	// TODO check we have selection proof for contribution
		//	// TODO check slot == duty slot
		//	// TODO check beacon block root somehow? maybe all beacon block roots should be equal?
		//
		//}
		return nil
	}
}
