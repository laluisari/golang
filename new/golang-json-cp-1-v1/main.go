package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
	// TODO: answer here
}

type Study struct {
	StudyName   string `json:"study_name"`
	StudyCredit int    `json:"study_credit"`
	Grade       string `json:"grade"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return Report{}, err
	}
	var study Report
	err = json.Unmarshal(jsonData, &study)
	if err != nil {
		return Report{}, err
	}

	return study, nil
}

func GradePoint(report Report) float64 {
	var nilai float64
	var credit int
	var ipk float64

	for _, val := range report.Studies {
		if val.Grade == "A" {
			nilai += 4 * float64(val.StudyCredit)
		} else if val.Grade == "AB" {
			nilai += 3.5 * float64(val.StudyCredit)
		} else if val.Grade == "B" {
			nilai += 3 * float64(val.StudyCredit)
		} else if val.Grade == "BC" {
			nilai += 2.5 * float64(val.StudyCredit)
		} else if val.Grade == "C" {
			nilai += 2 * float64(val.StudyCredit)
		} else if val.Grade == "CD" {
			nilai += 1.5 * float64(val.StudyCredit)
		} else if val.Grade == "D" {
			nilai += 1 * float64(val.StudyCredit)
		} else if val.Grade == "DE" {
			nilai += 0.5 * float64(val.StudyCredit)
		} else {
			nilai += 0
		}
		credit += val.StudyCredit
	}
	if len(report.Studies) > 0 {
		ipk = nilai / float64(credit)
		return ipk
	} else {
		return 0.0
	}

	// TODO: replace this
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
