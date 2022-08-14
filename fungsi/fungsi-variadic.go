package main

import "fmt"

func main() {
	sum := calculate("+", 1, 2, 3)
	fmt.Println("Hasil Perjumlahan ", sum)
	perkalian := calculate("*", 1, 2, 3)
	fmt.Println("Hasil Perkalian ", perkalian)
	pengurangan := calculate("-", 1, 2, 3)
	fmt.Println("Hasil pengurangan ", pengurangan)
}

func calculate(operator string, number ...int) int {
	var res int
	for _, c := range number {
		if operator == "-" {
			res -= c
		} else if operator == "+" {
			res += c
		} else if operator == ":" {
			res /= c
		} else if operator == "*" {
			res *= c
		} else {
			break
		}
	}

	return res
}
