package lib

import (
	in "golang-puzzlers/src/puzzlers/article3/q4/flag/internal"
	"os"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
