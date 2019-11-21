package main

import (
	"flag"
	"fmt"
)

// quetion：声明变量有几种方式？

func main() {
	var name string                                                   //[1]
	flag.StringVar(&name, "name", "everyone", "The greeting object.") //[2]

	// 方式1
	// var name = flag.String("name", "everyone", "The greeting object.") // [3] 基于表达式的类型判断，表达式类型就是对表达式进行求值后得到结果的类型

	// 方式2
	// name := flag.String("name", "everyone", "The greeting object.") // [4] 短变量声明，只能在函数内部使用

	flag.Parse()
	fmt.Printf("Hello, %v\n", name)

	// 适用于方式1和方式2
	// fmt.Printf("Hello, %v!\n", *name) // [3]
}
