package main

import "fmt"

func FindMin(num ...int) int {
	temp := num[0]

	for a := 1; a < len(num); a++ {
		if temp < num[a] {
			result := temp
			temp = result
		} else if temp > num[a] {
			result := num[a]
			temp = result
		} else {
			result := temp
			temp = result
		}
	}

	return temp
}

func FindMax(num ...int) int {
	temp := num[0]

	for a := 1; a < len(num); a++ {
		if temp > num[a] {
			result := temp
			temp = result
		} else if temp < num[a] {
			result := num[a]
			temp = result
		} else {
			result := temp
			temp = result
		}
	}

	return temp
}

func SumMinMax(num ...int) int {
	result := FindMin(num...) + FindMax(num...)
	return result

}

func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
