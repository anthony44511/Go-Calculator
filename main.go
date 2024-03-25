package main

import "fmt"

func main() {
	var number1, number2 float64
	var operation string

	fmt.Print("enter the first number: ")
	fmt.Scanln(&number1)
	fmt.Print("enter the second number: ")
	fmt.Scanln(&number2)

	fmt.Print("enter the operation (+, -, *, /): ")
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		fmt.Printf("%f %s %f = %f\n", number1, operation, number2, number1+number2)
	case "-":
		fmt.Printf("%f %s %f = %f\n", number1, operation, number2, number1+number2)
	}
}
