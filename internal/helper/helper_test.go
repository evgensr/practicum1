package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHash(t *testing.T) {

	// var hash string
	testCases := []struct {
		name    string
		text    string
		result  string
		isValid bool
	}{
		{
			name:    "valid",
			text:    "1",
			result:  "c4ca4238a0b923820dcc509a6f75849b",
			isValid: true,
		},
		{
			name:    "invalid",
			text:    "123",
			result:  "00000000000000000000000000000000",
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hash := GetHash(tc.text)
			if tc.isValid {
				assert.Equal(t, tc.result, hash)
			} else {
				assert.NotEqual(t, tc.result, hash)
			}
			assert.Len(t, hash, 32)
		})
	}
}

func TestGenerateRandom(t *testing.T) {

	testCases := []struct {
		name    string
		size    int
		result  []byte
		isValid bool
	}{
		{
			name:    "valid",
			size:    5,
			result:  []byte("12345"),
			isValid: true,
		},
		{
			name:    "invalid",
			size:    10,
			result:  []byte("12"),
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GenerateRandom(tc.size)
			assert.NoError(t, err)
			if tc.isValid {
				assert.Len(t, tc.result, len(result))
			} else {
				assert.Less(t, tc.result, result)
			}
		})
	}
}

func BenchmarkGetHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetHash("text")
	}
}

func BenchmarkGetShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetShort("text")
	}
}

func BenchmarkGeneratorUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratorUUID()
	}
}

func BenchmarkGenerateRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandom(10)
	}
}

func BenchmarkAddSlash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddSlash("test")
	}
}
