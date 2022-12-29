package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	var a float64 = float64(height)
	if gender == "laki-laki" {
		temp := (a - 100) - ((a - 100) * 10 / 100)
		return temp
	} else {
		temp := (a - 100) - ((a - 100) * 15 / 100)
		return temp
	}
}

func main() {
	fmt.Println(BMICalculator("laki-laki", 170))
}
