package main

import "fmt"

func sayHello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Printf(sayHello("Quim"))
}
