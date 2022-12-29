package main

import "fmt"

func EmailInfo(email string) string {
	var domain string
	var tld string
	for a := 0; a < len(email); a++ {
		if string(email[a]) == "@" {

			for b := a + 1; b < len(email); b++ {
				if string(email[b]) != "." {
					domain += string(email[b])
					if string(email[b+1]) == "." {
						for c := b + 2; c < len(email); c++ {
							tld += string(email[c])
						}
						b = len(email) - 1
					}
				}
			}

		}
	}
	// fmt.Println(domain)
	// fmt.Println(tld)
	result := "Domain: " + domain + " dan TLD: " + tld
	return result // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
