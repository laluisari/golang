package main

import "fmt"

func ExchangeCoin(amount int) []int {
	money := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1} // tentuang

	changesMoney := make([]int, 0)
	for _, m := range money {
		totalMoney := amount / m // mengeloop berapa kali isi array digunakan

		for a := 1; a <= totalMoney; a++ {
			changesMoney = append(changesMoney, m)

		}
		amount = amount % m

	}

	return changesMoney // TODO: replace this
}
func main() {
	fmt.Println(ExchangeCoin(1752))
}
