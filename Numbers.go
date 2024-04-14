package main

import "fmt"

func main() {
    count := 0
    fmt.Print("Enter number:")
    var n int
    fmt.Scanln(&n)
    for count < n {
        fmt.Print(fmt.Sprintf("%d->",count))
        count++
    }
    fmt.Print(count)
}
