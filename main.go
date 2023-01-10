package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches   []string
	waitgroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
	rootInput string
	fileName  string
)

func fileSearch(root, fileName string) {
	fmt.Println("searching in ", root)

	files, err := ioutil.ReadDir(root)

	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), filepath.Join(root, file.Name())) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()), fileName)
		}

		if file.IsDir() {
			waitgroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), fileName)
		}
	}
	waitgroup.Done()

}

func main() {
	//fetching filename
	fmt.Print("enter filename <name.ext>: ")
	fmt.Scanln(&fileName)

	//fetching root directory
	fmt.Print("enter absolute root directory: ")
	fmt.Scanln(&rootInput)

	waitgroup.Add(1)
	go fileSearch(rootInput, fileName)
	waitgroup.Wait()

	for _, file := range matches {
		fmt.Println("Matched: ", file)
	}
}
