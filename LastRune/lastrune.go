package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(lastRune("Hello"))
	z01.PrintRune(lastRune("Salut"))
	z01.PrintRune(lastRune("Ola"))
	z01.PrintRune('\n')
	z01.PrintRune(firstRune("Hello!"))
	z01.PrintRune(firstRune("Salut!"))
	z01.PrintRune(firstRune("Ola!"))
	z01.PrintRune('\n')
	z01.PrintRune(nRune("Hello!", 3))
	z01.PrintRune(nRune("Salut!", 2))
	z01.PrintRune(nRune("Bye!", -1))
	z01.PrintRune(nRune("Bye!", 5))
	z01.PrintRune(nRune("Ola!", 4))
	z01.PrintRune('\n')
}

func lastRune(s string) rune {

	runes := []rune(s)
	var value rune
	count := 0

	for range runes {
		count++
	}

	for i := 0; i <= count-1; i++ {
		if runes[i] == runes[count-1] {
			value = runes[i]
		}
	}

	return value
}

func firstRune(s string) rune {

	runes := []rune(s)
	var value rune
	count := 0

	for range runes {
		count++
	}

	for i := 0; i <= count-1; i++ {
		if runes[i] == runes[0] {
			value = runes[i]
		}
	}

	return value
}

func nRune(s string, n int) rune {

	runes := []rune(s)
	var value rune
	count := 0

	for range runes {
		count++
	}
	if n < 0 {
		return value
	} else if n >= count {
		return value
	} else {
		for i := 0; i <= count-1; i++ {
			if runes[i] == runes[n] {
				value = runes[i] - 1
			}
		}
	}
	return value
}
