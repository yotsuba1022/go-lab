package main

import (
	"fmt"
	"go-lab/pkg/calculator"
)

func main() {
	fmt.Println("Hello, World!")

	result := calculator.Add(5, 3)
	fmt.Printf("The result of adding 5 and 3 is: %d\n", result)
}
