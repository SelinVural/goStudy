package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Print("Hesaptaki para yetersiz")
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("Paranız hazırlanıyor")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Println("Bitti. Hesaptaki Para:" + fmt.Sprintf("%v", hesap))

}
