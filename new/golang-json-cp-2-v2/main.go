package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type LoanData struct {
	StartBalance int
	Data         []Loan
	Employees    []Employee
}

type Loan struct {
	Date        string
	EmployeeIDs []string
}

type Employee struct {
	ID       string
	Name     string
	Position string
}

// json structure
type LoanRecord struct {
	MonthDate    string     `json:"month_date"`
	StartBalance int        `json:"start_balance"`
	EndBalance   int        `json:"end_balance"`
	Borrowers    []Borrower `json:"borrowers"`
}

type Borrower struct {
	Identity string `json:"id"`
	Total    int    `json:"total_loan"`
}

func LoanReport(data LoanData) LoanRecord {
	var rawmap = map[string]int{}
	result := LoanRecord{}
	var startbal = data.StartBalance
	var temp int
	var res1 = []Borrower{}
	var res2 = Borrower{}
	for _, tanggal := range data.Data {
		splitter := strings.Split(tanggal.Date, "-")
		result.MonthDate = splitter[1] + " " + splitter[2]
		for _, res := range tanggal.EmployeeIDs {
			if startbal >= 50000 {
				rawmap[res] += 50000
				startbal -= 50000
				temp += 50000
			} else {
				rawmap[res] += startbal
				temp += startbal
			}
		}
	}

	for key, value := range rawmap {
		res2.Identity = key // tanggal
		res2.Total = value  // startbal
		res1 = append(res1, res2)
	}
	sort.Slice(res1, func(i, j int) bool { return res1[i].Identity < res1[j].Identity })
	result.StartBalance = data.StartBalance
	result.EndBalance = data.StartBalance - temp
	result.Borrowers = res1
	return result // TODO: replace this
}

func RecordJSON(record LoanRecord, path string) error {
	jsonData, err := json.Marshal(record)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

// gunakan untuk debug
func main() {
	records := LoanReport(LoanData{
		StartBalance: 500000,
		Data: []Loan{
			{"01-January-2021", []string{"1", "2"}},
			{"02-January-2021", []string{"1", "2", "3"}},
			{"03-January-2021", []string{"2", "3"}},
			{"04-January-2021", []string{"1", "3"}},
		},
		Employees: []Employee{
			{"1", "Employee A", "Manager"},
			{"2", "Employee B", "Staff"},
			{"3", "Employee C", "Staff"},
		},
	})

	err := RecordJSON(records, "loan-records.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)
}