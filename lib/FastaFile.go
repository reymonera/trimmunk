package FastaFile

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		fmt.Println(line1)
		fmt.Println(line2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al escanear el archivo:", err)
		return
	}
}
