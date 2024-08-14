package lib

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"os"
)

//func ImportFile(file string) (string, error) {
//	data, err := ioutil.ReadFile(file)
//	if err != nil {
//		fmt.Println(err)
//		return "Oops! An error ocurred!", err
//	}
//	return string(data), nil
//}

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
		scanner.Scan()
		fastaline2 := scanner.Text()
		fmt.Println(fastaline1)
		fmt.Println(fastaline2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning:", err)
		return
	}
}
