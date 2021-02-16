// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	a := 3.0
	b := 2.0
	pA := &a
	pB := &b

	swapValues(&a, &b)
	fmt.Printf("before swap pA: %p pB: %p *pA: %v  *pB: %v\n", pA, pB, *pA, *pB)
	pA, pB = swapAddresses(pA, pB)
	fmt.Printf("a: %v b: %v\n", a, b)
	fmt.Printf("after swap pA: %p pB: %p *pA: %v  *pB: %v\n", pA, pB, *pA, *pB)
}

func swapValues(a *float64, b *float64) {
	*a, *b = *b, *a
}

func swapAddresses(a *float64, b *float64) (*float64, *float64) {
	return b, a
}
