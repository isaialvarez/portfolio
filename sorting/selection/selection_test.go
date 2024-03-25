package selection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		testName string
		input    []int
		expected []int
	}{
		{
			testName: "should sort",
			input:    []int{64, 25, 12, 22, 11},
			expected: []int{11, 12, 22, 25, 64},
		},
		{
			testName: "should sort ordered value",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tCase := range testCases {
		actual := Sort(tCase.input)
		assert.Equal(t, tCase.expected, actual)
	}
}
