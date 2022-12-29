package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	var hasil string
	ArrIndex := strings.Split(str, " ")

	for a, index := range ArrIndex {
		var temp string
		for b := len(index) - 1; b >= 0; b-- {
			if unicode.IsUpper(rune(index[0])) && b == len(index)-1 {
				temp = temp + strings.ToUpper(string(index[b]))
			} else if unicode.IsUpper(rune(index[b])) {
				temp = temp + strings.ToLower(string(index[b]))
			} else {
				temp += string(index[b])
			}
		}
		if a == len(ArrIndex)-1 {
			hasil += temp
		} else if a != len(ArrIndex)-1 {
			hasil += temp + " "
		} else {
			return hasil
		}

	}

	return hasil

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	//fmt.Println(ReverseWord("A bird fly to the Sky"))

}
