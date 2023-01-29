package fastaFile

import (
	"fmt"
	"io/ioutil"
)

func ImportFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(data), nil
}

func getSingleLine() {

}
