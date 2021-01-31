// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// --- #1 ---
	// 1. create a new slice named: games
	var games []string
	// 2. print the length and capacity of the games slice
	fmt.Printf("games length: %d\ngames capacity: %d\n", len(games), cap(games))
	// 3. comment out the games slice
	//    then declare it as an empty slice
	games2 := []string(nil)
	// 4. print the length and capacity of the games slice
	fmt.Printf("games2 length: %d\ngames2 capacity: %d\n", len(games2), cap(games2))
	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	games = append(games, "pacman", "mario", "tetris", "doom")
	// 6. print the length and capacity of the games slice
	fmt.Printf("games length: %d\ngames capacity: %d\n", len(games), cap(games))
	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 3)
	game3 := []string{"pacman", "mario", "tetris", "doom"}
	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	for i := 1; i <= 4; i++ {
		x := game3[i-1 : i]
		fmt.Printf("game[%d:%d]: Length: %d Capacity: %d\n", i-1, i, len(x), cap(x))
	}

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	zero := game3[:0]
	// 2. print the games and the new slice's len and cap
	fmt.Printf("game[:0] length = %d and capacity = %d\n", len(zero), cap(zero))
	// 3. append a new element to the new slice
	zero = append(zero, "dark souls")
	// 4. print the new slice's lens and caps
	fmt.Printf("new length after appending length = %d and capacity = %d\n",
		len(zero), cap(zero))
	// 5. repeat the last two steps 5 times (use a loop)
	new_games := [5]string{"demon souls", "dark souls2", "dark souls 3",
		"sekiro", "bloodborne"}
	for i, val := range new_games {
		zero = append(zero, val)
		fmt.Printf("append #%d: length = %d and capacity = %d\n", i, len(zero),
			cap(zero))
	}
	// 6. notice the growth of the capacity after the 5th append
	//
	// Use this slice's elements to append to the new slice:
	other_games := []string{"ultima", "dagger", "pong", "coldspot", "zetra"}
	zero = append(zero, other_games...)
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	// for ... {
	//   ...
	//   fmt.Printf("zero's      len: %d cap: %d\n", ...)
	// }

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	for i := range zero {
		if i-1 < 0 {
			x := zero[0:i]
			fmt.Printf("zero[%d:%d] : length = %d cap = %d\n", 0, i,
				len(x), cap(x))
		} else {
			x := zero[i-1 : i]
			fmt.Printf("zero[%d:%d] : length = %d cap = %d\n", i-1, i, len(x),
				cap(x))
		}

	}
	// observe that, the range loop only loops for the length, not the cap.
	fmt.Println("Using the length to declare slice")
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(s), cap(s))
	}

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	zero = zero[:cap(zero)]
	// 2. print the elements of the zero slice in the loop.
	fmt.Println("Using the capacity instead of length to declare the slice")
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(s), cap(s))
	}

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	zero[0] = "Remnant"
	// 2. change the same element of the games slice
	games[0] = "Fallout"
	// 3. print the games and the zero slices
	fmt.Printf("games: %s\nzero: %s\n", games, zero)
	// 4. observe that they don't have the same backing array
	fmt.Println()
	// ...
	// fmt.Printf("zero  : %q\n", zero)
	// fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
}
