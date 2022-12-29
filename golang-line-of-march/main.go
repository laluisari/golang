package main

import (
	"fmt"
	"sort"
)

func Sortheight(height []int) []int {
	//if i < j, jika benar  kembalikan kondisi tersebut
	sort.Slice(height, func(i, j int) bool {
		return height[i] < height[j]
	})
	result := make([]int, 0)
	for _, v := range height {
		result = append(result, v)
	}

	return result // TODO: replace this
}

func main() {
	data := []int{172, 170, 150, 165, 144, 155, 159}

	fmt.Println(Sortheight(data))

}
