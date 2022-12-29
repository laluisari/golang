package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	nilai := (math + science + english + indonesia) / 4
	if nilai >= 100 {

		return "Sempurna"
	} else if nilai >= 90 && nilai < 100 {

		return "Sangat Baik"

	} else if nilai >= 80 && nilai < 90 {

		return "Baik"

	} else if nilai >= 70 && nilai < 80 {

		return "Cukup"

	} else if nilai >= 60 && nilai < 70 {

		return "Kurang"

	} else {

		return "Sangat kurang"
	}
	// TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
