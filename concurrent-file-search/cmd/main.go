package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matchesST []string
	matchesMT []string
	lock      sync.Mutex
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fileSearchST(root string, fileName string) {
	files, err := os.ReadDir(root)
	check(err)

	for _, file := range files {
		path := filepath.Join(root, file.Name())
		// match found
		if strings.Contains(file.Name(), fileName) {
			matchesST = append(matchesST, path)
		}
		// it is a directory
		if file.IsDir() {
			fileSearchST(path, fileName)
		}
	}
}

func fileSearchMT(root string, fileName string, wg *sync.WaitGroup) {
	defer wg.Done()

	files, err := os.ReadDir(root)
	check(err)

	for _, file := range files {
		path := filepath.Join(root, file.Name())
		// match found
		if strings.Contains(file.Name(), fileName) {
			// protect access to shared resource across threads
			lock.Lock()
			matchesMT = append(matchesMT, path)
			lock.Unlock()
		}
		// it is a directory
		if file.IsDir() {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				fileSearchMT(path, fileName, wg)
			}(wg)
		}
	}
}

func main() {
	fileSearchST("/Users/vivekchandela/personal", "README.md")

	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fileSearchMT("/Users/vivekchandela/personal", "README.md", wg)
	}(&wg)
	wg.Wait()

	fmt.Println("no. of matches in ST: ", len(matchesST))
	fmt.Println("no. of matches in MT: ", len(matchesMT))
}
