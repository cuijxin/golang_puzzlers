// question:对于程序实体，还有其他的访问权限规则吗？
// 我们可以通过创建internal代码包让一些程序实体仅仅能被当前模块中的其他代码引用。
// 这被称为 Go 程序实体的第三种访问权限：模块级私有。
// 具体规则是，internal代码包中声明的公开程序实体仅能被该代码包的直接父包及其子包中的代码引用。
// 当然，引用前需要先导入这个internal包。对于其他代码包，导入该internal包都是非法的，无法通过编译。
package main

import (
	"flag"
	in "golang-puzzlers/src/puzzlers/article3/q4/lib/internal" // 此行无法通过编译。
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	// lib.Hello(name)
	in.Hello(name)
}
