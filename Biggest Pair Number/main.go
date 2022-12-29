package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	parse := strconv.Itoa(numbers)
	var angka int
	var result int
	for a := 1; a < len(parse); a++ {
		b, _ := strconv.Atoi(string(parse[a]))
		c, _ := strconv.Atoi(string(parse[a-1]))
		hitung := b + c
		temp := string(parse[a-1]) + string(parse[a])
		//fmt.Println(temp)
		if hitung > angka {
			temp2, _ := strconv.Atoi(temp)
			angka = hitung
			result = temp2
		}
	}

	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
