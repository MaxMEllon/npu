package main

import (
	"fmt"
	"log"
	"os"

	"github.com/maxmellon/nvu/packages"
	"github.com/maxmellon/nvu/registry"
)

func main() {
	path := os.Args[1]
	packages, err := packages.Read(path)
	if err != nil {
		panic(err)
	}
	client := registry.NewClient()
	for k, v := range packages.Modules {
		// TODO: go func with chan
		version, err := client.GetLatest(k)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s %s => %s\n", k, v.Raw, version)
	}
}
