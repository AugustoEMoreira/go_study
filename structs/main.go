package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}
	joao := person{firstName: "Joao", lastName: "Pereira"}
	var vitor person
	vitor.firstName = "Vitor"
	vitor.lastName = "Pereira"
	fmt.Println(joao)
	fmt.Println(alex)
	fmt.Printf("%+v", vitor)
}
