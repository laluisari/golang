package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	money := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1} // tentuang
	totalPrice := 0
	for _, listProduct := range products {
		totalPrice += listProduct.Price + listProduct.Tax
	}
	totalChanges := amount - totalPrice //uang - total harga
	fmt.Println(totalChanges)
	changesMoney := make([]int, 0)

	for _, m := range money {
		totalMoney := totalChanges / m // total harga dibagi element money
		fmt.Print(totalMoney)
		for a := 1; a <= totalMoney; a++ { //di input berapa kali element money
			changesMoney = append(changesMoney, m) //diinputin
		}
		totalChanges = totalChanges % m //total tadi kita ambil sisanya

	}

	return changesMoney // TODO: replace this
}

// func calculateTotalPrice(products []Product) int {
// 	total := 0
// 	for _, listProduct := range products {
// 		total += listProduct.Price + listProduct.Tax
// 	}
// 	return total
// }

func main() {
	fmt.Println(MoneyChanges(10000, []Product{{Name: "Baju", Price: 5000, Tax: 500}, {Name: "Celana", Price: 3000, Tax: 300}}))
}
