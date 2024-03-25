package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymetricDifference(t *testing.T) {
	testCases := []struct {
		testName string
		inputs   [][]int
		expected []int
	}{
		{
			testName: "Test 2 items",
			inputs: [][]int{
				{1, 2, 3},
				{5, 2, 1, 4},
			},
			expected: []int{3, 4, 5},
		},
		{
			testName: "Test 2 items with repeated numbers",
			inputs: [][]int{
				{1, 2, 3, 3},
				{5, 2, 1, 4},
			},
			expected: []int{3, 4, 5},
		},
		{
			testName: "Test 3 items",
			inputs: [][]int{
				{1, 2, 5},
				{2, 3, 5},
				{3, 4, 5},
			},
			expected: []int{1, 4, 5},
		},
		{
			testName: "Test 3 items with duplicates",
			inputs: [][]int{
				{1, 1, 2, 5},
				{2, 2, 3, 5},
				{3, 4, 5, 5},
			},
			expected: []int{1, 4, 5},
		},
		{
			testName: "Test 4 items",
			inputs: [][]int{
				{3, 3, 3, 2, 5},
				{2, 1, 5, 7},
				{3, 4, 6, 6},
				{1, 2, 3},
			},
			expected: []int{2, 3, 4, 6, 7},
		},
		{
			testName: "Test 6 items",
			inputs: [][]int{
				{3, 3, 3, 2, 5},
				{2, 1, 5, 7},
				{3, 4, 6, 6},
				{1, 2, 3},
				{5, 3, 9, 8},
				{1},
			},
			expected: []int{1, 2, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tCase := range testCases {
		actual := SymetricDifference(tCase.inputs...)
		assert.Equal(t, tCase.expected, actual, tCase.testName)
	}
}

func TestPairwise(t *testing.T) {
	testCases := []struct {
		testName string
		input1   []int
		input2   int
		expected int
	}{
		{
			testName: "Test 1",
			input1:   []int{1, 1, 2},
			input2:   3,
			expected: 2,
		},
		{
			testName: "Test 2",
			input1:   []int{7, 9, 11, 13, 15},
			input2:   20,
			expected: 6,
		},
		{
			testName: "Test 3",
			input1:   []int{1, 4, 2, 3, 0, 5},
			input2:   7,
			expected: 11,
		},
		{
			testName: "Test 4",
			input1:   []int{1, 3, 2, 4},
			input2:   4,
			expected: 1,
		},
		{
			testName: "Test 5",
			input1:   []int{1, 1, 1},
			input2:   2,
			expected: 1,
		},
		{
			testName: "Test 6",
			input1:   []int{0, 0, 0, 0, 1, 1},
			input2:   1,
			expected: 10,
		},
		{
			testName: "Test 7",
			input1:   []int{},
			input2:   100,
			expected: 0,
		},
	}

	for _, tCase := range testCases {
		actual := Pairwise(tCase.input1, tCase.input2)
		assert.Equal(t, tCase.expected, actual, tCase.testName)
	}
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		testName string
		input1   int
		input2   int
		expected int
	}{
		{
			testName: "should sort",
			input1:   3,
			input2:   2,
			expected: 6,
		},
	}
	for _, tCase := range testCases {
		actual := CallupdateInventory(tCase.input1, tCase.input2)
		assert.Equal(t, tCase.expected, actual)
	}
}
