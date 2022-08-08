package main

import "fmt"

func main() {
	var firstName string = "Jhon"
	lastName := "Chena"
	var first, second, third string = "1", " 2", "3"
	four, five, six := 4, 5, "6"

	name := new(string)

	fmt.Printf("My name is %s %s \n", firstName, lastName)
	fmt.Printf("\tMy name is %s %s", firstName, lastName)

	fmt.Printf("\n%s,%s,%s", first, second, third) //deklarasi multi variable
	fmt.Printf("\n%v,%v,%v", four, five, six)      //deklarasi multi variable

	fmt.Printf("\n %v", name)  //alamat memori
	fmt.Printf("\n %v", *name) //alamat value

}
