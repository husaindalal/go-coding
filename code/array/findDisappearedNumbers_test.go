package array

import (
	"reflect"
	"sort"
	"testing"
)

// sortInts is a helper function to sort an integer slice for comparison
func sortInts(nums []int) []int {
	sort.Ints(nums)
	return nums
}

// Test cases for findDisappearedNumbers functions
func TestFindDisappearedNumbers(t *testing.T) {
	// Create a large array with numbers 1-100, but with many duplicates for the large test case
	largeN := 100
	largeInput := make([]int, largeN)
	for i := 0; i < largeN; i++ {
		// Use only numbers 1-50, causing 51-100 to be missing
		largeInput[i] = (i % 50) + 1
	}
	
	// Expected: all numbers from 51 to 100
	largeExpected := make([]int, 50)
	for i := 0; i < 50; i++ {
		largeExpected[i] = i + 51
	}

	testCases := []struct {
		name     string
		input    []int
		expected []int
		desc     string
	}{
		{
			name:     "Example 1 - Normal case with some missing numbers",
			input:    []int{4, 3, 2, 7, 8, 2, 3, 1},
			expected: []int{5, 6},
			desc:     "Happy path: Input array with some missing numbers should return those missing numbers",
		},
		{
			name:     "Example 2 - Normal case with repeated numbers",
			input:    []int{1, 1},
			expected: []int{2},
			desc:     "Happy path: Input array with repeated numbers should identify missing numbers correctly",
		},
		{
			name:     "Edge case - Empty array",
			input:    []int{},
			expected: []int{},
			desc:     "Edge case: With no numbers, there are no missing numbers in the range [1, n] because n=0",
		},
		{
			name:     "Edge case - No missing numbers",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{},
			desc:     "Edge case: When all numbers in range [1, n] are present, should return empty array",
		},
		{
			name:     "Edge case - Only one number repeated multiple times",
			input:    []int{1, 1, 1, 1, 1},
			expected: []int{2, 3, 4, 5},
			desc:     "Edge case: When only one number is repeated multiple times, should identify all other missing numbers",
		},
		{
			name:     "Edge case - Large array with many duplicates",
			input:    largeInput,
			expected: largeExpected,
			desc:     "Edge case: Tests the algorithm's performance with larger inputs where half the range is missing",
		},
	}

	// Test the brute force solution
	for _, tc := range testCases {
		t.Run("BruteForce: "+tc.name, func(t *testing.T) {
			t.Logf("Description: %s", tc.desc)
			
			inputCopy := make([]int, len(tc.input))
			copy(inputCopy, tc.input)

			result := findDisappearedNumbersBruteForce(inputCopy)
			if !reflect.DeepEqual(sortInts(result), sortInts(tc.expected)) {
				t.Errorf("Expected %v but got %v for input %v", tc.expected, result, tc.input)
			}
		})
	}

	// Test the optimized solution
	for _, tc := range testCases {
		t.Run("Optimized: "+tc.name, func(t *testing.T) {
			t.Logf("Description: %s", tc.desc)
			
			inputCopy := make([]int, len(tc.input))
			copy(inputCopy, tc.input)

			result := findDisappearedNumbers(inputCopy)
			if !reflect.DeepEqual(sortInts(result), sortInts(tc.expected)) {
				t.Errorf("Expected %v but got %v for input %v", tc.expected, result, tc.input)
			}
		})
	}

	// Test the follow-up solution
	for _, tc := range testCases {
		t.Run("FollowUp: "+tc.name, func(t *testing.T) {
			t.Logf("Description: %s", tc.desc)
			
			inputCopy := make([]int, len(tc.input))
			copy(inputCopy, tc.input)

			result := findDisappearedNumbersFollowUp(inputCopy)
			if !reflect.DeepEqual(sortInts(result), sortInts(tc.expected)) {
				t.Errorf("Expected %v but got %v for input %v", tc.expected, result, tc.input)
			}
		})
	}
}

// Run the tests with: go test -v ./code/array/...
