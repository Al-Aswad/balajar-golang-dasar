package main

import "fmt"

func main() {
	var fruitsA = []string{"banana", "apple", "grape"}           //slice
	var fruitsB = [4]string{"banana", "apple", "grape", "melon"} //array
	var fruitsC = [...]string{"banana", "apple", "grape"}        //array

	var newFruits = fruitsA[0:2] //mengambil nilai mulai dari index ke o sampai nilai sebelum index ke 2
	var newFruitsA = fruitsB[1:3]

	newFruits[0] = "melon" //banana to melon, dan slice lama akan ikut berubah

	fmt.Println("Fruit A ", fruitsA)
	fmt.Println("Fruit B ", fruitsB)
	fmt.Println("Fruit C ", fruitsC)
	fmt.Println("New Fruits", newFruits)
	fmt.Println("New Fruits", newFruitsA)
	fmt.Println("New Fruits", fruitsA)

	fmt.Println("Len Fruits A ", len(fruitsA))
	fmt.Println("Len New Fruits A ", len(newFruitsA))
	fmt.Println("Cap Fruits A ", cap(fruitsA))
	fmt.Println("Cap New Fruits A ", cap(newFruitsA))

	fmt.Println("Len Fruits B ", len(fruitsB))
	fmt.Println("Len New Fruits B ", len(newFruitsA))
	fmt.Println("Cap Fruits B ", cap(fruitsB))
	fmt.Println("Cap New Fruits B ", cap(newFruitsA))
}
