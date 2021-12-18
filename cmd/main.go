package main

import (
	"fmt"
	"struct-go/internal/adapters/core/arithmetic"
)

func main() {
	// var core ports.ArithmeticPort
	// core = arithmetic.NewAdapter()

	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(1, 3)
	if err != nil {
		fmt.Println("pipec")
		return
	}

	fmt.Println(result)
}
