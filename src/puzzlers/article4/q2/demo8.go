// question:Go语言的类型推断可以带来哪些好处？
// 在声明变量的时候不指定类型，变量的类型通过给他赋值的表达式的类型推导得出，可以方便的对代码进行重构
// 即，不用改变变量的声明语句，来实现改变程序声明并赋值的逻辑
package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = getTheFlag() // 此处不用做修改，name的类型通过getTheFlag的实现来改变

	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name)
}

// func getTheFlag() *string {
// 	return flag.String("name", "everyone", "The greeting object.")
// }

func getTheFlag() *int {
	return flag.Int("num", 1, "The number of greeting object.")
}
