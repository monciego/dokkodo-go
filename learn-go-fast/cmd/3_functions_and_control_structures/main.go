package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	// passing the printValue variable in printMe parameter
	// type is enforced we cant pass int in printMe function
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)

	// using printF function can format the string easier using variable
	/*
		%s  string
		%d  integer
		%f  float
		%t  bool
		%v  default/general
		%T  type
		%q  quoted string
		%x hexadecimal
		%b binary
	*/

	// using if statement
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	// }

	// using switch statement
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of integer division is %v\n", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was exact\n")
	case 1, 2:
		fmt.Printf("The division was close\n")
	default:
		fmt.Printf("The division was not close\n")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

// return a type
// you can return multiple value (returned result, remainder and an error)
// err is built in type let's handle it if it is cannot divided by 0
// nil is a default value for many types
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero\n")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, nil
}

/*
Notes:

nil is the zero value for several types, but not all.
nil applies to:

- pointer
- slice
- map
- function
- channel
- interface

For example:

var p *int    // nil
var s []int   // nil
var m map[string]int // nil
*/

