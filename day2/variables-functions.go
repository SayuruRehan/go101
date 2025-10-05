package main

import "fmt"

func main() {
	var foo = 3
	bar := "Hello"
	greet(bar, foo)
}

func greet(greeting string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(greeting)
	}
}

/* 
EXPLANATION:
1. We declare the main package and import the "fmt" package for formatted I/O.
2. In the main function, we declare a variable 'foo' with an initial value of 3 using the 'var' keyword.
3. We also declare a variable 'bar' and initialize it with the string "Hello" using the short variable declaration syntax ':='.
4. We call the 'greet' function, passing 'bar' and 'foo' as arguments.
5. The 'greet' function takes two parameters: a string 'greeting' and an integer 'times'.
6. Inside the 'greet' function, we use a for loop to print the greeting message the specified number of times.
7. The program will output "Hello" three times when executed.
*/
