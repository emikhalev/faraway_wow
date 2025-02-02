package pow

import (
	"fmt"
	"strings"
	"testing"
)

const (
	tokenSize  = 32
	difficulty = 20
)

// Тестируем функцию calcZeros
func Test_calcZeros(t *testing.T) {
	tests := []struct {
		name  string
		hash  [32]byte
		zeros int
	}{
		{
			name:  "No leading zeros",
			hash:  pseudoHash(0b11001010),
			zeros: 0,
		},
		{
			name:  "One leading zero",
			hash:  pseudoHash(0b01100101),
			zeros: 1,
		},
		{
			name:  "Multiple leading zeros",
			hash:  pseudoHash(0b00000000, 0b00000001),
			zeros: 15,
		},
		{
			name:  "All zeros",
			hash:  [32]byte{},
			zeros: 256,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcZeros(tt.hash)
			if got != tt.zeros {
				t.Errorf("calcZeros() = %v, want %v (%s)", got, tt.zeros, toBinaryString(tt.hash))
			}
		})
	}
}

func TestVerify(t *testing.T) {
	// Пример данных для теста
	challenge := []byte("test-challenge")

	// Пример успешных решений (реально найденные nonce для сложности)
	tests := []struct {
		challenge  []byte
		difficulty int
		nonce      uint64
		expected   bool
	}{
		{challenge, 16, 51216, true},
		{challenge, 8, 273, true},
		{challenge, 16, 43, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("difficulty=%d,nonce=%d", tt.difficulty, tt.nonce), func(t *testing.T) {
			result := verify(tt.challenge, tt.nonce, tt.difficulty)
			if result != tt.expected {
				t.Errorf("verify() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	c, err := challenge(32)
	if err != nil {
		t.Errorf("expecting no error, got %s", err)
	}

	nonce := solve(c, difficulty)
	if !verify(c, nonce, difficulty) {
		t.Errorf("not solved")
	}
}

func toBinaryString(hash [32]byte) string {
	var builder strings.Builder
	for _, b := range hash {
		builder.WriteString(fmt.Sprintf("%08b", b))
	}
	return builder.String()
}

func pseudoHash(b ...byte) [32]byte {
	hash := [32]byte{}
	for i := 0; i < len(b) && i < 32; i++ {
		hash[i] = b[i]
	}
	return hash
}
