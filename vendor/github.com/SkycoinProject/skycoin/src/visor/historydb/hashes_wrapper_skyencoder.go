// Code generated by github.com/SkycoinProject/skyencoder. DO NOT EDIT.

package historydb

import (
	"errors"
	"math"

	"github.com/SkycoinProject/skycoin/src/cipher"
	"github.com/SkycoinProject/skycoin/src/cipher/encoder"
)

// encodeSizeHashesWrapper computes the size of an encoded object of type hashesWrapper
func encodeSizeHashesWrapper(obj *hashesWrapper) uint64 {
	i0 := uint64(0)

	// obj.Hashes
	i0 += 4
	{
		i1 := uint64(0)

		// x1
		i1 += 32

		i0 += uint64(len(obj.Hashes)) * i1
	}

	return i0
}

// encodeHashesWrapper encodes an object of type hashesWrapper to a buffer allocated to the exact size
// required to encode the object.
func encodeHashesWrapper(obj *hashesWrapper) ([]byte, error) {
	n := encodeSizeHashesWrapper(obj)
	buf := make([]byte, n)

	if err := encodeHashesWrapperToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeHashesWrapperToBuffer encodes an object of type hashesWrapper to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeHashesWrapperToBuffer(buf []byte, obj *hashesWrapper) error {
	if uint64(len(buf)) < encodeSizeHashesWrapper(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Hashes length check
	if uint64(len(obj.Hashes)) > math.MaxUint32 {
		return errors.New("obj.Hashes length exceeds math.MaxUint32")
	}

	// obj.Hashes length
	e.Uint32(uint32(len(obj.Hashes)))

	// obj.Hashes
	for _, x := range obj.Hashes {

		// x
		e.CopyBytes(x[:])

	}

	return nil
}

// decodeHashesWrapper decodes an object of type hashesWrapper from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeHashesWrapper(buf []byte, obj *hashesWrapper) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Hashes

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length != 0 {
			obj.Hashes = make([]cipher.SHA256, length)

			for z1 := range obj.Hashes {
				{
					// obj.Hashes[z1]
					if len(d.Buffer) < len(obj.Hashes[z1]) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.Hashes[z1][:], d.Buffer[:len(obj.Hashes[z1])])
					d.Buffer = d.Buffer[len(obj.Hashes[z1]):]
				}

			}
		}
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeHashesWrapperExact decodes an object of type hashesWrapper from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeHashesWrapperExact(buf []byte, obj *hashesWrapper) error {
	if n, err := decodeHashesWrapper(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
