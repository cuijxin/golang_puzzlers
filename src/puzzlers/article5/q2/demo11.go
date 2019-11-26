package main

import (
	// "fmt"
	"fmt"
	. "fmt"
)

func Printf(val interface{}) {
	fmt.Printf("The value is %v\n", val)
}

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])
}
