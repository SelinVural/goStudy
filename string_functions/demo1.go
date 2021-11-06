package string_functions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Selin"
	fmt.Println(s.Count(isim, "l"))
	fmt.Println(s.Contains(isim, "b"))
	sonuc := s.Index(isim, "b")

	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
}
