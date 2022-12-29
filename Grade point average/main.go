package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Study struct {
	Study_name   string `json:"study_name"`
	Study_credit int    `json:"study_credit"`
	Grade        string `json:"grade"`
}

type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

func ReadJSON(filename string) (Report, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Report{}, err
	}
	defer f.Close()

	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return Report{}, nil
	}

	var report Report

	err = json.Unmarshal(jsonData, &report)
	if err != nil {
		return Report{}, nil
	}
	return report, nil

}

func GradePoint(report Report) float64 {
	Alphabet := map[string]float64{
		"A":  4.0,
		"AB": 3.5,
		"B":  3.0,
		"BC": 2.5,
		"C":  2.0,
		"CD": 1.5,
		"D":  1.0,
		"DE": 0.5,
		"E":  0,
	}
	grade := 0.0
	sc := 0
	if len(report.Studies) <= 0 {
		return grade
	} else {
		for a := 0; a < len(report.Studies); a++ {
			for k, v := range Alphabet {
				if report.Studies[a].Grade == k {
					grade += v * float64(report.Studies[a].Study_credit)
					sc += report.Studies[a].Study_credit
				}
			}
		}
		grade = grade / float64(sc)

		return grade // TODO: replace this
	}
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
