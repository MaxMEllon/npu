package entrypoint

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/maxmellon/nvu/packages"
	"github.com/maxmellon/nvu/registry"
	"github.com/maxmellon/nvu/semver"
)

// All - show status of all modules
func All(path string) {
	packages, err := packages.Read(path)
	if err != nil {
		panic(err)
	}
	client, err := registry.NewClient()
	if err != nil {
		panic(err)
	}
	wg := &sync.WaitGroup{}
	for k, v := range packages.Modules {
		wg.Add(1)
		go func(moduleName string, version *semver.Semver) {
			defer wg.Done()
			latest, err := client.GetLatest(moduleName)
			l, _ := semver.Parse(latest)
			if err != nil {
				log.Println(err)
			}
			versions, err := client.GetAllVersions(moduleName)
			if err != nil {
				log.Println(err)
			}
			max, _ := semver.Parse(strings.Join([]string{
				string(version.Major),
				string(version.Miner),
				"-1",
			}, "."))
			for k := range versions {
				ve, _ := semver.Parse(k)
				if ve.Major == version.Major && ve.Miner == version.Miner && ve.Patch != version.Patch {
					if max.Patch < ve.Patch {
						max = ve
					}
				}
			}
			if version.GreaterThanPatch(max) {
				fmt.Printf("\U0001f300 [PATCH] %45s %12s => %-12s (%s)\n", moduleName, version.Raw, max.Raw, latest)
			} else if !version.Eq(l) {
				fmt.Printf("\U0001f195 [MAJOR] %45s %12s => %-12s\n", moduleName, version.Raw, latest)
			} else {
				fmt.Printf("\U0001f49a [ KEEP] %45s %12s\n", moduleName, version.Raw)
			}
		}(k, v)
	}
	wg.Wait()
}
