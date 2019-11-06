package main

import (
	piscine ".."
	"github.com/01-edu/z01"
)

func main() {
	result := piscine.Rot13("Hello How are You")
	arrayRune := []rune(result)

	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
