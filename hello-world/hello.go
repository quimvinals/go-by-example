package main

import "fmt"

const englishHelloPrefix = "Hello, "

func sayHello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Printf(sayHello("Quim"))
}
