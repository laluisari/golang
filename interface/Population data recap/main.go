package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {
	//	"[name];[age];[address];[height];[is_married]"
	result := make([]map[string]any, 0)

	for _, check := range data {
		population := make(map[string]any)
		split := strings.Split(check, ";")
		population["name"] = split[0]
		age, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err)
		} else {
			population["age"] = age
		}
		population["address"] = split[2]

		if split[3] != "" {
			height, err := strconv.ParseFloat(split[3], 64)
			if err != nil {
				fmt.Println(err)
			} else {
				population["height"] = height
			}
		}

		if split[4] != "" {
			isMarried, err := strconv.ParseBool(split[4])
			if err != nil {
				fmt.Println(err)
			} else {
				population["isMarried"] = isMarried
			}
		}
		result = append(result, population)
	}

	return result // TODO: replace this
}

// func FormatData(temp string) []map[string]any {

// }

func main() {
	test := []string{
		"Budi;23;Jakarta;160.42;true",
		"Joko;30;Bandung;;true",
		"Susi;25;Bogor;;false",
	}
	fmt.Println(PopulationData(test))
}
