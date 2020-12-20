package main

import (
	"os"

	"github.com/ddddddO/extxt"
)

func main() {
	serviceAccountFile := "/mnt/c/Users/lbfde/Downloads/tag-mng-b8e1b87744fc.json"
	in := "../../testdata/image.JPG"
	out := os.Stdout

	if err := extxt.Run(out, in, serviceAccountFile); err != nil {
		panic(err)
	}
}
