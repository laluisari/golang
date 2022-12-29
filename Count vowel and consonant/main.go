package main

import (
	"fmt"
	"strings"
)

func checkAIUEO(char rune) bool {
	if char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o' {
		return true
	} else {
		return false
	}
}

func CountVowelConsonant(str string) (int, int, bool) {
	str = strings.ToLower(str)
	vokal := 0
	konsonan := 0
	spasi := 0

	for _, a := range str {
		if checkAIUEO(a) {
			vokal += 1
		} else if a == ' ' || a == ',' || a == '/' {
			spasi += 1
		} else {

			konsonan += 1
		}

	}
	if vokal == 0 {
		return vokal, konsonan, true
	} else {
		return vokal, konsonan, false
	}
}

func main() {
	fmt.Println(CountVowelConsonant("SEMANGAT PAGI, itu kata orang yang baru datang ke rumahku"))
}
