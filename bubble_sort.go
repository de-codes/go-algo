package main

import (
	"fmt"
	"math/rand"
)

func makeRandomSlice(numItems, max int) []int {
	res := make([]int, numItems)

	for i := 0; i < numItems; i++ {
		res[i] = rand.Intn(max)
	}

	return res
}

// Print at most numItems items.
func printSlice(slice []int, numItems int) {
	n := len(slice)

	if numItems < n {
		n = numItems
	}

	fmt.Println(slice[:n])

}

// Verify that the slice is sorted.
func checkSorted(slice []int) {
	isSorted := true

	for i := 1; i < len(slice); i++ {

		if slice[i] < slice[i-1] {
			isSorted = false
		}

	}

	if isSorted {
		fmt.Println("The Slice is sorted")
	} else {
		fmt.Println("The Slice is NOT sorted")
	}
}

// 3,2,1,7

func bubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		isSorted := true
		n := len(slice) - i

		for j := 1; j < n; j++ {

			if slice[j] < slice[j-1] {
				isSorted = false
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}

		}

		if isSorted {
			break
		}
	}
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
