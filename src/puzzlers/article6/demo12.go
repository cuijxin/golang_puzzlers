package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])
}

// Question：怎样判断一个变量的类型？
// 使用“类型断言”表达式：value, ok := interface{}(container).([]string)
