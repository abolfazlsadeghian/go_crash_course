package main

import "fmt"

func greeting (name string) string{
return "Hello "+ name
}

func getSum(number1 int, number2 int)int {
	return number1+number2
}

func main(){
	fmt.Println("Hello World")
	fmt.Println(greeting("Abolfazl"))
	fmt.Println(getSum(10,20))
}