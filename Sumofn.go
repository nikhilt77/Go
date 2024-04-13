package main

import "fmt"

func sumOfNNumbers(n int) int {
	return n * (n + 1) / 2
}

func main() {
    fmt.Print("Enter Number:");
    var n int
    fmt.Scanln(&n)
	fmt.Println(sumOfNNumbers(n))
}
