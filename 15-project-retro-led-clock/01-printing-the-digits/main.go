// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	type placeholder [5]string
	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	sep := placeholder{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}
	sep2 := placeholder{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

	numbers := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine, sep,
	}
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	for {
		screen.MoveTopLeft()
		t := time.Now()
		h, m, s := t.Hour(), t.Minute(), t.Second()

		var clock [8]placeholder
		if s%2 != 0 {
			clock[2], clock[5] = sep, sep
		} else {
			clock[2], clock[5] = sep2, sep2
		}
		clock[0], clock[3], clock[6] = numbers[h/10], numbers[m/10], numbers[s/10]
		clock[1], clock[4], clock[7] = numbers[h%10], numbers[m%10], numbers[s%10]

		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], "   ")
			}
			fmt.Print("\n")
		}
		time.Sleep(1 * time.Second)
		// c := exec.Command("clear")
		// c.Stdout = os.Stdout
		// c.Run()
		screen.Clear()
	}

	// fmt.Printf("The time is %d : %d : %d\n", h, m, s)

}
