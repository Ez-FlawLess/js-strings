package main

import (
	"fmt"
	"js-strings/config"
	"js-strings/enums"
	"log"
	"os"
	"sync"
)

func main() {
	c := config.New()

	stringEnums := enums.New(c)

	fmt.Println(stringEnums)

	var wg sync.WaitGroup

	for _, stringFile := range c.StringFiles {

		wg.Add(1)
		func(sf config.StringFiles) {
			defer wg.Done()
			file, err := os.OpenFile(sf.Path, os.O_RDWR, os.ModeAppend)
			if err != nil {
				log.Fatalf("can't open string file with path %s with error: %v", sf.Path, err)
			}
			defer file.Close()

			n, err := file.WriteAt([]byte("\n"), 0)
			fmt.Println(n, err)
		}(stringFile)

	}

	wg.Wait()

}
