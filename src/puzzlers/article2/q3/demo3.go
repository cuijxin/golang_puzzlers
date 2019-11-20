package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

// 方法3
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	// 方法2
	// flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	// flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	// flag.CommandLine.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }

	// flag.StringVar(&name, "name", "everyone", "The greeting object.")

	// 方法3
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	// 方法1
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }

	// flag.Parse()

	// 方法3
	cmdLine.Parse(os.Args[1:])

	fmt.Printf("Hello, %s!\n", name)
}
