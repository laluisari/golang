package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	strCount := CountTransaction(transactions)

	err := os.WriteFile(path, []byte(strCount), 0644)

	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func CountTransaction(transactions []Transaction) string {
	//masukkan nanti ke map dari elemen struct
	var recapData = make(map[string]int)

	for a := 0; a < len(transactions); a++ {
		if transactions[a].Type == "expense" {
			//kata kunci nya diberi nilai
			recapData[transactions[a].Date] -= transactions[a].Amount
		} else {
			recapData[transactions[a].Date] += transactions[a].Amount
		}
	}

	//kemudian data di atas di tampung di array
	var dates []string
	//kita loop recap datanya saa
	for date := range recapData {
		dates = append(dates, date)
	}

	sort.Strings(dates)

	var str string

	for _, d := range dates {
		if recapData[d] < 0 {
			totalExp := int(math.Abs(float64(recapData[d])))
			str += fmt.Sprintf("%s;expense;%d\n", d, totalExp)
		} else {
			str += fmt.Sprintf("%s;income;%d\n", d, recapData[d])
		}
	}

	if len(str) > 0 {
		return str[:len(str)-1]
	} else {
		return str
	}

}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"01/01/2021", "expense", 30000},
		{"01/01/2021", "income", 20000},
		{"02/01/2021", "expense", 10000},
		{"02/01/2021", "expense", 10000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
