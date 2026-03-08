package main

import (
	//"strconv"
	"github.com/neticdk/go-stdlib/diff/myers"
	"fmt"
)

func main() {
	diff, err := myers.Diff("hello\nworld", "hello\nthere\nworld",
	myers.WithContextLines(3), myers.WithShowLineNumbers(false))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(diff)
	}
}
