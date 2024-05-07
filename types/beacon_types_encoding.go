// Code generated by fastssz. DO NOT EDIT.
// Hash: d24b700b53f81f415e6c695470210775a880e70469a64411980f3e36dd703ea0
// Version: 0.1.3
package types

import (
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconDuty object
func (b *BeaconDuty) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconDuty object to a target array
func (b *BeaconDuty) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(109)

	// Field (0) 'Type'
	dst = ssz.MarshalUint64(dst, uint64(b.Type))

	// Field (1) 'PubKey'
	dst = append(dst, b.PubKey[:]...)

	// Field (2) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Field (3) 'ValidatorIndex'
	dst = ssz.MarshalUint64(dst, uint64(b.ValidatorIndex))

	// Field (4) 'CommitteeIndex'
	dst = ssz.MarshalUint64(dst, uint64(b.CommitteeIndex))

	// Field (5) 'CommitteeLength'
	dst = ssz.MarshalUint64(dst, b.CommitteeLength)

	// Field (6) 'CommitteesAtSlot'
	dst = ssz.MarshalUint64(dst, b.CommitteesAtSlot)

	// Field (7) 'ValidatorCommitteeIndex'
	dst = ssz.MarshalUint64(dst, b.ValidatorCommitteeIndex)

	// Offset (8) 'ValidatorSyncCommitteeIndices'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.ValidatorSyncCommitteeIndices) * 8

	// Field (9) 'IsStopped'
	dst = ssz.MarshalBool(dst, b.IsStopped)

	// Field (8) 'ValidatorSyncCommitteeIndices'
	if size := len(b.ValidatorSyncCommitteeIndices); size > 13 {
		err = ssz.ErrListTooBigFn("BeaconDuty.ValidatorSyncCommitteeIndices", size, 13)
		return
	}
	for ii := 0; ii < len(b.ValidatorSyncCommitteeIndices); ii++ {
		dst = ssz.MarshalUint64(dst, b.ValidatorSyncCommitteeIndices[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconDuty object
func (b *BeaconDuty) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 109 {
		return ssz.ErrSize
	}

	tail := buf
	var o8 uint64

	// Field (0) 'Type'
	b.Type = BeaconRole(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'PubKey'
	copy(b.PubKey[:], buf[8:56])

	// Field (2) 'Slot'
	b.Slot = spec.Slot(ssz.UnmarshallUint64(buf[56:64]))

	// Field (3) 'ValidatorIndex'
	b.ValidatorIndex = spec.ValidatorIndex(ssz.UnmarshallUint64(buf[64:72]))

	// Field (4) 'CommitteeIndex'
	b.CommitteeIndex = spec.CommitteeIndex(ssz.UnmarshallUint64(buf[72:80]))

	// Field (5) 'CommitteeLength'
	b.CommitteeLength = ssz.UnmarshallUint64(buf[80:88])

	// Field (6) 'CommitteesAtSlot'
	b.CommitteesAtSlot = ssz.UnmarshallUint64(buf[88:96])

	// Field (7) 'ValidatorCommitteeIndex'
	b.ValidatorCommitteeIndex = ssz.UnmarshallUint64(buf[96:104])

	// Offset (8) 'ValidatorSyncCommitteeIndices'
	if o8 = ssz.ReadOffset(buf[104:108]); o8 > size {
		return ssz.ErrOffset
	}

	if o8 < 109 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (9) 'IsStopped'
	b.IsStopped = ssz.UnmarshalBool(buf[108:109])

	// Field (8) 'ValidatorSyncCommitteeIndices'
	{
		buf = tail[o8:]
		num, err := ssz.DivideInt2(len(buf), 8, 13)
		if err != nil {
			return err
		}
		b.ValidatorSyncCommitteeIndices = ssz.ExtendUint64(b.ValidatorSyncCommitteeIndices, num)
		for ii := 0; ii < num; ii++ {
			b.ValidatorSyncCommitteeIndices[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconDuty object
func (b *BeaconDuty) SizeSSZ() (size int) {
	size = 109

	// Field (8) 'ValidatorSyncCommitteeIndices'
	size += len(b.ValidatorSyncCommitteeIndices) * 8

	return
}

// HashTreeRoot ssz hashes the BeaconDuty object
func (b *BeaconDuty) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconDuty object with a hasher
func (b *BeaconDuty) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Type'
	hh.PutUint64(uint64(b.Type))

	// Field (1) 'PubKey'
	hh.PutBytes(b.PubKey[:])

	// Field (2) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (3) 'ValidatorIndex'
	hh.PutUint64(uint64(b.ValidatorIndex))

	// Field (4) 'CommitteeIndex'
	hh.PutUint64(uint64(b.CommitteeIndex))

	// Field (5) 'CommitteeLength'
	hh.PutUint64(b.CommitteeLength)

	// Field (6) 'CommitteesAtSlot'
	hh.PutUint64(b.CommitteesAtSlot)

	// Field (7) 'ValidatorCommitteeIndex'
	hh.PutUint64(b.ValidatorCommitteeIndex)

	// Field (8) 'ValidatorSyncCommitteeIndices'
	{
		if size := len(b.ValidatorSyncCommitteeIndices); size > 13 {
			err = ssz.ErrListTooBigFn("BeaconDuty.ValidatorSyncCommitteeIndices", size, 13)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.ValidatorSyncCommitteeIndices {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(b.ValidatorSyncCommitteeIndices))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(13, numItems, 8))
	}

	// Field (9) 'IsStopped'
	hh.PutBool(b.IsStopped)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconDuty object
func (b *BeaconDuty) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}

// MarshalSSZ ssz marshals the BeaconVote object
func (b *BeaconVote) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconVote object to a target array
func (b *BeaconVote) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'BlockRoot'
	dst = append(dst, b.BlockRoot[:]...)

	// Field (1) 'Source'
	if b.Source == nil {
		b.Source = new(spec.Checkpoint)
	}
	if dst, err = b.Source.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Target'
	if b.Target == nil {
		b.Target = new(spec.Checkpoint)
	}
	if dst, err = b.Target.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconVote object
func (b *BeaconVote) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 112 {
		return ssz.ErrSize
	}

	// Field (0) 'BlockRoot'
	copy(b.BlockRoot[:], buf[0:32])

	// Field (1) 'Source'
	if b.Source == nil {
		b.Source = new(spec.Checkpoint)
	}
	if err = b.Source.UnmarshalSSZ(buf[32:72]); err != nil {
		return err
	}

	// Field (2) 'Target'
	if b.Target == nil {
		b.Target = new(spec.Checkpoint)
	}
	if err = b.Target.UnmarshalSSZ(buf[72:112]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconVote object
func (b *BeaconVote) SizeSSZ() (size int) {
	size = 112
	return
}

// HashTreeRoot ssz hashes the BeaconVote object
func (b *BeaconVote) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconVote object with a hasher
func (b *BeaconVote) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'BlockRoot'
	hh.PutBytes(b.BlockRoot[:])

	// Field (1) 'Source'
	if b.Source == nil {
		b.Source = new(spec.Checkpoint)
	}
	if err = b.Source.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Target'
	if b.Target == nil {
		b.Target = new(spec.Checkpoint)
	}
	if err = b.Target.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconVote object
func (b *BeaconVote) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
