package main

import "fmt"

func intersection(a, b []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, num := range a {
		if _, ok := counter[num]; !ok {
			counter[num] = 1
		} else {
			counter[num] += 1
		}
	}

	for _, num := range b {
		if count, ok := counter[num]; ok && count > 0 {
			counter[num] -= 1
			result = append(result, num)
		}
	}
	return result

}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 6, 7}
	fmt.Printf("%v", intersection(a, b))
}
