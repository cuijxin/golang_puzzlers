package main

import "fmt"

// 问题2: 切片的底层数组什么时候会被替换？
// 确切地说，一个切片的底层数组永远不会被替换。
// 为什么？虽然在扩容的时候 Go 语言一定会生成新的底层数组，但是它也同时生成了新的切片。
// 它只是把新的切片作为了新底层数组的窗口，而没有对原切片，及其底层数组做任何改动。
// 请记住，在无需扩容时，append函数返回的是指向原底层数组的新切片，而在需要扩容时，append函数返回的是指向新底层数组的新切片。
// 所以，严格来讲，“扩容”这个词用在这里虽然形象但并不合适。
// 不过鉴于这种称呼已经用得很广泛了，我们也没必要另找新词了。
// 顺便说一下，只要新长度不会超过切片的原容量，那么使用append函数对其追加元素的时候就不会引起扩容。
// 这只会使紧邻切片窗口右边的（底层数组中的）元素被新的元素替换掉。
func main() {
	// 示例1
	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n",
		a1, len(a1), cap(a1))
	s9 := a1[1:4]
	fmt.Printf("s9: %v (len: %d, cap: %d)\n",
		s9, len(s9), cap(s9))
	for i := 1; i <= 5; i++ {
		s9 = append(s9, i)
		fmt.Printf("s9(%d): %v (len: %d, cap: %d)\n",
			i, s9, len(s9), cap(s9))
	}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n", a1, len(a1), cap(a1))
	fmt.Println()
}