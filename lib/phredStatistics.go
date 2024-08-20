package lib

import "fmt"

func asciiInterpreter(ascii_line string) {
	for _, char := range ascii_line {
		phred_slice := []int{}
		phred_score := int(char)
		phred_slice = append(phred_slice, phred_score)
	}
	fmt.Println("Y ahora se va graficando en Python lol")
}
