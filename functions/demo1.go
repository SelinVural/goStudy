package functions

import "fmt"

func SelamVer(kullaniciAdi string) {
	fmt.Println("merhaba", kullaniciAdi)
}

func Topla(sayi1 int, sayi2 int) int {
	var toplam = sayi1 + sayi2
	return toplam

}
