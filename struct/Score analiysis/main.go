package main

import (
	"fmt"
	"sort"
)

func Add(a, b int) int {
	return 0
}

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	for a := 0; a < len(grades); a++ {
		s.Grades = append(s.Grades, grades[a])
	}

}

func Analysis(s School) (float64, int, int) {
	avg := 0.0
	count := 0.0
	iterate := 0
	if len(s.Grades) == 1 {
		avg = float64(s.Grades[0])
		min := int(s.Grades[0])
		max := int(s.Grades[0])
		return avg, min, max
	} else if len(s.Grades) > 1 {
		for a := 0; a < len(s.Grades); a++ { //kita loop, setelah itu kita bandingin dari 1 ke berikutnya
			count += float64(s.Grades[a])
			iterate = a
		}
		fmt.Println(count)
		avg = count / float64(iterate+1)
		sort.Ints(s.Grades)
		min := s.Grades[0]
		max := s.Grades[len(s.Grades)-1]

		return avg, min, max // TODO: replace this
	} else {
		avg = 0.0
		min := 0
		max := 0
		return avg, min, max
	}

}

// gunakan untuk melakukan debugging
func main() {
	//var sch School
	//sch.AddGrade(100, 90, 80, 70, 60, 60, 100, 100, 100, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57)

	avg, min, max := Analysis(School{
		Name:    "mody",
		Address: "selong",
		//Grades:  []int{100, 90, 80, 70, 60, 60, 100, 100, 100, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57},
		Grades: []int{100, 90, 95, 100},
	})
	fmt.Println(avg, min, max)

	// number := []int{14, 6, 9, 24}
	// sort.Ints(number)
	// fmt.Printf("Hasil sorting: %d\n", number)
	// s := sort.IntsAreSorted(number)
	// fmt.Println("Sudah di sorting: ", s)

}
