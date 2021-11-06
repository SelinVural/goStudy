package loops

import "fmt"

func Demo3() {
	tahminEdilenSayi := 0
	aklimdakiSayi := 80
	tahminSayisi := 0

	fmt.Println("Bir sayı tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi = tahminSayisi + 1

	for aklimdakiSayi != tahminEdilenSayi {
		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("daha büyük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}

		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("daha küçük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
	}

	basariDurumu := ""
	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Süper"

	} else if tahminSayisi <= 6 {
		basariDurumu = "Güzel"

	} else {
		basariDurumu = "Fena değil"
	}

	if tahminEdilenSayi == aklimdakiSayi {
		fmt.Printf("Bravo, bildiniz! %v tahmin %v", tahminSayisi, basariDurumu)
	}

}
