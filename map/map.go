package main

import "fmt"

func main() {
	var chicken = map[string]int{}
	chicken["januari"] = 10

	var chicken4 = make(map[string]int)

	chicken4["Januari"] = 100
	chicken4["Februari"] = 101
	chicken4["Maret"] = 102
	chicken4["April"] = 102
	chicken4["Mei"] = 103
	chicken4["Juni"] = 104
	chicken4["Juli"] = 104
	chicken4["Agustus"] = 105
	chicken4["September"] = 106
	chicken4["Oktober"] = 107
	chicken4["November"] = 108
	chicken4["Desember"] = 109

	fmt.Println("Chicken ", chicken["januari"])
	fmt.Println("Chicken4 ", chicken4["Januari"])

	delete(chicken4, "Januari")

	for index, value := range chicken4 {
		fmt.Println("Chicken 4 Key ", index)
		fmt.Println("Chicken 4 ", value)
	}

	value, isExist := chicken4["Januari"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item not exist")
	}

	var students = []map[string]string{
		{"name": "Jhon", "gender": "Male"},
		{"name": "Eko", "gender": "Male"},
		{"name": "Prili", "gender": "Famale"},
	}

	for _, c := range students {
		fmt.Printf("\n\tName Student %v Gender %v\n", c["name"], c["gender"])
	}

}
