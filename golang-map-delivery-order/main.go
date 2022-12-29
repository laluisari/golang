package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {

	lingkupLokasi := map[string]map[string]bool{
		"JKT": {"senin": true, "selasa": true, "rabu": true, "kamis": true, "jumat": true, "sabtu": true},
		"BDG": {"rabu": true, "kamis": true, "sabtu": true},
		"BKS": {"selasa": true, "kamis": true, "jumat": true},
		"DPK": {"senin": true, "selasa": true},
	}
	biayaAdmin := map[string]float32{
		"senin":  10.0,
		"selasa": 5.0,
		"rabu":   10.0,
		"kamis":  5.0,
		"jumat":  10.0,
		"sabtu":  5.0,
	}

	//make output
	result := make(map[string]float32)
	for a := 0; a < len(data); a++ {
		//sambung
		value := strings.Split(data[a], ":")
		FLname := value[0] + "-" + value[1]
		tempPrice, _ := strconv.Atoi(value[2]) //ambil harga
		price := float32(tempPrice)
		name := value[3]
		kodeLokasi := lingkupLokasi[name][day]

		if kodeLokasi {
			price = price + (price * (biayaAdmin[day] / 100.0))
			result[FLname] = price
		} else {
			log.Println("hari dan lokasi tidak sesuai")
		}

	}

	return result
	// TODO: replace this
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Andi:Sukirman:20000:JKT",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
