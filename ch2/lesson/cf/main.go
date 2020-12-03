package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/training/ch2/lesson/tempcomv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempcomv.Fahrenheit(t)
		c := tempcomv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempcomv.FToC(f), c, tempcomv.CToF(c))
	}
}
