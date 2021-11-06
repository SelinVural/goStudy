package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("ürün kaydedildi:", p.productName)
	defer Log()

}

func Log() {
	fmt.Println("loglandı")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}
	//defer p.save()
	// deferden önce laptop olduğu için output laptop olur. deferü mouseun altına yazsaydık sonuç mouse olurdu.
	p = product{productName: "Mouse", unitPrice: 150}

	fmt.Println("işlem başarılı")
	fmt.Println(p.productName)
	defer p.save()

}
