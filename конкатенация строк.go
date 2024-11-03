package main

import (
	"fmt"
	"strings"
)

func concat(values []string) string {
	total := 0
	for i := 0; i < len(values); i++ { // проводятся итерации по каждой строке для вычисления общего числа байтов
		total += len(values[i])
	}
	sb := strings.Builder{}
	sb.Grow(total) // вызывается Grow с аргументом, равным этому общему числу
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

func main() {
	newList := []string{"hello", "hi"}
	result := concat(newList)
	fmt.Print(result)

}
