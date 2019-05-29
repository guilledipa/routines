// Just playing around with Go routines
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

const (
	winDir   = "d:/Test"
	routines = 10
)

var wg sync.WaitGroup

func findFiles(dirPath string, c chan os.FileInfo) {

	files, err := ioutil.ReadDir(winDir)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		c <- file
	}
	close(c)
}

func worker(c chan os.FileInfo) {

}

func handler(c chan os.FileInfo, i int) {
	for f := range c {
		fmt.Println(i, f.Name())
	}
	defer wg.Done()
}

func main() {

	c := make(chan os.FileInfo)
	go findFiles(winDir, c)

	for i := 0; i < routines; i++ {
		wg.Add(1)
		go handler(c, i)
	}

	wg.Wait()
	fmt.Println("main finished")
}
