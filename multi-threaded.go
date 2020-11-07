package main

import (
    "fmt"
    "sync"
)

func main() {
	slice := []string{"a", "b", "c", "d", "e"}
    sliceLength := len(slice)

    var wg sync.WaitGroup
    // Tell the 'wg' WaitGroup how many threads/goroutines
    //   that are about to run concurrently.
    wg.Add(sliceLength)

    fmt.Println("Running for loop...")

    for i := 0; i < sliceLength; i++ {
        // Spawn a thread for each iteration in the loop.
        // Pass 'i' into the goroutine's function
        //   in order to make sure each goroutine
        //   uses a different value for 'i'.
        go func(i int) {
            // At the end of the goroutine, tell the WaitGroup
            //   that another thread has completed. 
            defer wg.Done()
            val := slice[i]
            fmt.Printf("i: %v, val: %v\n", i, val)
        }(i)
    }
    
    fmt.Println("Doing other stuff")

    // Wait for `wg.Done()` to be exectued the number of times
    //   specified in the `wg.Add()` call.
    // `wg.Done()` should be called the exact number of times
    //   that was specified in `wg.Add()`.
    // If the numbers do not match, `wg.Wait()` will either
    //   hang infinitely or throw a panic error.
    wg.Wait()

    fmt.Println("Finished for loop")
}