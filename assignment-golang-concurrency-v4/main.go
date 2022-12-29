package main

import (
	"errors"
	"fmt"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func GetTLD(domain string) (TLD string, IDN_TLD string) {
	var ListIDN_TLD = map[string]string{
		".com": ".co.id",
		".org": ".org.id",
		".gov": ".go.id",
	}

	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			TLD = domain[i:]
			break
		}
	}

	if _, ok := ListIDN_TLD[TLD]; ok {
		return TLD, ListIDN_TLD[TLD]
	} else {
		return TLD, TLD
	}
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	TLD, IDN_TLD := GetTLD(website.Domain)
	fmt.Println(TLD, IDN_TLD)
	if website.Domain == "" {
		chErr <- errors.New("domain name is empty")
		return
	} else if website.RefIPs == -1 {
		chErr <- errors.New("domain RefIPs not valid")
		return
	} else if !website.Valid {
		chErr <- errors.New("domain not valid")
		return
	}
	row := RowData{}
	row.RankWebsite = website.RankWebsite
	row.Domain = website.Domain
	row.TLD = TLD
	row.IDN_TLD = IDN_TLD
	row.RefIPs = website.RefIPs
	row.Valid = website.Valid

	ch <- row
	// TODO: replace this
}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)
	for _, website := range data {
		//run concurrency
		go FuncProcessGetTLD(website, ch, errCh)
	}

	tempData := []RowData{}
	for i := 0; i < len(data); i++ {
		select {
		case getTld := <-ch:
			ch <- getTld
		case err := <-errCh:
			return nil, err
		}

		rowData := <-ch
		newTld, _ := GetTLD(rowData.Domain)
		if TLD == newTld {
			ch <- rowData
			newGetTLD := <-ch
			tempData = append(tempData, newGetTLD)
		}
	}

	return tempData, nil
	// TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	rows, err := FilterAndFillData(".com", []RowData{
		{1, "google.com", "", "", true, 100},
		{2, "facebook.com", "", "", true, 100},
		{3, "golang.org", "", "", true, 100},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}
