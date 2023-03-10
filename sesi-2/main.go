package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}
	for j := 0; j < 11; j++ {
		if j == 5 {
			unicode := "\u0421\u0410\u0428\u0410\u0420\u0412\u041E"
			for rusia := 0; rusia < len(unicode); {
				r, size := utf8.DecodeRuneInString(unicode[rusia:])
				fmt.Printf("Character U+%04X '%c' starts at byte position %d\n", r, r, rusia)
				rusia += size
			}
			continue
		}
		fmt.Println("Nilai j = ", j)
	}
}
