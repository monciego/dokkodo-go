package main

import "fmt"

const greetings = "Hello, "

func Hello(name string) string {
	if name == "" {
		return greetings + "World"
	}

	return greetings + name
}

func main() {
	fmt.Println(Hello("Jericho"))
}
