package main

import "fmt"

func BinarySearch(in []int, searchFor int) (int, bool) {
	if len(in) == 0 {
		return 0, false
	}

	var first, last = 0, len(in) - 1

	for first <= last {
		var mid = ((last - first) / 2) + first

		if in[mid] == searchFor {
			return mid, true
		} else if in[mid] > searchFor { // нужно искать в "левой" части слайса
			last = mid - 1
		} else if in[mid] < searchFor { // нужно искать в "правой" части слайса
			first = mid + 1
		}
	}

	return 0, false
}

func main() {
	textSlice := []int{1, 3, 4, 6, 8, 10, 55, 56, 59, 70, 79, 81, 91, 10001}
	result, isIn := BinarySearch(textSlice, 55)
	fmt.Print(result, isIn)
}
