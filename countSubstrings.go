// countSubstrings.go
package main

import (
	"os"
	//"io"
	"strings"
	"bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func countSubstrings(stringToFind string, inputFilename string, portionSize int64) (counter int64) {
	substringLen := len(stringToFind);
	buff := make([]byte, portionSize)
	prevPortion := make([]byte, 0)
	f, err := os.Open(inputFilename)
	check(err)
	defer f.Close()
	
	fi, err := f.Stat();
	check(err)
	filelength := fi.Size()
	
	buffReader := bufio.NewReader(f);
	var readPos int64
	for readPos=0; readPos<filelength; readPos+=portionSize{
		n, err := buffReader.Read(buff)
		check(err)
		stringToCheck := append(prevPortion, buff...)
		counter += int64(strings.Count(string(stringToCheck), stringToFind))
		if int64(n) == portionSize && substringLen>1 {
			prevPortion = buff[portionSize-int64(substringLen)+1:]
		}
	}
	return
}

