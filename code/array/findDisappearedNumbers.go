package array

// Brute force solution
// This solution uses a map to track the presence of each number in the input array.
// It then iterates through the range [1, n] to find the missing numbers.
func findDisappearedNumbersBruteForce(nums []int) []int {
	n := len(nums)
	result := []int{}
	numMap := make(map[int]bool)

	// Populate the map with the numbers present in the array
	for _, num := range nums {
		numMap[num] = true
	}

	// Check for numbers in the range [1, n] that are not in the map
	for i := 1; i <= n; i++ {
		if !numMap[i] {
			result = append(result, i)
		}
	}

	return result
}

// Optimized solution
// This solution marks the presence of each number by negating the value at the index corresponding to the number.
// It then iterates through the array to find the indices with positive values, which correspond to the missing numbers.
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	result := []int{}

	// Mark the presence of each number by negating the value at the corresponding index
	for i := 0; i < n; i++ {
		index := abs(nums[i]) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	// Collect the indices with positive values, which correspond to the missing numbers
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Follow-up solution
// This solution sorts the array in-place such that each number is placed at its correct index.
// It then iterates through the array to find the indices where the numbers are not in their correct positions.
func findDisappearedNumbersFollowUp(nums []int) []int {
	n := len(nums)
	result := []int{}

	// Place each number at its correct index
	for i := 0; i < n; i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// Collect the indices where the numbers are not in their correct positions
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			result = append(result, i+1)
		}
	}

	return result
}

// Edge conditions:
// - The input array is empty: return an empty array.
// - All numbers in the range [1, n] are present: return an empty array.
// - The input array contains duplicates: handle duplicates correctly in the optimized and follow-up solutions.