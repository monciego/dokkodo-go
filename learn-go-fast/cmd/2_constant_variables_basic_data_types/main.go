package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// arithmetic operations cannot be performed with mixed types:
	// var result_a float32 = floatNum32 + intNum32 // error
	// so you will cast one of the numbers (intNum32) to a common type (which is float)
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 23
	var intNum2 int = 6
	// integer division results in an integer and will round it down
	fmt.Println(intNum1 / intNum2) // 23 / 6 = 3
	// if you want to get the remainder use a % (modulo) sign
	fmt.Println(intNum1 % intNum2) // 32 % 6 = 5

	// you can use a doublequote("") or backtic/backquote(``)
	var myString string = "Hello World" + " " + "World"
	fmt.Println(myString)

	// you can use the backquote to format your string in the next without using \n
	var shelby string = `in the
bleak midwinter`
	fmt.Println(shelby)

	// get the length of the string using len() function
	// note: it prints the number of bytes not the number of characters
	println(len("test")) // 4 (this is byte btw)
	println(len("h"))    // 1
	// to check i will print the length of a japanese letter
	println(len("は")) // 3
	// another example gamma symbol (γ)
	println(len("γ")) // 2

	// if you want to use some fancy strings and want the length of the string in the number of characters
	// you can use the built-in package `unicode/utf8`
	// and call  the RuneCountInString function
	println(utf8.RuneCountInString("γ")) // 1

	konnichiwa := "こんにちは"
	println(len(konnichiwa))                    // 15
	println(utf8.RuneCountInString(konnichiwa)) // 5

	// single string is called a rune
	var myRune rune = 'a'
	fmt.Println(myRune)

	// boolean (true or false)
	var myBoolean bool = false
	fmt.Println(myBoolean)

	// default values (depends on its type)
	// for ints, unsigned ints, float and rune is 0
	// for strings is '' (empty string)
	// and for boolean is false
	var intNumA int
	var intNum3 rune
	var emptyString string
	// the first bool here is variable name, i found out that it is not in keyword
	// and can use it as a variable and i find it funny bulbol
	var bool bool

	println(intNumA)         // 0
	fmt.Println(intNum3)     // 0
	fmt.Println(emptyString) //
	fmt.Println(bool)        // false

	// you can set a variable but omit(exclude) the type righ away
	// the type is inferred
	// var myVar = "text" (type string)

	// shorthand version
	myVar := "text"
	fmt.Println(myVar)

	// adding a type when it is not obvious is a best practice
	// for example i use myVar := foo()
	// we don't know what is the return type of the foo unless you check it
	// so it is better to add the type when it's not obvious e.g var myVar string = foo()

	// you can initialize multiple variables at once
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// you cant change a value of a constant variable
	const myConst string = "const value"
	fmt.Println(myConst)
	// pi is the best example to use a constant because it is not supposed to be changed
	const pi float32 = 3.1415
}
