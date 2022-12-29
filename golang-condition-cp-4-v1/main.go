package main

func GetTicketPrice(vip int, regular int, student int, day int) float32 {
	checkTgl := day % 2
	jmlhTiket := vip + regular + student
	hargaTotal := (vip * 30) + (regular * 20) + (student * 10)

	if hargaTotal >= 100 {
		if checkTgl != 0 {
			if jmlhTiket < 5 {
				harga := (hargaTotal) - (hargaTotal * 15 / 100)
				return float32(harga)
			} else {
				harga := (hargaTotal) - (hargaTotal * 25 / 100)
				return float32(harga)
			}
		} else {
			if jmlhTiket < 5 {
				harga := (hargaTotal) - (hargaTotal * 10 / 100)
				return float32(harga)
			} else {
				harga := (hargaTotal) - (hargaTotal * 20 / 100)
				return float32(harga)
			}
		}
	} else {
		return float32(hargaTotal)
	}

}
