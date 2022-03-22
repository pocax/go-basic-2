package main

import "fmt"

func main() {
	// city := "Jakarta"

	// for i := 0; i < len(city); i++ {
	// 	fmt.Printf("index: %d, byte: %d\n", i, city[i])
	// }

	//implement byte
	var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	fmt.Println(string(city))

	//implement rune by rune
	str1 := "Makan"
	str2 := "Manca"

	fmt.Printf("str1 byte lenght ==> %d\n", len(str1))
	fmt.Printf("str2 byte lenght ==> %d\n", len(str2))

	for i, s := range str1 {
		fmt.Printf("index: %d, rune: %s\n", i, string(s))
	}
}
