package variables

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya!"
	fmt.Println(metin)

	var kdv int = 10
	fmt.Println(kdv)

	fmt.Println(100 + 100*10/100)

	kdv3 := 25
	fmt.Println(kdv3)

	var durum bool

	var metin1 string = "Engin"
	var metin2 string = "Ahmet"

	durum = metin1 != metin2

	fmt.Println(durum)
	fmt.Println(!durum)

}
