package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	fmt.Println(now)

	mins := now % 2
	fmt.Println(mins)

	switch mins {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	}

	fmt.Println("- - - - - - - - - - - - - - - -")

	num := 3
	v := num % 2
	switch v {
	case 0:
		fmt.Println("even")
	case 3 - 2:
		fmt.Println("odd")
	}

	fmt.Println("- - - - - - - - - - - - - - - -")

	score := 111
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("mhm did you cheat?")
	default:
		fmt.Println(score, "off the chart")
	}

	fmt.Println("- - - - - - - - - - - - - - - -")

	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough //if case matches all other conditions will be executed too
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
		fallthrough
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
		fallthrough
	default:
		fmt.Println("Try again!")
	}

	fmt.Println("- - - - - - - - - - - - - - - -")

	m := 3
	switch m {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("<= 1")
		fallthrough
	case 2:
		fmt.Println("<= 2")
		fallthrough
	case 3:
		fmt.Println("<= 3")
		if time.Now().Unix()%2 == 0 {
			fmt.Println("un pasito pa' lante Maria") // <= wasn't expecting that one
			break
		}
		fallthrough
	case 4:
		fmt.Println("<= 4")
		fallthrough
	case 5:
		fmt.Println("<= 5")
	}
}
