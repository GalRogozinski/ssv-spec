package spectest

import (
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/committee"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/dutyexe"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/partialsigcontainer"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/runner"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/runner/consensus"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/runner/duties/newduty"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/runner/duties/proposer"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/runner/duties/synccommitteeaggregator"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/runner/postconsensus"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/runner/preconsensus"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/valcheck/valcheckattestations"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/valcheck/valcheckduty"
	"github.com/ssvlabs/ssv-spec/ssv/spectest/tests/valcheck/valcheckproposer"
)

var AllTests = []tests.TestF{
	runner.FullHappyFlow,

	postconsensus.TooManyRoots,
	postconsensus.TooFewRoots,
	postconsensus.UnorderedExpectedRoots,
	postconsensus.UnknownSigner,
	postconsensus.InconsistentBeaconSigner,
	postconsensus.PostFinish,
	postconsensus.NoRunningDuty,
	postconsensus.InvalidMessageSignature,
	postconsensus.InvalidBeaconSignatureInQuorum,
	postconsensus.DuplicateMsgDifferentRoots,
	postconsensus.DuplicateMsgDifferentRootsThenQuorum,
	postconsensus.DuplicateMsg,
	postconsensus.InvalidExpectedRoot,
	postconsensus.PreDecided,
	postconsensus.PostQuorum,
	postconsensus.InvalidMessage,
	postconsensus.InvalidOperatorSignature,
	postconsensus.InvalidMessageSlot,
	postconsensus.ValidMessage,
	postconsensus.ValidMessage7Operators,
	postconsensus.ValidMessage10Operators,
	postconsensus.ValidMessage13Operators,
	postconsensus.Quorum,
	postconsensus.Quorum7Operators,
	postconsensus.Quorum10Operators,
	postconsensus.Quorum13Operators,
	postconsensus.InvalidDecidedValue,
	postconsensus.InvalidThenQuorum,
	postconsensus.InvalidQuorumThenValidQuorum,
	postconsensus.InconsistentOperatorSigner,
	postconsensus.NilSSVMessage,
	postconsensus.InvalidValidatorIndex,
	postconsensus.PartialInvalidRootQuorumThenValidQuorum,
	postconsensus.PartialInvalidSigQuorumThenValidQuorum,
	postconsensus.MixedCommittees,

	newduty.ConsensusNotStarted,
	newduty.NotDecided,
	newduty.PostDecided,
	newduty.Finished,
	newduty.Valid,
	newduty.PostWrongDecided,
	newduty.PostInvalidDecided,
	newduty.PostFutureDecided,
	newduty.DuplicateDutyFinished,
	newduty.DuplicateDutyNotFinished,
	newduty.FirstHeight,

	committee.StartDuty,
	committee.StartNoDuty,
	committee.ValidBeaconVote,
	committee.WrongBeaconVote,
	committee.Decided,
	committee.HappyFlow,
	committee.PastMessageDutyNotFinished,
	committee.PastMessageDutyFinished,
	committee.PastMessageDutyDoesNotExist,
	committee.ProposalWithConsensusData,

	consensus.FutureDecidedNoInstance,
	consensus.FutureDecided,
	consensus.InvalidDecidedValue,
	consensus.FutureMessage,
	consensus.PastMessage,
	consensus.PostFinish,
	consensus.PostDecided,
	consensus.ValidDecided,
	consensus.ValidDecided7Operators,
	consensus.ValidDecided10Operators,
	consensus.ValidDecided13Operators,
	consensus.ValidMessage,
	consensus.InvalidSignature,

	synccommitteeaggregator.SomeAggregatorQuorum,
	synccommitteeaggregator.NoneAggregatorQuorum,
	synccommitteeaggregator.AllAggregatorQuorum,

	proposer.ProposeBlindedBlockDecidedRegular,
	proposer.ProposeRegularBlockDecidedBlinded,
	proposer.BlindedRunnerAcceptsNormalBlock,
	proposer.NormalProposerAcceptsBlindedBlock,

	// pre_consensus_justifications.PastSlot,
	// pre_consensus_justifications.InvalidData,
	// pre_consensus_justifications.FutureHeight,
	// pre_consensus_justifications.PastHeight,
	// pre_consensus_justifications.InvalidMsgType,
	// pre_consensus_justifications.WrongBeaconRole,
	// pre_consensus_justifications.InvalidConsensusData,
	// pre_consensus_justifications.InvalidSlot,
	// pre_consensus_justifications.UnknownSigner,
	// pre_consensus_justifications.InvalidJustificationSignature,
	// pre_consensus_justifications.DuplicateJustificationSigner,
	// pre_consensus_justifications.DuplicateRoots,
	// pre_consensus_justifications.InconsistentRootCount,
	// pre_consensus_justifications.InconsistentRoots,
	// pre_consensus_justifications.InvalidJustification,
	// pre_consensus_justifications.MissingQuorum,
	// pre_consensus_justifications.DecidedInstance,
	// pre_consensus_justifications.ExistingValidPreConsensus,
	// pre_consensus_justifications.Valid,
	// pre_consensus_justifications.Valid7Operators,
	// pre_consensus_justifications.Valid10Operators,
	// pre_consensus_justifications.Valid13Operators,
	// pre_consensus_justifications.ValidFirstHeight,
	// pre_consensus_justifications.ValidNoRunningDuty,
	// pre_consensus_justifications.ValidRoundChangeMsg,
	// pre_consensus_justifications.HappyFlow,

	preconsensus.NoRunningDuty,
	preconsensus.TooFewRoots,
	preconsensus.TooManyRoots,
	preconsensus.UnorderedExpectedRoots,
	preconsensus.InvalidSignedMessage,
	preconsensus.InvalidOperatorSignature,
	preconsensus.InvalidExpectedRoot,
	preconsensus.DuplicateMsg,
	preconsensus.DuplicateMsgDifferentRoots,
	preconsensus.PostFinish,
	preconsensus.PostDecided,
	preconsensus.PostQuorum,
	preconsensus.Quorum,
	preconsensus.Quorum7Operators,
	preconsensus.Quorum10Operators,
	preconsensus.Quorum13Operators,
	preconsensus.ValidMessage,
	preconsensus.InvalidMessageSlot,
	preconsensus.ValidMessage7Operators,
	preconsensus.ValidMessage10Operators,
	preconsensus.ValidMessage13Operators,
	preconsensus.InconsistentBeaconSigner,
	preconsensus.UnknownSigner,
	preconsensus.InvalidBeaconSignatureInQuorum,
	preconsensus.InvalidMessageSignature,
	preconsensus.InvalidThenQuorum,
	preconsensus.InvalidQuorumThenValidQuorum,
	preconsensus.InconsistentOperatorSigner,
	preconsensus.NilSSVMessage,

	valcheckduty.WrongValidatorIndex,
	valcheckduty.WrongValidatorPK,
	valcheckduty.WrongDutyType,
	valcheckduty.FarFutureDutySlot,

	valcheckattestations.Slashable,
	valcheckattestations.SourceHigherThanTarget,
	valcheckattestations.FarFutureTarget,
	valcheckattestations.BeaconVoteDataNil,
	valcheckattestations.Valid,
	valcheckattestations.MinoritySlashable,
	valcheckattestations.MajoritySlashable,

	valcheckproposer.BlindedBlock,

	dutyexe.WrongDutyRole,
	dutyexe.WrongDutyPubKey,
	partialsigcontainer.OneSignature,
	partialsigcontainer.Quorum,
	partialsigcontainer.Duplicate,
	partialsigcontainer.DuplicateQuorum,
	partialsigcontainer.Invalid,
}
