package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("FOO", "1")
	os.Setenv("BAR", "This is .env bar")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println("NOENV:", os.Getenv("NOENV"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
