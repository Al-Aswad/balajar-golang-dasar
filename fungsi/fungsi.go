package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"Jhon", "Eko"}
	printMessage("Halo", names)

	rand.Seed(time.Now().Unix())

	var randomValue int

	randomValue = randomWithRange(1, 10)
	fmt.Println("Random Value ", randomValue)
	randomValue = randomWithRange(1, 10)
	fmt.Println("Random Value ", randomValue)
	randomValue = randomWithRange(1, 10)
	fmt.Println("Random Value ", randomValue)

}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, "-")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	return rand.Int()%(max-min+1) + min
}
