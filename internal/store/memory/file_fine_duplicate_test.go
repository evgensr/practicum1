package memory

import (
	"testing"

	"github.com/evgensr/practicum1/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestFineDuplicate(t *testing.T) {
	var items Box

	line := []Line{

		{
			User:  helper.GeneratorUUID(),
			URL:   "http://url1",
			Short: helper.GetHash("http://url1"),
		},
		{
			User:  helper.GeneratorUUID(),
			URL:   "http://url2",
			Short: helper.GetHash("http://url2"),
		},
		{
			User:  helper.GeneratorUUID(),
			URL:   "http://url3",
			Short: helper.GetHash("http://url3"),
		},
	}

	items = Box{
		Items:           line,
		fileStoragePath: "string",
	}

	testCases := []struct {
		name    string
		payload string
		isValid bool
	}{

		{
			name:    "valid",
			payload: helper.GetHash("http://url0"),
			isValid: false,
		},
		{
			name:    "Invalid",
			payload: helper.GetHash("http://url1"),
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.isValid, fineDuplicate(&items, tc.payload))
		})
	}
}
