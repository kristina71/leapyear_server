package year

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYear(t *testing.T) {
	testCases := []struct {
		year   uint64
		expect bool
	}{
		{
			year:   1700,
			expect: false,
		},
		{
			year:   1800,
			expect: false,
		},
		{
			year:   1600,
			expect: true,
		},
	}
	for _, ts := range testCases {
		t.Run(strconv.FormatUint(uint64(ts.year), 10), func(t *testing.T) {
			require.Equal(t, ts.expect, IsLeapYear(ts.year))
		})
	}
}
