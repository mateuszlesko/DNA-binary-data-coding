package main

import (
	"fmt"
	"strings"
)

func codeToDNA(data uint8) string {

	mDNA := map[int]string{
		0: "A",
		1: "G",
		2: "C",
		3: "T",
	}
	var sb strings.Builder
	var i uint8 = 0

	//8 is the upper limit of shift for a byte
	for x := 8; x >= 0; x -= 2 {
		i = (data >> x) & 0b00000011
		sb.WriteString(mDNA[int(i)])
	}
	return sb.String()
}

func main() {
	fmt.Println(codeToDNA(255))
	fmt.Println(codeToDNA(2))
	fmt.Println(codeToDNA(55))
	fmt.Println(codeToDNA(66))
}
