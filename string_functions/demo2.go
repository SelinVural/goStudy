package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Selin"
	fmt.Println(s.HasPrefix(isim, "Sel"))
	fmt.Println(s.HasSuffix(isim, "lin"))
	fmt.Println(s.Index(isim, "el"))
	harfler := []string{"s", "e", "l", "i", "n"}
	//	fmt.Println(s.Join(harfler, "-"))
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", 3))
	fmt.Println(s.Replace(sonuc, "*", "+", -1))

	fmt.Println(s.Split(sonuc, "-"))

	fmt.Println(s.Repeat(sonuc, 3))

}
