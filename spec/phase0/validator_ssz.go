// Code generated by fastssz. DO NOT EDIT.
// Hash: fa97a18404df62375826ed5ee7ded00832a9164171537b8327e9357589e54c7a
// Version: 0.1.3
package phase0

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Validator object
func (v *Validator) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the Validator object to a target array
func (v *Validator) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'PublicKey'
	dst = append(dst, v.PublicKey[:]...)

	// Field (1) 'WithdrawalCredentials'
	if size := len(v.WithdrawalCredentials); size != 32 {
		err = ssz.ErrBytesLengthFn("Validator.WithdrawalCredentials", size, 32)
		return
	}
	dst = append(dst, v.WithdrawalCredentials...)

	// Field (2) 'EffectiveBalance'
	dst = ssz.MarshalUint64(dst, uint64(v.EffectiveBalance))

	// Field (3) 'Slashed'
	dst = ssz.MarshalBool(dst, v.Slashed)

	// Field (4) 'ActivationEligibilityEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.ActivationEligibilityEpoch))

	// Field (5) 'ActivationEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.ActivationEpoch))

	// Field (6) 'ExitEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.ExitEpoch))

	// Field (7) 'WithdrawableEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.WithdrawableEpoch))

	return
}

// UnmarshalSSZ ssz unmarshals the Validator object
func (v *Validator) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 121 {
		return ssz.ErrSize
	}

	// Field (0) 'PublicKey'
	copy(v.PublicKey[:], buf[0:48])

	// Field (1) 'WithdrawalCredentials'
	if cap(v.WithdrawalCredentials) == 0 {
		v.WithdrawalCredentials = make([]byte, 0, len(buf[48:80]))
	}
	v.WithdrawalCredentials = append(v.WithdrawalCredentials, buf[48:80]...)

	// Field (2) 'EffectiveBalance'
	v.EffectiveBalance = Gwei(ssz.UnmarshallUint64(buf[80:88]))

	// Field (3) 'Slashed'
	v.Slashed = ssz.UnmarshalBool(buf[88:89])

	// Field (4) 'ActivationEligibilityEpoch'
	v.ActivationEligibilityEpoch = Epoch(ssz.UnmarshallUint64(buf[89:97]))

	// Field (5) 'ActivationEpoch'
	v.ActivationEpoch = Epoch(ssz.UnmarshallUint64(buf[97:105]))

	// Field (6) 'ExitEpoch'
	v.ExitEpoch = Epoch(ssz.UnmarshallUint64(buf[105:113]))

	// Field (7) 'WithdrawableEpoch'
	v.WithdrawableEpoch = Epoch(ssz.UnmarshallUint64(buf[113:121]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Validator object
func (v *Validator) SizeSSZ() (size int) {
	size = 121
	return
}

// HashTreeRoot ssz hashes the Validator object
func (v *Validator) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the Validator object with a hasher
func (v *Validator) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'PublicKey'
	hh.PutBytes(v.PublicKey[:])

	// Field (1) 'WithdrawalCredentials'
	if size := len(v.WithdrawalCredentials); size != 32 {
		err = ssz.ErrBytesLengthFn("Validator.WithdrawalCredentials", size, 32)
		return
	}
	hh.PutBytes(v.WithdrawalCredentials)

	// Field (2) 'EffectiveBalance'
	hh.PutUint64(uint64(v.EffectiveBalance))

	// Field (3) 'Slashed'
	hh.PutBool(v.Slashed)

	// Field (4) 'ActivationEligibilityEpoch'
	hh.PutUint64(uint64(v.ActivationEligibilityEpoch))

	// Field (5) 'ActivationEpoch'
	hh.PutUint64(uint64(v.ActivationEpoch))

	// Field (6) 'ExitEpoch'
	hh.PutUint64(uint64(v.ExitEpoch))

	// Field (7) 'WithdrawableEpoch'
	hh.PutUint64(uint64(v.WithdrawableEpoch))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Validator object
func (v *Validator) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(v)
}