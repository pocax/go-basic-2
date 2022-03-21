package main

import "fmt"

func main() {
	// var currentYear = 2022

	// if age := currentYear - 1993; age < 17 {
	// 	fmt.Println("Kamu belum boleh membuat kartu SIM")
	// } else {
	// 	fmt.Println("Kamu boleh membuat kartu SIM")
	// }
	// var score = 8

	// switch score {
	// case 8:
	// 	fmt.Println("Perfect")
	// case 7:
	// 	fmt.Println("Good")
	// case 6:
	// 	fmt.Println("Ok")
	// }

	// var score = 6

	// switch {
	// case score == 8:
	// 	fmt.Println("Perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("not bad")
	// default:
	// 	{
	// 		fmt.Println("study harder")
	// 		fmt.Println("you need to learn more")
	// 	}
	// }

	// var score = 9

	// switch {
	// case score == 8:
	// 	fmt.Println("Perfect")
	// case score < 8 && score > 3:
	// 	fmt.Println("not bad")
	// 	fallthrough
	// case score < 5:
	// 	fmt.Println("its ok but please study harder")
	// default:
	// 	{
	// 		fmt.Println("study harder")
	// 		fmt.Println("You dont have a good score yet")
	// 	}
	// }

	var score = 10

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice")
		}
	} else {
		if score == 5 {
			fmt.Println("Not bad")
		} else if score == 3 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("You can do it")
			if score == 0 {
				fmt.Println("try harder")
			}
		}
	}
}
