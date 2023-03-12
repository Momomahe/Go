package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}
	for j := 0; j < 11; j++ {
		if j == 5 {
			r := []rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041E'}
			for posisi, loop := range r {
				fmt.Printf("Character %U '%c' starts at byte position %d\n", loop, loop, posisi*2)
			}
			continue
		}
		fmt.Println("Nilai j = ", j)
	}
}
