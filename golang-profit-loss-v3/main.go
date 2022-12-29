package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	//bukak
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//baca
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if len(b) == 0 {
		return []string{}, nil
	} else {
		lines := strings.Split(string(b), "\n")
		return lines, nil

	}
}

func CalculateProfitLoss(data []string) string {

	date := ""
	ket := ""
	profit := 0
	result := ""

	for a := 0; a < len(data); a++ {
		split := strings.Split(data[a], ";")
		date = split[0]
		uang, _ := strconv.Atoi(split[2])

		if split[1] == "income" {
			profit += uang
		} else {
			profit -= uang
		}

		if profit < 0 {
			ket = "loss"
		} else {
			ket = "profit"
		}
	}

	if profit < 0 {
		profit = -profit
	}
	result = fmt.Sprintf("%s;%s;%d", date, ket, profit)

	return result // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)

	// str := "1/1/2021;INCOME;100000"
	// tes := strings.Split(str, ";")
	// fmt.Println(tes)
}
