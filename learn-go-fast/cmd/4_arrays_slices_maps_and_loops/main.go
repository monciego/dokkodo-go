package main

import "fmt"

// array, length inferred
// slice, can grow
// map - a set of key-value pair

/* Notes:
Arrays - a fixed-length collection of data, same type, indexable, and stored in contiguous (next to each other in memory, with no gaps) memory locations.
*/

func main() {
	// array hold 3 elements, the length of array cannot be changed after it's initialized
	// default value of int32 is 0, so if it is just defined the default values are [0 0 0]
	// int32 is 4 bytes of memory and we store 3 elements, Go allocates 12 bytes of contiguos memory when initialized
	var intArr [3]int32

	// change the value of the element by index
	intArr[1] = 123
	fmt.Println(intArr)
	fmt.Println(intArr[0])
	// access element 1 and 2
	fmt.Println(intArr[1:3])

	// print the memory location using &
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// we can initialize an array using this syntax
	// var intArr2 [3]int32 = [3]int32{1, 2, 3}
	// intArr2 := [3]int32{1, 2, 3} // shorthand
	// you can ommited the 3 and have the number inferred by the compiler using the ... syntax
	intArr2 := [...]int32{1, 2, 3, 4, 5} // shorthand
	fmt.Println(intArr2)

	// slices are just wrapper around arrays, to give more general, powerful, convinient interface to sequence of data
	// under the hood slices are just arrays with additional functionality
	// by omitting the length value, we now have a slice
	var intSlice []int32 = []int32{6, 7, 8}
	// cap shows the capacity of the slice
	// how many elements it can hold before it needs to allocate a new underlying array.
	fmt.Printf("The length %v with capacity %v\n", len(intSlice), cap(intSlice))

	// add in slice using append
	intSlice = append(intSlice, 9)
	fmt.Printf("The length %v with capacity %v\n", len(intSlice), cap(intSlice))

	// you can append multiple values by using spread operator ...
	intSlice2 := []int32{10, 11}
	intSlice = append(intSlice, intSlice2...)

	// another way to create a slice is by using the make function
	// you can specificy the length of the slice and optionally specify the cap, by default the capacity is the length of the slice
	//example:
	// 	[]int{1, 2, 3}       → values you want
	// make([]int, 3)       → length you want
	// make([]int, 3, 10)   → length + capacity you want
	var intSlice3 []int32 = make([]int32, 3, 8)

	fmt.Println(intSlice)
	fmt.Println(intSlice3)      // [0 0 0]
	fmt.Println(cap(intSlice3)) // 8

	// # Map
	// key-value pair (can look up a value by its key)
	// to be continue...
}
