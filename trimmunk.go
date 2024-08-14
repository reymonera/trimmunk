package main

import (
	"flag"
	"fmt"
	"trimmunk/lib"
	//"io/ioutil"
    //"os"
)

func main() {

	// Define the flags
    filePath1 := flag.String("f1", "", "Path to the first file")
    filePath2 := flag.String("f2", "", "Path to the second file")

    // Parse the flags
    flag.Parse()

    // Check if both file paths are provided
    if *filePath1 == "" || *filePath2 == "" {
        fmt.Println("Please provide both file paths using the -f1 and -f2 flags")
        return
    }

	// Announcing the content proceedings are complete.
	fmt.Println("Completed reading the .fastq files!")

	for _, filePath := range []string{*filePath1, *filePath2} {
		//fasta, err := lib.GetSingleLine(filePath)
		lib.GetSingleLine(filePath)
		//if err != nil {
		//	fmt.Println(err)
		//} else {
		//	fmt.Println(fasta)
		//}
	}
}