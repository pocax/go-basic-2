package main

import "fmt"

func main() {
	//declare variable with uint8 data type
	var a uint8 = 10
	var b byte // this byte is an alias from uin8 data type

	b = a // assign value from a to b
	_ = b // ignore b

	//2nd example
	type second = uint
	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour) //hour type: uint
}
