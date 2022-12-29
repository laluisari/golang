package main

import (
	"fmt"
	"unicode"
)

func SlurredTalk(words *string) {
	temp := ""
	for _, check := range *words {
		if unicode.IsUpper(check) {
			if check == 'S' || check == 'R' || check == 'Z' {
				check = 'L'
			}
		} else {
			if check == 's' || check == 'r' || check == 'z' {
				check = 'l'
			}
		}
		temp += string(check)
	}
	*words = temp
}

func main() {

	// bisa dicoba untuk pengujian test case
	var words string = "Steven"
	SlurredTalk(&words)
	fmt.Println(words)
}
