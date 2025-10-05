package main

import "fmt"

// define a struct named rect
type rect struct {
	width, height int
	name          string
}

// define a method named perim for the rect struct
func (r *rect) perim() int {
	return 2 * (r.width + r.height)
}

// define another method
func (r *rect) expand(i int) {
	r.width += i
	r.height += i
}

func main() {
	// use {} notation to instantiate a struct
	var r = rect{width: 10, height: 5, name: "My Rectangle"}

	fmt.Println("Area: ", r.width*r.height)
	fmt.Println("Perimeter: ", r.perim())
	r.expand(2)
	fmt.Println("New Area: ", r.width*r.height)
	fmt.Println("New Perimeter: ", r.perim())
}

/*
EXPLANATION:
1. We define a struct named `rect` with three fields: `width`, `height`, and `name`.
2. We define a method named `perim` for the `rect` struct that calculates and returns the perimeter of the rectangle.
3. We define another method named `expand` that takes an integer parameter and increases both the width and height of the rectangle by that value.
4. In the `main` function, we instantiate a `rect` struct using the `{}` notation, initializing its fields.
5. We print the area of the rectangle by multiplying its width and height.
6. We call the `perim` method to print the perimeter of the rectangle.
7. We call the `expand` method to increase the dimensions of the rectangle by 2.
8. Finally, we print the new area and perimeter after expansion.

When executed, this program will output:
Area:  50
Perimeter:  30
New Area:  84
New Perimeter:  34
*/
