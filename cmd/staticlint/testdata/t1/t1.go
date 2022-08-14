package main

import (
	"fmt"
	"os"
)

func otherFunc() {
	os.Exit(1)
}

func main() {
	fmt.Println("test")
	x := 0
	fmt.Println(x)
	// os.Exit(x)
	otherFunc()
}
