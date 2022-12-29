package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	temp := make(map[string]int)
	rawOutput := []string{}
	output := ""
	var err error

	for i := 0; i < len(transactions); i++ {
		if transactions[i].Type == "income" {
			temp[transactions[i].Date] += transactions[i].Amount
		} else {
			temp[transactions[i].Date] -= transactions[i].Amount
		}
	}
	// fmt.Println(temp)

	for key, val := range temp {
		status := ""
		if val < 0 {
			status = "expense"
			val = -val
		} else {
			status = "income"
		}
		store := key + ";" + status + ";" + strconv.Itoa(val)
		rawOutput = append(rawOutput, store)
	}

	sort.Strings(rawOutput)
	for i, val := range rawOutput {
		if i == len(rawOutput)-1 {
			output += val
		} else {
			output += val + "\n"
		}
	}

	err = os.WriteFile(path, []byte(output), 0644)
	if err != nil {
		return err
	}

	// fmt.Print(output)
	return err // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"01/01/2021", "expense", 30000},
		{"01/01/2021", "income", 20000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
