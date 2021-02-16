// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var p *computer
	fmt.Println(p)

	// compare the pointer variable to a nil value, and say it's nil
	var c computer
	c = computer{
		brand: "apple",
	}

	fmt.Println(c)
	// create an apple computer by putting its address to a pointer variable
	apple := &computer{brand: "apple"}
	// put the apple into a new pointer variable
	newApple := apple
	// compare the apples: if they are equal say so and print their addresses
	if apple == newApple {
		fmt.Printf("pointer: %p and %p &apple: %p &newApple: %p\n", apple, newApple, &apple, &newApple)
	}
	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{"sony"}
	// compare apple to sony, if they are not equal say so and print their
	// addresses

	if sony != apple {
		fmt.Printf("pointer: %p and %p &apple: %p &sony: %p\n", apple, sony, &apple, &sony)
	}

	// put apple's value into a new ordinary variable
	c = *apple
	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("&apple: %p apple: %p  &c: %p\n", &apple, apple, &c)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if c == *apple {
		fmt.Println("c and *apple are equal")
	}
	// print the values:
	// the value that is pointed by the apple and the new variable
	fmt.Printf("*apple: %v c: %v\n", *apple, c)

	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand
	change(sony, "hp")
	fmt.Println(sony.brand)
	// write a func that returns the value that is pointed by the given *computer
	// print the returned value
	v := returnValue(sony)
	fmt.Printf("%v\n", v)
	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
	constructor()
	constructor()
	constructor()
}

func change(c *computer, newBrand string) {
	c.brand = newBrand
}

func returnValue(p *computer) computer {
	return *p
}

func constructor() (p *computer) {
	p = &computer{"mac"}
	fmt.Printf("%p", &p)
	return p
}
