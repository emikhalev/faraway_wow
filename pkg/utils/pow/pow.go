package pow

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
)

func Challenge(size int64) ([]byte, error) {
	token := make([]byte, size)
	_, err := rand.Read(token)
	return token, err
}

func Solve(challenge []byte, difficulty int) uint64 {
	var nonce uint64
	for {
		data := append(challenge, make([]byte, 8)...)
		binary.BigEndian.PutUint64(data[len(challenge):], nonce)

		hash := sha256.Sum256(data)
		if calcZeros(hash) >= difficulty {
			return nonce
		}
		nonce++
	}
}

func Verify(challenge []byte, nonce uint64, difficulty int) bool {
	data := append(challenge, make([]byte, 8)...)
	binary.BigEndian.PutUint64(data[len(challenge):], nonce)

	hash := sha256.Sum256(data)
	return calcZeros(hash) >= difficulty
}

func calcZeros(hash [32]byte) int {
	byteIndex := 0
	bitIndex := 7
	zerosCount := 0

	for {
		if (hash[byteIndex]>>bitIndex)&1 == 0 {
			zerosCount++
		} else {
			break
		}

		if bitIndex == 0 {
			bitIndex = 7
			byteIndex++
		} else {
			bitIndex--
		}

		if byteIndex >= len(hash) {
			break
		}
	}

	return zerosCount
}
