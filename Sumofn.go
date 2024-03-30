package main

import "fmt"

func sumOfNNumbers(n int) int {
	return n * (n + 1) / 2
}

func main() {
	fmt.Println(sumOfNNumbers(10))
}
