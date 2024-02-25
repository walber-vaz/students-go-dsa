package main

import (
	"fmt"
	"time"
)

// Time complexity: O(n)

// Jeito mais eficiente
// func sumNumbers(n int) int {
// 	return n * (n + 1) / 2
// }

// Jeito menos eficiente
func sumNumbers(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

func main() {
	timeStart := time.Now()
	_ = sumNumbers(9999999999)
	timeEnd := time.Now()

	fmt.Println("Tempo de execução: ", timeEnd.Sub(timeStart).Milliseconds(), "ms")
}
