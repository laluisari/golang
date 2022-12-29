package main

import (
	"fmt"
	"strings"
)

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	jmlh := 0
	huruf := strings.ToLower(text)
	for _, scanHuruf := range huruf {
		if scanHuruf == 'r' || scanHuruf == 's' ||
			scanHuruf == 't' || scanHuruf == 'z' {
			jmlh += 1
		}
	}

	return jmlh // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Remaja Riang"))
}
