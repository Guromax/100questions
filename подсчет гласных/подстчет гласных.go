package main

import "fmt"

func vowelCount(str string) int {
	count := 0
	for _, char := range str {
		switch char {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			count++
		}
	}
	return count
}

func main() {
	result := vowelCount("I see dead people")
	fmt.Print(result)
}
