package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Reverse(str string) string {
	var temp string
	for a := len(str) - 1; a >= 0; a-- {
		temp += string(str[a])
	}
	return temp // TODO: replace this
}

func Generate(str string) string {
	str = Reverse(str)

	var temp string
	for _, str := range str {

		if unicode.IsUpper(str) {
			if str == 'A' {
				str = 'e'
			} else if str == 'E' {
				str = 'i'
			} else if str == 'I' {
				str = 'o'
			} else if str == 'O' {
				str = 'u'
			} else if str == 'U' {
				str = 'a'
			}
		} else {
			if str == 'a' {
				str = 'E'
			} else if str == 'e' {
				str = 'I'
			} else if str == 'i' {
				str = 'O'
			} else if str == 'o' {
				str = 'U'
			} else if str == 'u' {
				str = 'A'
			}
		}

		temp += string(str)

	}
	var temp2 string

	for _, huruf := range temp {
		if huruf == 'a' || huruf == 'i' || huruf == 'u' || huruf == 'e' || huruf == 'o' ||
			huruf == 'A' || huruf == 'I' || huruf == 'U' || huruf == 'E' || huruf == 'O' {
			temp2 += string(huruf)
		} else {
			if unicode.IsUpper(huruf) {
				temp2 += strings.ToLower(string(huruf))
			} else {
				temp2 += strings.ToUpper(string(huruf))
			}
		}

	}
	temp2 = strings.ReplaceAll(temp2, " ", "")

	return temp2 // TODO: replace this
}

func CheckPassword(str string) string {
	str = Generate(str)
	response := ""
	num := "1234567890"
	sim := "%%&`$^(){}-=+!@#"
	// for _, char := range str {
	// 	if char != '1' || char != '2' || char != '3' || char != '4' || char != '5' ||
	// 	char != '6' || char != '7' || char != '8' || char != '9' || char
	// }

	if len(str) < 7 {
		response += "sangat lemah"
	} else if len(str) >= 7 && strings.ContainsAny(str, sim) == false {
		response += "lemah"
	} else if len(str) >= 7 && len(str) < 14 {
		if strings.ContainsAny(str, num) || strings.ContainsAny(str, sim) {
			response += "sedang"
		}
	} else if len(str) >= 14 || strings.ContainsAny(str, num) || strings.ContainsAny(str, sim) {
		response += "kuat"
	}
	// TODO: replace this
	return response
}

func PasswordGenerator(base string) (string, string) {
	base = Generate(base)
	base2 := CheckPassword(base)
	return base, base2 // TODO: replace this
}

func main() {

	// fmt.Println(CheckPassword("Adm1n!@"))

	data := "Semangat Pagi" // bisa digunakan untuk melakukan debug

	res, check := PasswordGenerator(data)

	fmt.Println(res)
	fmt.Println(check)

}

// || strings.ContainsAny(str, num) {
// 	response += "lemah"
// } else if len(str) >= 7 && len(str) < 14 && strings.ContainsAny(str, num) && strings.ContainsAny(str, sim) {
// 	response += "sedang"
// } else if len(str) >= 14 && strings.ContainsAny(str, num) && strings.ContainsAny(str, sim) {
// 	response += "kuat"
