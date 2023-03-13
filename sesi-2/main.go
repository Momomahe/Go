package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}
	for j := 0; j < 11; j++ {
		if j == 5 {
			r := "САШАРВО"
			//r := []rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041E'}
			/*
			utk variabel r jg bsa digunakan, akan ttpi variabel posisi dibawah dikalikan dngn 2
			"САШАРВО" adlh alfabet Cyrillic yg dignakan basaha rusia, olh karena itu terdapat error krn mengetik seperti biasa, 
			shngga mmbuat hasil unicode berbeda degn output yg diperintahkan, hal it disbbkn krn hasil hasil ketikan biasanya
			menggunakan alfabet latin/roman
			*/
			for posisi, loop := range r {
				fmt.Printf("Character %U '%c' starts at byte position %d\n", loop, loop, posisi)
				//fmt.Printf("Character %U '%c' starts at byte position %d\n", loop, loop, posisi*2)
			}
			continue
		}
		fmt.Println("Nilai j = ", j)
	}
}
