package main

import "fmt"

const greetings = "Hello, "

func Hello(name string) string {
	return greetings + name
}

func main() {
	fmt.Println(Hello("Jericho"))
}
