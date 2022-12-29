package main

import (
	"fmt"
)

func CountProfit(data [][][2]int) []int {

	if len(data) == 0 {
		return []int{}
	} else {
		var profit = make([]int, len(data[0]))
		for a := 0; a < len(data); a++ {
			for b := 0; b < len(data[a]); b++ {
				test := data[a][b][0] - data[a][b][1]
				profit[b] += test
			}
		}

		return profit[:]
	}

	// TODO: replace this
}

func main() {

	//Creating slice
	slc := [][][2]int{
		{
			{1000, 500}, {500, 200},
		},
		{
			{1200, 200}, {1000, 800},
		},
		{
			{500, 100}, {700, 100},
		},
	}

	fmt.Println(CountProfit(slc))

}
