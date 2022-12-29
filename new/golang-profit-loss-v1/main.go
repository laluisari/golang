package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	res := make([]string, 0)
	data, err := os.ReadFile(path)

	if err != nil {
		return res, err
	}

	temp := string(data[:])
	if len(temp) == 0 {
		return res, err
	}
	output := strings.Split(temp, "\n")
	fmt.Println(output)

	return output, nil // TODO: replace this
}

func CalculateProfitLoss(data []string) string {
	date := ""
	status := ""
	profit := 0

	for i := 0; i < len(data); i++ {
		split := strings.Split(data[i], ";")
		date = split[0]
		uang, _ := strconv.Atoi(split[2])

		if split[1] == "income" {
			profit += uang
		} else {
			profit -= uang
		}
		// fmt.Println(profit)
		// income += in
		if profit < 0 {
			status = "loss"
		} else {
			status = "profit"
		}
	}
	if profit < 0 {
		profit = -profit
	}
	output := date + ";" + status + ";" + strconv.Itoa(profit)
	// fmt.Println(output)
	return output // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
