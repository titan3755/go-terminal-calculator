package main

import (
	"fmt"
	"os"
)

func initCalc() {
	fmt.Println("Welcome to Go terminal calculator!")
	fmt.Println("\n1. Add\n2.Subtract\n3.Multiply\n4.Divide")
}

func compute(mode int, numberOne float32, numberTwo float32) float32 {
	var result float32 
	if mode == 1 {
		result = numberOne + numberTwo
	} else if mode == 2 {
		result = numberOne - numberTwo
	} else if mode == 3 {
		result = numberOne * numberTwo
	} else if mode == 4 {
		result = numberOne / numberTwo
	} 
	return result
}

func mainCalc() {
	var (
		promptResponse string
		input string
	 	numOne float32
	 	numTwo float32
	)
	fmt.Print(">> ")
	fmt.Scan(&input)
	fmt.Print("\n NumOne: ")
	fmt.Scan(&numOne)
	fmt.Print("\n NumTwo: ")
	fmt.Scan(&numTwo)
	if input == "1" {
		fmt.Printf("\n Result of Addition: %v", compute(1, numOne, numTwo))
	} else if input == "2" {
		fmt.Printf("\n Result of Subtraction: %v", compute(2, numOne, numTwo))
	} else if input == "3" {
		fmt.Printf("\n Result of Multiplication: %v", compute(3, numOne, numTwo))
	} else if input == "4" {
		fmt.Printf("\n Result of Division: %v", compute(4, numOne, numTwo))
	} else {
		fmt.Printf("Invalid Inputs!")
	}
	fmt.Print("\nCalculate again? [Y/n]")
	fmt.Scan(&promptResponse)
	if promptResponse == "y" || promptResponse == "Y" {
		{}
	} else {
		os.Exit(0)
	}
}

func main() {
	initCalc()
	for {
		mainCalc()
	}
}

