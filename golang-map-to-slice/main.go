package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	slice := [][]string{}
	for key, value := range mapData {
		slice = append(slice, []string{key, value}) //setiap data kita bungkus jadi 1 slice

	}

	return slice
}

func main() {
	var data = map[string]string{
		"hello": "world",
		"John":  "Doe",
		"age":   "14",
	}
	fmt.Println(MapToSlice(data))
}
