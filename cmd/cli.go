package main

import (
	"os"

	"github.com/maxmellon/npu/packages/packages"
)

func main() {
	err := packages.Read(os.Args[1])
	if err != nil {
		panic(err)
	}
}
