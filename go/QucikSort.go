package main

import "fmt"

func QuickSort(arr []int, low, high int) {
	fmt.Printf("=== QuickSort called with range [%d:%d] ===\n", low, high)
	fmt.Printf("Current array: %v\n", arr)
	fmt.Printf("Subarray being sorted: %v\n", arr[low:high+1])

	if low < high {
		fmt.Printf("Condition (low=%d < high=%d) is true, proceeding with sorting\n", low, high)

		// Partition step
		fmt.Printf("\n--- Starting Partition ---\n")
		pivotIndex := Partition(arr, low, high)
		fmt.Printf("--- Partition Complete ---\n")
		fmt.Printf("Pivot settled at index %d with value %d\n", pivotIndex, arr[pivotIndex])
		fmt.Printf("Array after partitioning: %v\n", arr)

		// Left recursive call
		if low < pivotIndex-1 {
			fmt.Printf("\n>>> Recursively sorting LEFT side: [%d:%d]\n", low, pivotIndex-1)
			QuickSort(arr, low, pivotIndex-1)
			fmt.Printf("<<< Returned from LEFT side sorting [%d:%d]\n", low, pivotIndex-1)
		} else {
			fmt.Printf("\n>>> LEFT side [%d:%d] has 0 or 1 elements, skipping\n", low, pivotIndex-1)
		}

		// Right recursive call
		if pivotIndex+1 < high {
			fmt.Printf("\n>>> Recursively sorting RIGHT side: [%d:%d]\n", pivotIndex+1, high)
			QuickSort(arr, pivotIndex+1, high)
			fmt.Printf("<<< Returned from RIGHT side sorting [%d:%d]\n", pivotIndex+1, high)
		} else {
			fmt.Printf("\n>>> RIGHT side [%d:%d] has 0 or 1 elements, skipping\n", pivotIndex+1, high)
		}

		fmt.Printf("\n=== QuickSort [%d:%d] completed ===\n", low, high)
		fmt.Printf("Final array state: %v\n\n", arr)
	} else {
		fmt.Printf("Condition (low=%d < high=%d) is false, base case reached\n", low, high)
		fmt.Printf("=== QuickSort [%d:%d] completed (base case) ===\n\n", low, high)
	}
}

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	fmt.Printf("Partitioning range [%d:%d] with pivot = %d (at index %d)\n", low, high, pivot, high)
	fmt.Printf("Array before partitioning: %v\n", arr[low:high+1])

	i := low
	fmt.Printf("Starting with i = %d (partition index)\n", i)

	for j := low; j < high; j++ {
		fmt.Printf("  Comparing arr[%d] = %d with pivot = %d", j, arr[j], pivot)
		if arr[j] <= pivot {
			fmt.Printf(" -> %d <= %d, swapping arr[%d] with arr[%d]\n", arr[j], pivot, i, j)
			fmt.Printf("    Before swap: arr[%d] = %d, arr[%d] = %d\n", i, arr[i], j, arr[j])
			arr[i], arr[j] = arr[j], arr[i]
			fmt.Printf("    After swap:  arr[%d] = %d, arr[%d] = %d\n", i, arr[i], j, arr[j])
			i++
			fmt.Printf("    Incrementing i to %d\n", i)
		} else {
			fmt.Printf(" -> %d > %d, no swap needed\n", arr[j], pivot)
		}
		fmt.Printf("    Current array: %v\n", arr[low:high+1])
	}

	fmt.Printf("Final step: placing pivot at correct position\n")
	fmt.Printf("  Swapping pivot arr[%d] = %d with arr[%d] = %d\n", high, arr[high], i, arr[i])
	arr[i], arr[high] = arr[high], arr[i]
	fmt.Printf("  After final swap: %v\n", arr[low:high+1])
	fmt.Printf("Pivot %d is now at index %d\n", pivot, i)

	return i
}

func main() {
	fmt.Println("-- QUICk SORT ---")
	testingKit := []struct {
		input    []int
		expected []int
		high     int
		low      int
	}{
		// Test Case 1: Typical unsorted array
		{
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		// Test Case 2: Empty array
		{
			input:    []int{},
			expected: []int{},
		},
		// Test Case 3: Single element
		{
			input:    []int{1},
			expected: []int{1},
		},
		// Test Case 4: Array with duplicates
		{
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9},
		},
		// Test Case 5: Already sorted array
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		// Test Case 6: Reverse sorted array
		{
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		// Test Case 7: All identical elements
		{
			input:    []int{4, 4, 4, 4, 4},
			expected: []int{4, 4, 4, 4, 4},
		},
		// Test Case 8: Negative numbers
		{
			input:    []int{-5, 2, -10, 0, 8, -3},
			expected: []int{-10, -5, -3, 0, 2, 8},
		},
		// Test Case 9: Large numbers
		{
			input:    []int{1000000, 500, 10000, 50, 100000},
			expected: []int{50, 500, 10000, 100000, 1000000},
		},
		// Test Case 10: Mixed positive and negative with duplicates
		{
			input:    []int{0, -2, 3, -2, 0, 5, 3},
			expected: []int{-2, -2, 0, 0, 3, 3, 5},
		},
	}

	for i := range testingKit {
		testingKit[i].high = len(testingKit[i].input) - 1
		testingKit[i].low = 0
	}

	for i := range testingKit {
		QuickSort(testingKit[i].input, testingKit[i].low, testingKit[i].high)
	}

}
