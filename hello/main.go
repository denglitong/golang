package main

import "fmt"

import (
	"github.com/google/go-cmp/cmp"
	"golang/hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world."))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
