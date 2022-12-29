package invoice

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Marketing invoice
type MarketingInvoice struct {
	Date        string
	StartDate   string
	EndDate     string
	PricePerDay int
	AnotherFee  int
	Approved    bool
}

func (mi MarketingInvoice) RecordInvoice() (InvoiceData, error) {

	if mi.Date == "" {
		return InvoiceData{}, errors.New("invoice date is empty") // TODO: replace this
	}
	if mi.StartDate == "" || mi.EndDate == "" {
		return InvoiceData{}, errors.New("travel date is empty")
	}
	if mi.PricePerDay <= 0 {
		return InvoiceData{}, errors.New("price per day is not valid")
	}

	date := ChangeDate(mi.Date)
	sd, sm, year := SplitFormat(mi.StartDate)
	ed, em, eyear := SplitFormat(mi.EndDate)
	t1 := CDate(sd, sm, year)
	t2 := CDate(ed, em, eyear)

	days := (t2.Sub(t1).Hours() / 24) + 1

	totalInvoice := days*float64(mi.PricePerDay) + float64(mi.AnotherFee)

	return InvoiceData{
		Date:         date,
		TotalInvoice: totalInvoice,
		Departemen:   Marketing,
	}, nil // TODO: replace this
}

func CDate(year, month, day int) time.Time {

	return time.Date(day, time.Month(month), year, 0, 0, 0, 0, time.UTC)
}

func SplitFormat(dataSlice string) (sd int, sm int, year int) {
	res := strings.Split(dataSlice, "/")
	sd, _ = strconv.Atoi(res[0])
	sm, _ = strconv.Atoi(res[1])
	year, _ = strconv.Atoi(res[2])
	return
}
