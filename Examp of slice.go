package main

import "fmt"

func main() {
    sample := []int{} // it is creating a empyty slice
    fmt.Println(len(sample))
    fmt.Println(cap(sample))
    fmt.Println(sample)

    letters := []string{"a", "b", "c"}
    fmt.Println(len(letters)) // checking length
    fmt.Println(cap(letters)) // checking capacity
    fmt.Println(letters)
}