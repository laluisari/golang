package invoice

import "errors"

// Finance invoice

type FinanceInvoice struct {
	Date     string
	Status   InvoiceStatus // status: "paid", "unpaid"
	Approved bool
	Details  []Detail
}

type InvoiceStatus string

const (
	PAID   InvoiceStatus = "paid"
	UNPAID InvoiceStatus = "unpaid"
)

type Detail struct {
	Description string
	Total       int
}

func (fi FinanceInvoice) RecordInvoice() (InvoiceData, error) { //disini w logic bisnis
	if fi.Date == "" {
		return InvoiceData{}, errors.New("invoice date is empty") // TODO: replace this
	}
	if fi.Status != PAID && fi.Status != UNPAID || fi.Status == "" {
		return InvoiceData{}, errors.New("invoice status is not valid")
	}

	if len(fi.Details) == 0 {
		return InvoiceData{}, errors.New("invoice details is empty")
	} else {
		for _, data := range fi.Details {
			if data.Total <= 0 {
				return InvoiceData{}, errors.New("total price is not valid")
			}
		}
	}
	FormatDate := ""

	FormatDate = ChangeDate(fi.Date)
	Rest := 0.0

	if fi.Approved == true && fi.Status == PAID {
		for _, total := range fi.Details { // w get
			Rest += float64(total.Total)
		}
	} else {
		Rest = Rest + 0.0
	}

	return InvoiceData{
		Date:         FormatDate,
		TotalInvoice: Rest,
		Departemen:   Finance,
	}, nil // TODO: replace this

}

//nanti disini akan ada func untuk mangggil interface nya
