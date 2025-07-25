package main

import "fmt"

func removeDuplicate(arr []int) []int {

	writePointer := 1
	for i := 1; i < len(arr); i++ {

		if arr[i] != arr[i-1] {
			arr[writePointer] = arr[i]
			writePointer++
		}
	}
	return arr[:writePointer]
}
func main() {

	input := []int{1, 2, 2, 2, 3, 4, 5, 5, 5, 6, 6, 6, 6}
	res := removeDuplicate(input)
	fmt.Printf("Input - %d \n output - %d ", input, res)

}
