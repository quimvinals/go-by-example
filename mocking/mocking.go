package mocking

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	for i := 3; i >= 1; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
