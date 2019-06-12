package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign of arrays short form
	nameArr:=[2]string{"Azita","Hoda"}

	//Declaring slices it is like ArrayList in Java it doesn't have a fixed size you can add to it later
	numberSlice:=[]int{23,12,344,231,654}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(nameArr)
	fmt.Println(numberSlice)
	//Getting the length of the Slice
	fmt.Println(len(numberSlice))

	//Checking the range in a slice
	fmt.Println(numberSlice[1:3])
}
