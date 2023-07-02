package aesmthash

import (
	"crypto/aes"
	"encoding/binary"
	"math/bits"
)

const Size = aes.BlockSize

func Hash(data []byte) [Size]byte {
	var pad uint
	if len(data)%Size == 0 {
		pad = 0
	} else {
		pad = Size - uint(len(data))%Size
	}

	buffer := append(data[:len(data):len(data)], make([]byte, Size+pad)...)

	binary.LittleEndian.PutUint64(buffer[len(buffer)-Size:], uint64(len(data)))

	return greedyHash(buffer)
}

var zeros [Size]byte

// greedyHash's input length must be a multiple of [Size], it must not be empty either.
func greedyHash(data []byte) (r [Size]byte) {
	if len(data) == Size {
		c, err := aes.NewCipher(data)
		if err != nil {
			panic(err)
		}
		c.Encrypt(r[:], zeros[:])
		return
	}

	blocksize := uint(len(data)) / Size

	// log2 rounded down
	hashSize := uint(1) << (uint(bits.Len(blocksize)) - 1)
	if blocksize == hashSize {
		hashSize /= 2
	}

	lhs := greedyHash(data[:hashSize*Size])
	rhs := greedyHash(data[hashSize*Size:])
	c, err := aes.NewCipher(rhs[:])
	if err != nil {
		panic(err)
	}
	c.Encrypt(r[:], lhs[:])
	return
}
