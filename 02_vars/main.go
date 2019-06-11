package main

import "fmt"

var country = "Iran" //Global variable

func main() {
	//Main Data types
	//string
	//bool
	//int
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//byte - alias for uint8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128

	//using var
	var name string = "Abolfazl"
	//or you can just type "var name = "Abolfazl" it automatically knows
	//it is a string
	var name2 = "Sadeghian"
	var age int = 33
	const isCool = true

	// Shorthand variable declaration
	color := "Green"
	size := 1.3
	email := "email@gmail.com"

	// or you can declare multiple varibles in a single line like this
	name3, email2 := "Abolfazl2","email2@gmail.com"

	fmt.Println(name+" "+name2+" Age ", age, isCool)
	fmt.Printf("%T\n", isCool,color,size,email,name3,email2)

}
