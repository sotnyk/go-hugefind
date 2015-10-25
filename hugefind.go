// hugefind
package main

import (
	"fmt"
	"os"
	"errors"
	"path/filepath"
)

var inputFilename string
var stringToFind string

func main() {
	if err := parseArgs(); err!=nil{
		fmt.Println(err)
		fmt.Println("")
		fmt.Println("Usage: "+exeName()+" substring filename")
		return
	}

	counter := countSubstrings(stringToFind, inputFilename, 1024*1024)
	fmt.Println(fmt.Sprintf("%d entries found.", counter))
}

func parseArgs() error {
	if len(os.Args)!=3 { // 0 - exe filename
		return errors.New("It is expected 2 arguments.")
	}
	stringToFind = os.Args[1]
	inputFilename = os.Args[2]

	fmt.Println("String to find: '"+stringToFind+"'")
	fmt.Println("File to processing: '"+inputFilename+"'")

	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return errors.New("File '"+inputFilename+"' is not exist.")
	}

	if len(stringToFind)==0 {
		return errors.New("Substring should not be empty")
	}

	return nil;
}

func exeName() (result string){
	result = filepath.Base(os.Args[0])
	return
}