package main

import "fmt"

func main() {
	//var fruits = make([]string, 3)
	// _ = fruits

	// fruits[0] = "Apple"
	// fruits[1] = "Orange"
	// fruits[2] = "Banana"

	//append
	// fruits = append(fruits, "Apple", "Orange", "Banana")
	// fmt.Printf("%#v", fruits)

	// var fruits1 = []string{"Apple", "Orange", "Banana"}
	// var fruits2 = []string{"Mango", "Grape", "Pineapple"}

	// fruits1 = append(fruits1, fruits2...)

	// fmt.Printf("%#v", fruits1)

	//function copy slice
	// var fruits1 = []string{"Apple", "Orange", "Banana"}
	// var fruits2 = []string{"Mango", "Grape", "Pineapple"}

	// nn := copy(fruits1, fruits2)

	// fmt.Println("Fruits1 =>", fruits1)
	// fmt.Println("Fruits2 =>", fruits2)
	// fmt.Println("Total copy =>", nn)

	//slicing
	// var fruits1 = []string{"Apple", "Orange", "Banana", "Mango", "Grape", "Pineapple"}

	// var fruits2 = fruits1[1:4]
	// fmt.Printf("%#v\n", fruits2)

	// var fruits3 = fruits1[0:]
	// fmt.Printf("%#v\n", fruits3)

	// var fruits4 = fruits1[:3]
	// fmt.Printf("%#v\n", fruits4)

	// var fruits5 = fruits1[:]
	// fmt.Printf("%#v\n", fruits5)

	//slicing append
	// var fruits1 = []string{"Apple", "Orange", "Banana", "Mango", "Grape", "Pineapple"}

	// fruits1 = append(fruits1[3:], "Orange")

	// fmt.Printf("%#v\n", fruits1)

	//slice backing array
	// var fruits1 = []string{"Apple", "Orange", "Banana", "Mango", "Grape", "Pineapple"}
	// var fruits2 = fruits1[2:4]

	// fruits2[0] = "Banana"

	// fmt.Println("Fruits1 =>", fruits1)
	// fmt.Println("Fruits2 =>", fruits2)

	//slice cap
	// var fruits1 = []string{"Apple", "Orange", "Banana", "Mango", "Grape", "Pineapple"}

	// fmt.Println("Fruits1 Cap :", cap(fruits1))
	// fmt.Println("Fruits1 Len :", len(fruits1))

	// fmt.Println(strings.Repeat("#", 20))

	// var fruits2 = fruits1[0:3]

	// fmt.Println("Fruits2 Cap :", cap(fruits2))
	// fmt.Println("Fruits2 Len :", len(fruits2))

	// fmt.Println(strings.Repeat("#", 20))

	// var fruits3 = fruits1[1:]

	// fmt.Println("Fruits3 Cap :", cap(fruits3))
	// fmt.Println("Fruits3 Len :", len(fruits3))

	//slice new backing array

	cars := []string{"Ferrari", "Mercedes", "BMW", "Audi", "Porsche", "Lamborghini"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Mazda"
	fmt.Println("cars :", cars)
	fmt.Println("newCars :", newCars)
}
