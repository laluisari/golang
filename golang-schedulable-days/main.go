package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {
	duplicate := make(map[int]bool)
	result := make([]int, 0)
	for a := 0; a < len(date1); a++ {
		duplicate[date1[a]] = true
	}

	for _, data := range date2 {
		if duplicate[data] {
			result = append(result, data)

		}
	}

	return result
}

func main() {
	date2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	date1 := []int{14, 15, 16, 17, 18, 19, 20}

	fmt.Println(SchedulableDays(date1, date2))

}
