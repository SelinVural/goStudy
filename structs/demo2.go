package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {

	fmt.Println("eklendi:", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Selin", lastName: "Vural", age: 25}
	c.save()

}
