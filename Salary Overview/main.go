package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi ini akan mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
func FormatRupiah(number int) string {
	var result string

	numStr := strconv.Itoa(number)

	for i := len(numStr) - 1; i >= 0; i-- {
		result = string(numStr[i]) + result
		if (len(numStr)-i)%3 == 0 && i != 0 {
			result = "." + result
		}
	}

	return "Rp " + result
}

func Date(year, month, day int) time.Time {
	return time.Date(day, time.Month(month), year, 0, 0, 0, 0, time.UTC)
}

func SplitFormat(dataSlice string) (sd int, sm string, ed int, em string, year int) {
	res := strings.Split(dataSlice, " ")
	sd, _ = strconv.Atoi(res[0])
	sm = res[1]
	ed, _ = strconv.Atoi(res[3])
	em = res[4]
	year, _ = strconv.Atoi(res[5])
	return
}
func GetDayDifference(date string) int {
	var result int
	bulan := map[string]int{
		"January":   1, //abcdefghijklmnopqrstuvwxyz
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	sd, sm, ed, em, year := SplitFormat(date)

	if bulan[sm] > bulan[em] {

		fmt.Println("inputan bulan tidak valid")
	}

	checkSm := Date(0, bulan[sm]+1, year)
	// checkSd := Date(sd, bulan[sm]+1, year)
	// checkEd := Date(ed, bulan[sm]+1, year)

	if checkSm.Day() > 30 {
		last := Date(0, bulan[em]+1, year)
		fmt.Println(last)
		if ed >= last.Day() {
			fmt.Println("tanggal melebihi batas")
		} else {

			t1 := Date(sd, bulan[sm], year)
			t2 := Date(ed, bulan[em], year)
			days := (t2.Sub(t1).Hours() / 24)

			if int(days) == 0 {
				result = int(days) + 1
			} else {
				result = int(days)
			}

		}

	} else {
		last := Date(0, bulan[em]+1, year)
		if ed >= last.Day() {
			fmt.Println("tanggal melebihi batas")
		} else {
			t1 := Date(sd, bulan[sm], year)
			t2 := Date(ed, bulan[em], year)
			days := (t2.Sub(t1).Hours() / 24) + 1
			result = int(days)
		}
	}

	return result // TODO: replace this
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	duplicate := make(map[string]int)
	result := make(map[string]string)
	for a := 0; a < rangeDay; a++ {
		for b := 0; b < len(data[a]); b++ {
			nama := data[a][b] // ini adalah data nama
			_, exist := duplicate[nama]
			if exist {
				duplicate[nama] += 50000
				rp := FormatRupiah(duplicate[nama])
				result[nama] = rp
			} else {
				duplicate[nama] = 50000
				rp := FormatRupiah(duplicate[nama])
				result[nama] = rp
			}
		}

	}

	return result // TODO: replace this
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	day := GetDayDifference(dateRange)
	if day > len(data) {
		result := GetSalary(len(data), data)
		return result

	} else {
		result := GetSalary(day, data)
		return result
	}

	// TODO: replace this
}

func main() {

	res := GetSalaryOverview("21 February - 22 February 2021", [][]string{
		{"Andi", "Imam", "Eddy", "Deny"},
		{"Andi", "Imam"},
		{"Imam", "Deny"},
		{"Andi", "Deny"},
	})
	fmt.Println(res)
}
