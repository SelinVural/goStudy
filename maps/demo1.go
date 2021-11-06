package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("eleman sayısı:", len(sozluk))
	fmt.Println("sözlük:", sozluk)
	delete(sozluk, "book")
	fmt.Println("eleman sayısı:", len(sozluk))
	fmt.Println("sözlük:", sozluk)

	deger, varMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu:", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)

}
