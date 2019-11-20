// question：怎样把命令文件中的代码拆分到其他库源码文件？
package main

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	hello(name)
}
