package main

import (
	"golesson/slices"
)

func main() {

	//functions.SelamVer("Selin")
	//var sonuc = functions.Topla(2, 6)
	//fmt.Println(sonuc * 10)

	//sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 6)

	//fmt.Println("Toplam:", sonuc1)
	//fmt.Println("Çıkarım:", sonuc2)
	//fmt.Println("Çarpım:", sonuc3)
	//fmt.Println("Bölüm:", sonuc4)

	// fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))
	// // fmt.Println(functions.ToplaVariadic(1, 4))
	// fmt.Println(functions.ToplaVariadic())

	// sayilar := []int{4, 6, 7, 0, 11}
	// fmt.Println(functions.ToplaVariadic(sayilar...))

	//string_functions.Demo2()
	//sayilar := []int{1, 2, 3}
	//pointers.Demo2(sayilar)
	//fmt.Println("maindeki sayı:", sayilar[0])

	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)

	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	// ciftSayiToplam, testekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	// carpim := ciftSayiToplam * testekSayiToplam
	// fmt.Println("çarpım:", carpim)

	//time.Sleep(5 * time.Second)
	//fmt.Println("main bitti")

	// interfaces.Demo2()

	slices.Demo3()

}
