package slices

import "fmt"

func Demo3() {
	myarray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := myarray[0:3]
	fmt.Println(cap(myslice))

}
