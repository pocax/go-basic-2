package main

import "fmt"

func main() {
	// var numbers [4]int
	// numbers = [4]int{1, 2, 3, 4}

	// var strings = [3]string{"Hello", "World", "Go"}

	// fmt.Printf("%#v\n", numbers)
	// fmt.Printf("%#v\n", strings)

	// var fruits = [4]string{"Apple", "Orange", "Banana", "Mango"}
	// fruits[0] = "Apple"
	// fruits[1] = "Orange"
	// fruits[2] = "Banana"
	// fruits[3] = "Mango"

	// fmt.Printf("%#v\n", fruits)

	// var fruits = [4]string{"Apple", "Orange", "Banana", "Mango"}

	// //cara pertama
	// for i, v := range fruits {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, v)
	// }

	// fmt.Println(strings.Repeat("#", 20))

	// //care kedua
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	// }

	balance := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balance {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}
