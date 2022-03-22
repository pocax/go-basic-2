package main

import "fmt"

func main() {
	// var person map[string]string //deklarasi map
	// person = map[string]string{} //inisialisasi map

	// person["name"] = "Hactiv"
	// person["age"] = "6"
	// person["address"] = "Jl. Kebon Jeruk"

	// fmt.Println("Name:", person["name"])
	// fmt.Println("Age:", person["age"])
	// fmt.Println("Address:", person["address"])

	//looping with map
	// var person = map[string]string{
	// 	"name":    "Hactiv",
	// 	"age":     "6",
	// 	"address": "Jl. Kebon Jeruk",
	// }

	// fmt.Println("Name:", person["name"])
	// fmt.Println("Age:", person["age"])
	// fmt.Println("Address:", person["address"])

	//other : looping with map
	// for key, value := range person {
	// 	fmt.Println(key, ":", value)
	// }

	//delete value
	// fmt.Println("Before Deleting", person)
	// delete(person, "age")
	// fmt.Println("After deleting : ", person)

	//detect value
	// value, exist := person["age"]
	// if exist {
	// 	fmt.Println("Value:", value)
	// } else {
	// 	fmt.Println("Value not exist")
	// }

	//delete value
	// delete(person, "age")
	// value, exist = person["age"]

	// if exist {
	// 	fmt.Println("Value:", value)
	// } else {
	// 	fmt.Println("Value not exist")
	// }

	//combine slice with map using range loop
	var people = []map[string]string{
		{"name": "Hactiv", "age": "6"},
		{"name": "Kidz", "age": "2"},
		{"name": "Code", "age": "5"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}
