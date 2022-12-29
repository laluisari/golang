package main

func TicketPlayground(height, age int) int {
	if age < 5 {
		return -1
	} else if age >= 5 && age <= 7 && height >= 120 {
		if height >= 135 && height < 150 {
			return 25000
		} else if height >= 150 && height < 160 {
			return 40000
		} else if height >= 160 {
			return 60000
		}
		return 15000
	} else if age >= 8 && age <= 9 && height >= 120 {
		if height >= 135 && height <= 149 {
			return 25000
		} else if height >= 150 && height <= 160 {
			return 40000
		} else if height >= 160 {
			return 60000
		}
		return 25000
	} else if age >= 10 && age <= 11 && height >= 120 {
		if height >= 135 && height <= 149 {
			return 40000
		} else if height >= 150 && height <= 159 {
			return 40000
		} else if height >= 160 {
			return 60000
		}
		return 40000
	} else if age == 12 && height >= 120 {
		return 60000
	} else {
		return 100000
	}

}

func main() {
	TicketPlayground(150, 12)
}
