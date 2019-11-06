package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	//args := os.Args[1:]

	firstStr := ""

	for _, value := range os.Args[1:2] {
		//    fmt.Println(s)
		firstStr = firstStr + value
	}

	// fmt.Println(firstStr)

	rotRunes := []rune(firstStr)

	if argsLen(firstStr) < 1 {
		z01.PrintRune(10)
	} else {
		for i := 0; i <= runeLen(rotRunes); i++ {

			if rotRunes[i] >= 65 && rotRunes[i] <= 90 {
				if rotRunes[i] > 'M' {
					rotRunes[i] -= 13
					z01.PrintRune(rotRunes[i])
				} else {
					rotRunes[i] += 13
					z01.PrintRune(rotRunes[i])
				}

			} else if rotRunes[i] >= 97 && rotRunes[i] <= 122 {
				if rotRunes[i] > 'm' {
					rotRunes[i] -= 13
					z01.PrintRune(rotRunes[i])
				} else {
					rotRunes[i] += 13
					z01.PrintRune(rotRunes[i])
				}
			} else if rotRunes[i] == 32 {
				z01.PrintRune(rotRunes[i])
			}
		}
		z01.PrintRune(10)
	}
}

func argsLen(str string) int {
	count := 0

	for range str {
		count++
	}

	return count
}

func runeLen(runes []rune) int {
	count := 0

	for range runes {
		count++
	}

	return count - 1
}
