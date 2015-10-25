// hugefind
package main

import (
	"fmt"
	"os"
	"log"
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
}

func parseArgs() error {
	if len(os.Args)!=3 { // 0 - exe filename
		return errors.New("It is expected 2 arguments.")
	}
	stringToFind = os.Args[1]
	inputFilename = os.Args[2]

	log.Println("String to find: '"+stringToFind+"'")
	log.Println("File to processing: '"+inputFilename+"'")

	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return errors.New("File '"+inputFilename+"' is not exit.")
	}

	return nil;
}

func exeName() (result string){
	result = filepath.Base(os.Args[0])
	return
}