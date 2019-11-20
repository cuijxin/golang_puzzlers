// question1:怎样把命令源码文件中的代码拆分到其他代码包？
// question2:代码包的导入路径总会与其所在目录的相对路径一致吗？
//
// 请记住，源码文件所在的目录相对于 src 目录的相对路径就是它的代码包导入路径，
// 而实际使用其程序实体时给定的限定符要与它声明所属的代码包名称对应。

// question3:什么样的程序实体才可以被当前包外的代码引用？
//
// 超级简单，名称的首字母为大写的程序实体才可以被当前包外的代码引用，
// 否则它就只能被当前包内的其他代码引用。
// 通过名称，Go 语言自然地把程序实体的访问权限划分为了包级私有的和公开的。
// 对于包级私有的程序实体，即使你导入了它所在的代码包也无法引用到它。
package main

import (
	"flag"
	"golang-puzzlers/src/puzzlers/article3/q2/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}
