package utils

import (
	"testing"

	"gotest.tools/assert"
)

func TestFormatFloat(t *testing.T) {
	tests := []struct {
		f        float64
		digits   int
		expected float64
	}{
		{
			f:        0.000012345,
			digits:   4,
			expected: 0.00001235,
		},
		{
			f:        1.000012345,
			digits:   4,
			expected: 1,
		},
		{
			f:        0.12345,
			digits:   4,
			expected: 0.1235,
		},
		{
			f:        1.12345,
			digits:   4,
			expected: 1.1235,
		},
		{
			f:        1.000012345,
			digits:   4,
			expected: 1,
		},
		{
			f:        0.0000123456789,
			digits:   7,
			expected: 0.00001234568,
		},
		{
			f:        0.000012345,
			digits:   2,
			expected: 0.000012,
		},
		{
			f:        0.012345,
			digits:   2,
			expected: 0.012,
		},
		{
			f:        0.013545,
			digits:   2,
			expected: 0.014,
		},
	}

	for _, ts := range tests {
		ff := FormatFloat(ts.f, ts.digits)
		assert.Equal(t, ff, ts.expected)
	}
}
