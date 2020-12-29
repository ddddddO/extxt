package main

import (
	"os"

	"github.com/ddddddO/extxt"
)

func main() {
	in := "./testdata/image.JPG"
	out := os.Stdout

	if err := extxt.Run(out, in); err != nil {
		panic(err)
	}
}
