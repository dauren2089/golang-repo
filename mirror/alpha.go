package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	//args := os.Args[1:]

	firstStr := ""

	for _, value := range os.Args[1:2] {
		firstStr = firstStr + value
	}

	rotRunes := []rune(firstStr)

	for i := 0; i <= argsLen(firstStr)-1; i++ {
		if rotRunes[i] == 32 {
			z01.PrintRune(rotRunes[i])
		} else if rotRunes[i] >= 65 && rotRunes[i] <= 90 {
			rotRunes[i] = 90 - rotRunes[i] + 65
			z01.PrintRune(rotRunes[i])
		} else if rotRunes[i] >= 97 && rotRunes[i] <= 122 {
			rotRunes[i] = 122 - rotRunes[i] + 97
			z01.PrintRune(rotRunes[i])
		}
	}

	z01.PrintRune(10)
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

// func main() {

// 	str := "ABtynbek"
// 	runes := []rune(str)

// 	fmt.Println("initial: ", string(runes))

// 	for index := range runes {
// 		for i, j := 'a', 'A'; j <= 'Z'; i, j = i+1, j+1 {
// 			if runes[index] == i {
// 				runes[index] = 'z' - i + 'a'
// 				fmt.Println(string(runes), " index: ", index)
// 				break
// 			} else if runes[index] == j {
// 				runes[index] = 'Z' - j + 'A'
// 				fmt.Println(string(runes), " Rindex: ", index)
// 				break
// 			}
// 		}
// 		if runes[index] == ' ' {
// 			runes[index] = ' '
// 		}
// 	}
// }
