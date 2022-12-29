package main

import "fmt"

func SchedulableDays(villager [][]int) []int {
	if len(villager) == 0 {
		return []int{}
	} else {
		duplicate := make(map[int]int)

		for a := 0; a < len(villager); a++ {
			for b := 0; b < len(villager[a]); b++ {
				num := villager[a][b]
				_, exist := duplicate[num]
				if exist {
					duplicate[num] += 1
				} else {
					duplicate[num] = 1
				}
			}
		}
		//semuanya harus sepakat
		result := make([]int, 0)

		for key, isi := range duplicate {
			if isi == len(villager) {
				result = append(result, key)
			}
		}

		return result
		// TODO: replace this
	}
}

func main() {
	data := [][]int{
		{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19},
	}
	fmt.Println(SchedulableDays(data))

}
