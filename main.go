package main

import (
	"fmt"
	"os"
)

func main() {
	app := New()

	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
