package main

import (
	"go_by_example/mocking"
	"os"
)

func main() {
	mocking.Countdown(os.Stdout)
}
