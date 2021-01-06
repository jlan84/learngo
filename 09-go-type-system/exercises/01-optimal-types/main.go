// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Optimal Types
//
//  1. Choose the optimal data types for the given situations.
//  2. Print them all
//  3. Try converting them to lesser data types.
//     For example try converting int64 variable to int32.
//     Then observe the result.
//     Search the web why the result is so?
//
// NOTE
//  This is just an exercise for teaching you different
//  data types. Do not apply it to the real-life.
//
//  As I said in the lectures that, premature optimization
//  is not a good thing.
// ---------------------------------------------------------

func main() {
	// DONT FORGET: There are also unsigned data types.
	//              (For positive numbers)

	// DO NOT USE: int data type
	// Use only the ones with the bitsizes

	// ---

	// an english letter (search web for: ascii control code)
	type letter byte
	// a non-english letter (search web for: unicode codepoint)
	type uniLetter rune
	// a year in 4 digits like 2040
	type year uint16
	// a month in 2 digits: 1 to 12
	type month uint8
	// the speed of the light
	type lightSpeed uint32
	// angle of a circle
	type angle float32
	// PI number: 3.141592653589793
	type pi float64
}
