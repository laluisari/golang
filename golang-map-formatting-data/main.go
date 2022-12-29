package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	slice := make(map[string][]string)
	for _, dataSlice := range data {
		category, index, position, value := SplitFormat(dataSlice)
		_, exists := slice[category]

		if !exists {
			slice[category] = make([]string, 0)
		}

		//kemudian kita check berapa banyak indexnya

		if len(slice[category]) > index {
			if position {
				slice[category][index] = value
			} else {
				slice[category][index] = strings.Join([]string{slice[category][index], value}, " ")
			}
		} else {
			slice[category] = append(slice[category], value)
		}

	}
	return slice // TODO: replace this
}

func SplitFormat(dataSlice string) (category string, index int, position bool, value string) {
	res := strings.Split(dataSlice, "-")
	category = res[0]
	index, _ = strconv.Atoi(res[1]) //int
	position = res[2] == "first"
	value = res[3]
	return
}

func main() {
	list := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar"}
	fmt.Println(ChangeOutput(list))

}

/*

tempSlice := []string{}
	tempStr := ""
	x := make(map[string][]string)
	for _, words := range data {
		tempStr = words
		tempSlice = strings.Split(tempStr, "-")
		header := tempSlice[0]
		for a := 0; a < 1; a++ {
			x[header] = append(x[header], tempSlice[3])
		}

	}


#pikirin kalau kamu akan memasukkannya ke dalam map
#ide
- ambil index slice, kemudian convert to string
setelah itu check contains nya apakah terdapat, jika ada
masukin

oke diketahui kita disuruh memperbaikki format data


data saat ini [header]-[index]-[position]-[value]


tugasnya
mengubah format data menjadi sebuah map dengan tipe data
map[string][]string

okey pertama kita ambil index dari slice, jadi string
kemudian kita pecah string tersebut jadi array
nah setelah kita pecah jadi array,
aku ingin memasukkan nya ke dalam map , jadi nanti setiap map kita


**/
