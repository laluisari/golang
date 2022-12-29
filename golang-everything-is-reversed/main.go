package main

import (
	"fmt"
	"strconv"
)

func ReverseData(arr [5]int) [5]int {

	list := 0
	str := ""
	result := [5]int{}
	num := 0
	temp2 := [5]string{}
	for a := len(arr) - 1; a >= 0; a-- {
		num = arr[a]
		temp2[list] = strconv.Itoa(num)
		list++
	}

	for x := 0; x < len(temp2); x++ {
		chars := []rune(temp2[x])
		for a := len(chars) - 1; a >= 0; a-- {
			if a == 0 {
				str += string(chars[a]) + ""
			} else {
				str += string(chars[a])
			}
		}

		number, _ := strconv.Atoi(str)

		fmt.Println(number)
		result[x] = number

		str = ""

	}
	fmt.Println(result)
	// var string1 = strings.Split(str, ",")
	// fmt.Println(string1)
	return result // TODO: replace this
}

func main() {
	data := [5]int{
		0, 0, 0, 0, 0,
	}
	fmt.Println(ReverseData(data))

}
