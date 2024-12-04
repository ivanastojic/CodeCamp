package main

import (
	"fmt"
)

func main() {
	proizvodi := map[string]int{
		"torba":      300,
		"maramice":   5,
		"prozor":     1500,
		"tipkovnica": 50,
	}

	najskupljiproizvod(proizvodi)
}

func najskupljiproizvod(proizvodi map[string]int) {

	var najskuplji string
	var max int
	for imeproizvoda, cijena := range proizvodi {
		if cijena > max {
			max = cijena
			najskuplji = imeproizvoda
		}
	}
	fmt.Printf("Najskuplji proizvod je <%s> s cijenom od <%d>.\n", najskuplji, max)

	fmt.Println("Proizvodi koji imaju cijenu veću od 150 su:")
	for imeproizvoda, cijena := range proizvodi {
		if cijena > 150 {
			fmt.Printf("%s: %d\n", imeproizvoda, cijena)
		}
	}
}
