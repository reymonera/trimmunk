package lib

import (
	"bufio"
	"fmt"
	"os"
)

func GetSingleLine(filepath string) {
	fastafile, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error trying to open the file", err)
		return
	}
	defer fastafile.Close()

	scanner := bufio.NewScanner(fastafile)
	for scanner.Scan() {
		fastaline1 := scanner.Text()
		//fmt.Println("1st line")
		//fmt.Println(fastaline1)
		scanner.Scan()
		fastaline2 := scanner.Text()
		//fmt.Println("2nd line")
		//fmt.Println(fastaline2)
		scanner.Scan()
		fastaline3 := scanner.Text()
		//fmt.Println("3rd line")
		//fmt.Println(fastaline3)
		scanner.Scan()
		fastaline4 := scanner.Text()

		_ = fastaline1
		_ = fastaline2
		_ = fastaline3
		_ = fastaline4
		//fmt.Println("4th line")
		//fmt.Println(fastaline4)

		//if scanner.Scan(){
		//		fastaline2 := scanner.Text()
		//} else {
		//	fastaline2 = ""
		//	fmt.Println("Odd number of lines detected, or error occurred.")
		//	break
		//}

		//fmt.Println("Saved .fasta files in variables in Trimmunk")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning:", err)
		return
	}
}
