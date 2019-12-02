package main

import "fmt"

func main() {
	// aMap := map[string]int{
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// }

	// k := "two"
	// v, ok := aMap[k]
	// if ok {
	// 	fmt.Printf("The element of key %q: %d\n", k, v)
	// } else {
	// 	fmt.Println("Not found!")
	// }

	var badMap2 = map[interface{}]int{
		"1":      1,
		[]int{2}: 2,
		3:        3,
	}

	k := "two"
	v, ok := badMap2[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}
}
