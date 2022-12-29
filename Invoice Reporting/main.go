package main

import (
	"a21hc3NpZ25tZW50/invoice"
	"fmt"
	"log"
)

func RecapDataInvoice(data []invoice.Invoice) ([]invoice.InvoiceData, error) { //call func utama ne,  terus kelek so selapuk isi na
	total := 0.0                    //total
	tempDep := make([]string, 0)    //for name departemen
	tempDate := make([]string, 0)   //for date
	rest := []invoice.InvoiceData{} //for result

	//looping untuk ambil data buat check
	for a := 0; a < len(data); a++ {
		invoiceData, err := data[a].RecordInvoice()
		if err != nil {
			return rest, err
		}
		tempDate = append(tempDate, string(invoiceData.Date))     //tampung semua data date nya
		tempDep = append(tempDep, string(invoiceData.Departemen)) //tampung semua data date nya
	}

	//looping untuk membandingkan data yg buat check dengan data yg di invoiceData
	for a := 0; a < len(data); a++ {
		invoiceData, err := data[a].RecordInvoice()
		if err != nil {
			return rest, err
		}
		for b := 0; b < len(data); b++ {
			if string(invoiceData.Date) == tempDate[b] && string(invoiceData.Departemen) == tempDep[b] {
				if b <= 0 { //untuk ngisi data awal aja untuk bisa di check dengan data berikutnya

					total += invoiceData.TotalInvoice
					invoiceData.TotalInvoice = total
					tempDate[b] = ""
					tempDep[b] = ""
					rest = []invoice.InvoiceData{invoiceData}

				} else {
					//check apakah data sebelumnya yg sudah di masukkan ke rest sama atau tidak dengan data berikutnya
					//jika true

					if string(rest[len(rest)-1].Date) == tempDate[b] && string(rest[len(rest)-1].Departemen) == tempDep[b] {
						total += invoiceData.TotalInvoice
						invoiceData.TotalInvoice = total
						tempDate[b] = ""
						tempDep[b] = ""
						rest = []invoice.InvoiceData{invoiceData}

						//jika departemennya berbeda maka tinggal tambahin aja dengan total invoice departemen ini sendiri

					} else if string(rest[len(rest)-1].Date) == tempDate[b] && string(rest[len(rest)-1].Departemen) != tempDep[b] {
						total2 := 0.0
						total2 += invoiceData.TotalInvoice
						invoiceData.TotalInvoice = total2
						tempDate[b] = ""
						tempDep[b] = ""
						rest = append(rest, invoiceData)
					} else if string(rest[len(rest)-1].Date) != tempDate[b] && string(rest[len(rest)-1].Departemen) != tempDep[b] {
						total2 := 0.0
						total2 += invoiceData.TotalInvoice
						invoiceData.TotalInvoice = total2
						tempDate[b] = ""
						tempDep[b] = ""
						rest = append(rest, invoiceData)
					}

				}
			}
		}
	}

	return rest, nil

}

func main() {
	listInvoice := []invoice.Invoice{
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{"pembelian nota", 4000}, {"pembelian kuota", 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},

		invoice.FinanceInvoice{
			Date:     "02/02/2020",
			Details:  []invoice.Detail{{"pembelian nota", 4000}, {"pembelian kuota", 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},

		invoice.WarehouseInvoice{
			Date: "03-February-2020",
			Products: []invoice.Product{
				{"product A", 10, 10000, 0.1},
				{"product C", 5, 15000, 0.2},
			},
			InvoiceType: invoice.PURCHASE,
			Approved:    true,
		},
		invoice.MarketingInvoice{
			Date:        "04/02/2020",
			StartDate:   "20/01/2022",
			EndDate:     "25/01/2022",
			Approved:    true,
			PricePerDay: 10000,
			AnotherFee:  5000,
		},
	}

	result, err := RecapDataInvoice(listInvoice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}

// if string(invoiceData.Date) == tempDate[b] && string(invoiceData.Departemen) == tempDep[b] {
// 	fmt.Println("same date and dep")
// 	if b <= 0 { //untuk ngisi data awal aja untuk bisa di check dengan data berikutnya
// 		total += invoiceData.TotalInvoice
// 		invoiceData.TotalInvoice = total
// 		tempDate[b] = ""
// 		tempDep[b] = ""
// 		rest = []invoice.InvoiceData{invoiceData}
// 	} else {
// 		//check apakah data sebelumnya yg sudah di masukkan ke rest sama atau tidak dengan data berikutnya
// 		//jika true
// 		if string(rest[len(rest)-1].Departemen) == tempDep[b] {
// 			total += invoiceData.TotalInvoice
// 			invoiceData.TotalInvoice = total
// 			tempDate[b] = ""
// 			tempDep[b] = ""
// 			rest = []invoice.InvoiceData{invoiceData}

// 			//jika departemennya berbeda maka tinggal tambahin aja dengan total invoice departemen ini sendiri
// 		} else if string(rest[len(rest)-1].Departemen) != tempDep[b] {
// 			total2 := 0.0
// 			total2 += invoiceData.TotalInvoice
// 			invoiceData.TotalInvoice = total2
// 			tempDate[b] = ""
// 			tempDep[b] = ""
// 			rest = append(rest, invoiceData)
// 		}

// 	}
// } else if string(invoiceData.Date) != tempDate[b] && string(invoiceData.Departemen) != tempDep[b] {
// 	fmt.Println("different date and dep")
// }

// invoice.FinanceInvoice{
// 	Date:     "04/02/2020",
// 	Details:  []invoice.Detail{{"pembelian nota", 4000}, {"pembelian kuota", 4000}},
// 	Status:   invoice.PAID,
// 	Approved: true,
// },

// invoice.FinanceInvoice{
// 	Date:     "04/02/2020",
// 	Details:  []invoice.Detail{{"pembelian nota", 4000}, {"pembelian kuota", 4000}},
// 	Status:   invoice.PAID,
// 	Approved: true,
// },

// invoice.WarehouseInvoice{
// 	Date: "04/02/2020",
// 	Products: []invoice.Product{
// 		{"product A", 10, 10000, 0.1},
// 		{"product C", 5, 15000, 0.2},
// 	},
// 	InvoiceType: invoice.PURCHASE,
// 	Approved:    true,
// },
// invoice.MarketingInvoice{
// 	Date:        "04/01/2022",
// 	StartDate:   "20/01/2022",
// 	EndDate:     "25/01/2022",
// 	Approved:    true,
// 	PricePerDay: 10000,
// 	AnotherFee:  5000,
// },
