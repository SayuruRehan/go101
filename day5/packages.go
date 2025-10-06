package main

import (
	"fmt"

	"./util"
)

// Importing a local package named 'util' which contains the Factorial function.
func main() {
	fact := util.Factorial(5)
	fmt.Println("Factorial of 5 is:", fact)
}
