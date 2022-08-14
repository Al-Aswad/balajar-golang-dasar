package main

import "fmt"

func main() {
	var fruits = []string{"banana", "melon", "apple"}
	fruits = append(fruits, "banana")

	fmt.Println("Fruits ", fruits)
}
