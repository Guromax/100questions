package main

import "fmt"

func main() {
	arr := []int{2, 5, 6, 1, 3}
	// Получение всех значений
	min1, max1 := MinMax(arr)
	fmt.Println("min1 =", min1)
	fmt.Println("max1 =", max1)
	// Получение только первого значения
	min1, _ = MinMax(arr)
	fmt.Println("min1 =", min1)
	// Получение только второго значения
	_, max1 = MinMax(arr)
	fmt.Println("max1 =", max1)
}

func MinMax(arr []int) (int, int) {
	min1 := arr[0]
	max1 := arr[0]
	for _, value := range arr {
		if value < min1 {
			min1 = value
		}
		if value > max1 {
			max1 = value
		}
	}
	return min1, max1
}

// min = 1
// max = 6
// min = 1
// max = 6
