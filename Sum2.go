package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sumOfNNumbers(n int) int {
	return n * (n + 1) / 2
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	n, err := strconv.Atoi(str[:len(str)-1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumOfNNumbers(n))
}
