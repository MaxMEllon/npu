package main

import (
	"fmt"
	"os"

	"github.com/maxmellon/npu/packages"
)

func main() {
	path := os.Args[1]
	packages, err := packages.Read(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", packages.Dependencies)
}
