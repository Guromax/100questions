package main

import "fmt"

func main() {
	fmt.Println(Append([]int{1, 2, 3}, 4))
}

func Append[T any](dst []T, el T) []T {
	var res []T

	resLen := len(dst) + 1
	if resLen <= cap(dst) {
		res = dst[:resLen]
	} else {
		resCap := resLen
		if resCap < 2*len(dst) {
			resCap = 2 * len(dst)
		}

		res = make([]T, resLen, resCap)
		copy(res, dst)
	}

	res[len(dst)] = el
	return res
}
