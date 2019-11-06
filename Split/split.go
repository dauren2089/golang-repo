package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args1 := []string(os.Args[1:])
	// args2 := []string(os.Args[2:3])

	// "HelloHAhowHAareHAyou?"
	// "HA"
	// "Hello how are you?"

	if len(args1) <= 1 {
		return
	} else {
		str := ""
		for _, v := range os.Args[1:2] {
			str = str + v
		}

		sep := ""

		for _, v := range os.Args[2:3] {
			sep = sep + v
		}

		fmt.Println(str)
		fmt.Println(sep)

		splitStr := []rune(str)
		i := 0

		for _, value := range splitStr {
			if value == rune(sep[i]) {
				z01.PrintRune(value)
				i++
			}
		}
		z01.PrintRune(10)

	}

}
