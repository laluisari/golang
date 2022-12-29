package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {
	populations := make([]map[string]any, 0)
	for _, list := range data {
		populations = append(populations, formatPopulati(list))
	}
	return populations // TODO: replace this
}

func formatPopulati(datum string) map[string]any {
	population := make(map[string]any)

	split := strings.Split(datum, ";")

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

	return population

}
