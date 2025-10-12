package main

import "fmt"

// Answer is just a struct used to demonstrate pointers
type Answer struct {
	value int
}

func main() {
	answer := Answer{value: 42}
	answerAgain := answer
	answer.value = 43

	fmt.Println("Wihout pointers, things get passed by value")

	fmt.Println(answer)      // output: {43}
	fmt.Println(answerAgain) // output: {42} (surprise, JS and Java folks!)

	fmt.Println("You want to pass by reference? Use pointers")

	pAnswer := &answer
	answer.value = 44
	fmt.Println(answer)   // output: {44}
	fmt.Println(*pAnswer) // output: {44}   (whew!)
}

/*
EXPLANATION:

1. Package Declaration: The code starts with the declaration of the main package, which is the entry point for a Go application.
2. Import Statement: The "fmt" package is imported to enable formatted I/O operations, such as printing to the console.
3. Struct Definition: A struct named `Answer` is defined with a single field `value` of type `int`. This struct is used to demonstrate how values and pointers work in Go.
4. Main Function: The `main` function is defined, which serves as the entry point for the program execution.
5. Struct Initialization: An instance of the `Answer` struct named `answer` is created and initialized with a value of 42.
6. Value Copying: A new variable `answerAgain` is created by copying the value of `answer`. This demonstrates that structs in Go are passed by value, meaning that `answerAgain` holds a separate copy of the data.
7. Value Modification: The `value` field of the `answer` struct is modified to 43. This change does not affect `answerAgain`, which still holds the original value of 42.
8. Printing Values: The values of `answer` and `answerAgain` are printed to the console, showing that they are independent of each other.
9. Pointer Usage: A pointer `pAnswer` is created that points to the `answer` struct using the address-of operator (`&`). This demonstrates how to work with pointers in Go.
10. Further Modification: The `value` field of the `answer` struct is modified again to 44. Since `pAnswer` points to `answer`, dereferencing `pAnswer` (using the `*` operator) will reflect this change.
11. Final Printing: The values of `answer` and the dereferenced pointer `*pAnswer` are printed, showing that both reflect the updated value of 44.

This code illustrates the difference between passing by value and passing by reference using pointers in Go, highlighting how changes to a struct can be reflected through a pointer to that struct.
*/
