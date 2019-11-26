package main

import "fmt"

// question: 怎样正确估算切片的长度和容量
func main() {
	// 示例1
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capactiy of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
}

// 现在你需要跟着我一起想象：有一个窗口，你可以通过这个窗口看到一个数组，
// 但是不一定能看到该数组中的所有元素，有时候只能看到连续的一部分元素。
// 现在，这个数组就是切片s2的底层数组，而这个窗口就是切片s2本身。
// s2的长度实际上指明的就是这个窗口的宽度，决定了你透过s2，可以看到其底层数组中的哪几个连续的元素。

//
