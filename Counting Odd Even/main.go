package main

import "fmt"

func CountingNumber(n int) float64 {
	rest := 0.0
	a := float64(n)
	for temp := 1.0; temp <= a; temp += 0.5 {
		rest += temp
	}
	return rest // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(5))
}
