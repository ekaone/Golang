package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c", "d", "e"}
    sliceLength := len(slice)
    fmt.Println("Running for loop…")
    for i := 0; i < sliceLength; i++ {
        val := slice[i]
        fmt.Printf("i: %v, val: %v\n", i, val)
    }
    fmt.Println("Doing other stuff")
    fmt.Println("Finished for loop")
}
