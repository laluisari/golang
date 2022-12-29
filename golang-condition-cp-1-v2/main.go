package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	if score >= 70 && absent < 5 {
		return "lulus"
		// TODO: replace this
	} else if score < 70 || absent >= 5 {
		return "tidak lulus"
	} else {
		return ""
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(100, 4))
	// fmt.Println(GraduateStudent(80, 5))

}
