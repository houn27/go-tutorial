package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func RandAss(result int) {
	fmt.Printf("Goal is %v, start guess .... \n", result)
	var is_right = false
	for !is_right {
		var num = rand.Intn(100) + 1
		is_right = num == result
		if is_right {
			fmt.Println("Guess correctly!")
		} else {
			fmt.Printf("Guess wrongly with %v, continue .... \n", num)
		}
	}

}

func ScopeAss() {
	year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31
	is_leap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)

	switch month {
	case 2:
		if is_leap {
			daysInMonth = 28
		} else {
			daysInMonth = 29
		}

	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	fmt.Println(year, month, day)
}

func IfTest() {
	// string/if example
	var long_str, short_str = "walk outside", "walk"
	var short_in_long = strings.Contains(long_str, short_str)
	if short_in_long {
		fmt.Printf("\"%v\" conatins \"%v\" \n", long_str, short_str)
	} else {
		fmt.Printf("\"%v\" does not conatin \"%v\" \n", long_str, short_str)
	}
}

func RandTest() {
	// rand example
	var num = rand.Intn(10) + 1
	fmt.Println("rand.Intn(10) output: ", num)
}

func PrintTest() {
	// format print
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galatic", 100)
}

func ForTest() {
	fmt.Println("For test from 0 to 9: ")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {

}
