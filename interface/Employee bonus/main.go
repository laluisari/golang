package main

import "fmt"

type Employee interface { //ine interface na
	GetBonus() float64
}

type Junior struct { //ine data struc si kadunte temang ipak func
	Name         string
	BaseSalary   int
	WorkingMonth int
}
type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (e Junior) GetBonus() float64 { //ine func na
	return float64(e.BaseSalary) * GetProrate(e.WorkingMonth)
}
func (e Senior) GetBonus() float64 {
	return 2*float64(e.BaseSalary)*GetProrate(e.WorkingMonth) + GetPerformance(e.BaseSalary, e.PerformanceRate)
}
func (e Manager) GetBonus() float64 {
	return 2*float64(e.BaseSalary)*GetProrate(e.WorkingMonth) + GetPerformance(e.BaseSalary, e.PerformanceRate) + (e.BonusManagerRate * float64(e.BaseSalary))
}

func GetPerformance(baseSalary int, performance float64) float64 {
	return performance * float64(baseSalary)
}
func GetProrate(WorkingMonth int) float64 {
	if WorkingMonth > 12 {
		return 1.0
	}
	return float64(WorkingMonth) / 12.0
}
func EmployeeBonus(employee Employee) float64 { //ine call ne
	result := employee.GetBonus()
	return result // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	result := 0.0
	for _, check := range employees {
		result += check.GetBonus()
	}
	return result // TODO: replace this
}

func main() {
	junior := Junior{Name: "isa", BaseSalary: 100000, WorkingMonth: 12}
	senior := Senior{Name: "isa", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5}
	manager := Manager{Name: "isa", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5}
	fmt.Println(EmployeeBonus(junior))
	fmt.Println(EmployeeBonus(senior))
	fmt.Println(EmployeeBonus(manager))
	slice := make([]Employee, 0)
	slice = append(slice, junior, manager)
	fmt.Println(TotalEmployeeBonus(slice))
}
