package main

import (
	"fmt"
)

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var rest string
	b := 0
	for a := len(str) - 1; a >= b; a-- {
		if a == b {
			rest += string(str[a])
		} else if string(str[a]) == " " {
			//fmt.Print(string(str[a]))
			rest += string(str[a])
		} else if string(str[a-1]) == " " {
			rest += string(str[a])
		} else {
			rest += string(str[a]) + "_"
		}
	}
	return rest // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	// fmt.Println(ReverseString("I am a stranger"))
	fmt.Println(ReverseString("Hello World"))

}
