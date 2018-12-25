package main

import (
	"os"

	"github.com/maxmellon/nvu/entrypoint"
)

func main() {
	path := os.Args[1]
	entrypoint.All(path)
}
