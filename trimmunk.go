package main

import (
	"fmt"
	FastaFile "trimmunk/lib"
)

func main() {
	result, err := FastaFile.ImportFile("text.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
