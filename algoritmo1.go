package main

// Jeito mais eficiente
// func sumNumbers(n int) int {
// 	return n * (n + 1) / 2
// }

// Jeito menos eficiente
func sumNumbers(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	println(sumNumbers(10))
}
