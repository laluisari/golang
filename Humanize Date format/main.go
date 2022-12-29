package main

import "fmt"

func DateFormat(day int, month int, year int) string {

	date := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	if day < 10 {
		result := fmt.Sprint("0", day, "-", date[month], "-", year)
		return result
	} else {
		result := fmt.Sprint(day, "-", date[month], "-", year)
		return result
	}

}

func main() {
	fmt.Println(DateFormat(10, 1, 2020))
}
