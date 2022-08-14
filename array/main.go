package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "Eko"
	names[1] = "Jhon"
	names[2] = "Chena"
	names[3] = "Budi"

	var fruits = [4]string{"apple", "banana", "grape", "melon"}

	var number = [...]int{1, 2, 3, 4, 5, 6}

	var phone = make([]string, 1)
	phone[0] = "Samsung"

	fmt.Printf("Names %v\n", names)
	fmt.Printf("Fruits %v\n", fruits)
	fmt.Printf("Number %v\n", number)

	for i, c := range fruits {
		fmt.Printf("Fruit Index %v %v\n", i, c)
	}

	for i := range fruits {
		fmt.Printf("Fruit Index %v\n", i)
	}
	fmt.Printf("Phone %v\n", phone)
}
