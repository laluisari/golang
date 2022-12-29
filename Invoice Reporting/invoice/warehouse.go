package invoice

import (
	"errors"
	"strings"
)

// Warehouse invoice

type WarehouseInvoice struct {
	Date        string
	InvoiceType InvoiceTypeName
	Approved    bool
	Products    []Product
}

type InvoiceTypeName string

const (
	PURCHASE InvoiceTypeName = "purchase"
	SALES    InvoiceTypeName = "sales"
)

type Product struct {
	Name     string
	Unit     int
	Price    float64
	Discount float64
}

func (wi WarehouseInvoice) RecordInvoice() (InvoiceData, error) {
	FormatDate := ""
	if wi.Date == "" {
		return InvoiceData{}, errors.New("invoice date is empty")
	}
	if wi.InvoiceType != PURCHASE && wi.InvoiceType != SALES || wi.InvoiceType == "" {
		return InvoiceData{}, errors.New("invoice type is invalid")
	}
	if len(wi.Products) == 0 {
		return InvoiceData{}, errors.New("invoice products is empty")
	} else {
		for _, data := range wi.Products {
			if data.Unit <= 0 {
				return InvoiceData{}, errors.New("unit product is not valid")
			} else if data.Price <= 0 {
				return InvoiceData{}, errors.New("price product is not valid")
			}
		}
	}
	if strings.ContainsRune(wi.Date, '-') {
		FormatDate = wi.Date
	} else {
		FormatDate = ChangeDate(wi.Date)

	}
	totalInvoice := 0.0
	//RestTotal := 0.0
	for _, data := range wi.Products {
		Total := float64(data.Unit) * data.Price
		TotalDisc := Total * data.Discount
		totalInvoice += Total - TotalDisc

	}

	// unit * price - discout price
	return InvoiceData{
		Date:         FormatDate,
		TotalInvoice: totalInvoice,
		Departemen:   Warehouse,
	}, nil // TODO: replace this
}
