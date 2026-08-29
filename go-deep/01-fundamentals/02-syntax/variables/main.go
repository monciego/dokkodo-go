package main

import "fmt"

func main() {
	// declaration styles
	var name string = "Clark" // explicit type
	var age = 30              // type inferred
	var height float64        // zero value: 0.0
	city := "Gotham"          // short declaration (inside function only)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(city)

	// multiple variable declaration
	var x, y int = 10, 20
	var (
		a = 1
		b = "hello"
		c = true
	)
	books, author := "Deep Works", "Cal Newport"

	fmt.Println(x, y)
	fmt.Println(a, b, c)
	fmt.Printf("%s by %s", books, author)
}
