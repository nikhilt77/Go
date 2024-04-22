package main

import (
    "fmt"
    "sort"
)

func main() {
    li := []int{10, 20, 30, 40, 50}

    li1 := make([]int, 0)
    for _, i := range li {
        if i >= 30 {
            li1 = append(li1, i)
        }
    }
    fmt.Println(li1)

    li2 := make([]int, 0)
    for _, i := range li {
        if i >= 30 {
            li2 = append(li2, i)
        }
    }
    fmt.Println("filter=", li2)

    li3 := make([]int, len(li))
    for i, v := range li {
        li3[i] = v * 2
    }
    fmt.Println("map=", li3)

    count := len(li)
    fmt.Println("count=", count)

    for _, i := range li {
        fmt.Print(i*3, " ")
    }
    fmt.Println("\nforEach=", li)

    res := 0
    for _, i := range li {
        res += i
    }
    fmt.Println("reduce()= ", res)

    li5 := make([]int, len(li))
    copy(li5, li)
    sort.Ints(li5)
    fmt.Println("sorted=", li5)

    li6 := make([]int, len(li))
    copy(li6, li)
    sort.Slice(li6, func(i, j int) bool { return li6[i] > li6[j] })
    fmt.Println(li6)
}
