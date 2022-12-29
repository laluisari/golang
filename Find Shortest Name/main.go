package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {
	var sample []rune
	for _, char := range names {
		if char == ' ' || char == ',' || char == ';' {
			sample = append(sample, char)
		}
	}
	sliceNames := strings.Split(names, string(sample[0]))
	temp := sliceNames[0]
	var last [1]string

	for a := 1; a < len(sliceNames); a++ {
		if len(temp) < len(sliceNames[a]) {
			last[0] = temp
		} else if len(temp) > len(sliceNames[a]) {
			temp = sliceNames[a]
			last[0] = temp
		} else if len(temp) == len(sliceNames[a]) {
			if temp < sliceNames[a] {
				last[0] = temp
			} else {
				last[0] = sliceNames[a]
				temp = last[0]
			}
		}
	}

	return last[0]
}

func main() {
	fmt.Println(FindShortestName("angga isari"))
}
