package main

import (
	"fmt"
	"strconv"
	"strings"

	"a21hc3NpZ25tZW50/internal"
)

func AdvanceCalculator(calculate string) float32 {
	split := strings.Split(calculate, " ") //untuk menghilangkan semua spasinya
	if len(split) < 1 {
		return 0
	}
	base, _ := strconv.Atoi(split[0])
	calculator := internal.NewCalculator(float32(base))
	for a := 1; a < len(split)-1; a++ {
		operator := split[a]                  //mengambil semua charakter
		number, _ := strconv.Atoi(split[a+1]) //mengambil semua huruf
		switch operator {
		case "+":
			calculator.Add(float32(number))
		case "-":
			calculator.Subtract(float32(number))
		case "*":
			calculator.Multiply(float32(number))
		case "/":
			calculator.Divide(float32(number))
		}
	}

	return calculator.Result() // TODO: replace this
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
