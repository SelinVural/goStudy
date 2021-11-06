package loops

import "fmt"

func Demo4() {

	sayi := 0
	fmt.Println("bir sayı girin")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}
	if asalMi == true {
		fmt.Println("asal")

	} else {
		fmt.Println("asal değil")
	}

}
