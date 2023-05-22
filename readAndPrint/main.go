// instructions:
// create program that reads contents of file & prints contents to terminal
// get file like this -> go run main.go myfile.txt
// use os.Args to get arguments provided
// use os.Open for opening file

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	b, _ := ioutil.ReadAll(file)
	fmt.Print(string(b))
}
