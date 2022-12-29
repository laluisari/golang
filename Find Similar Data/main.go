package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	str := ""
	var slice []string

	for _, isi := range data {
		tes := strings.Contains(isi, input)
		if tes {
			slice = append(slice, isi)

		}
	}
	str += strings.Join(slice, ",")
	return str
}

func main() {

	fmt.Println(FindSimilarData("mobil", "mobil avanza", "mobil pajero", "motor nmax"))
}
