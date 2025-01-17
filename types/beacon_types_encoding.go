// Code generated by fastssz. DO NOT EDIT.
// Hash: 5e34e9f71b47bb74105d3a7de6974dc47fcbf337aadbffb4859a3262b9c29ba7
// Version: 0.1.2
package types

import (
	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Duty object
func (d *Duty) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the Duty object to a target array
func (d *Duty) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(108)

	// Field (0) 'Type'
	dst = ssz.MarshalUint64(dst, uint64(d.Type))

	// Field (1) 'PubKey'
	dst = append(dst, d.PubKey[:]...)

	// Field (2) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(d.Slot))

	// Field (3) 'ValidatorIndex'
	dst = ssz.MarshalUint64(dst, uint64(d.ValidatorIndex))

	// Field (4) 'CommitteeIndex'
	dst = ssz.MarshalUint64(dst, uint64(d.CommitteeIndex))

	// Field (5) 'CommitteeLength'
	dst = ssz.MarshalUint64(dst, d.CommitteeLength)

	// Field (6) 'CommitteesAtSlot'
	dst = ssz.MarshalUint64(dst, d.CommitteesAtSlot)

	// Field (7) 'ValidatorCommitteeIndex'
	dst = ssz.MarshalUint64(dst, d.ValidatorCommitteeIndex)

	// Offset (8) 'ValidatorSyncCommitteeIndices'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(d.ValidatorSyncCommitteeIndices) * 8

	// Field (8) 'ValidatorSyncCommitteeIndices'
	if size := len(d.ValidatorSyncCommitteeIndices); size > 13 {
		err = ssz.ErrListTooBigFn("Duty.ValidatorSyncCommitteeIndices", size, 13)
		return
	}
	for ii := 0; ii < len(d.ValidatorSyncCommitteeIndices); ii++ {
		dst = ssz.MarshalUint64(dst, d.ValidatorSyncCommitteeIndices[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Duty object
func (d *Duty) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 108 {
		return ssz.ErrSize
	}

	tail := buf
	var o8 uint64

	// Field (0) 'Type'
	d.Type = BeaconRole(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'PubKey'
	copy(d.PubKey[:], buf[8:56])

	// Field (2) 'Slot'
	d.Slot = spec.Slot(ssz.UnmarshallUint64(buf[56:64]))

	// Field (3) 'ValidatorIndex'
	d.ValidatorIndex = spec.ValidatorIndex(ssz.UnmarshallUint64(buf[64:72]))

	// Field (4) 'CommitteeIndex'
	d.CommitteeIndex = spec.CommitteeIndex(ssz.UnmarshallUint64(buf[72:80]))

	// Field (5) 'CommitteeLength'
	d.CommitteeLength = ssz.UnmarshallUint64(buf[80:88])

	// Field (6) 'CommitteesAtSlot'
	d.CommitteesAtSlot = ssz.UnmarshallUint64(buf[88:96])

	// Field (7) 'ValidatorCommitteeIndex'
	d.ValidatorCommitteeIndex = ssz.UnmarshallUint64(buf[96:104])

	// Offset (8) 'ValidatorSyncCommitteeIndices'
	if o8 = ssz.ReadOffset(buf[104:108]); o8 > size {
		return ssz.ErrOffset
	}

	if o8 < 108 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (8) 'ValidatorSyncCommitteeIndices'
	{
		buf = tail[o8:]
		num, err := ssz.DivideInt2(len(buf), 8, 13)
		if err != nil {
			return err
		}
		d.ValidatorSyncCommitteeIndices = ssz.ExtendUint64(d.ValidatorSyncCommitteeIndices, num)
		for ii := 0; ii < num; ii++ {
			d.ValidatorSyncCommitteeIndices[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Duty object
func (d *Duty) SizeSSZ() (size int) {
	size = 108

	// Field (8) 'ValidatorSyncCommitteeIndices'
	size += len(d.ValidatorSyncCommitteeIndices) * 8

	return
}

// HashTreeRoot ssz hashes the Duty object
func (d *Duty) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the Duty object with a hasher
func (d *Duty) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Type'
	hh.PutUint64(uint64(d.Type))

	// Field (1) 'PubKey'
	hh.PutBytes(d.PubKey[:])

	// Field (2) 'Slot'
	hh.PutUint64(uint64(d.Slot))

	// Field (3) 'ValidatorIndex'
	hh.PutUint64(uint64(d.ValidatorIndex))

	// Field (4) 'CommitteeIndex'
	hh.PutUint64(uint64(d.CommitteeIndex))

	// Field (5) 'CommitteeLength'
	hh.PutUint64(d.CommitteeLength)

	// Field (6) 'CommitteesAtSlot'
	hh.PutUint64(d.CommitteesAtSlot)

	// Field (7) 'ValidatorCommitteeIndex'
	hh.PutUint64(d.ValidatorCommitteeIndex)

	// Field (8) 'ValidatorSyncCommitteeIndices'
	{
		if size := len(d.ValidatorSyncCommitteeIndices); size > 13 {
			err = ssz.ErrListTooBigFn("Duty.ValidatorSyncCommitteeIndices", size, 13)
			return
		}
		subIndx := hh.Index()
		for _, i := range d.ValidatorSyncCommitteeIndices {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(d.ValidatorSyncCommitteeIndices))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(13, numItems, 8))
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Duty object
func (d *Duty) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(d)
}
