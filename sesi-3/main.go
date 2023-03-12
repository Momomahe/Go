package main

import (
	"fmt"
	"sort"
)

func main() {
	kalimat := "selamat malam"
	for _, huruf := range kalimat {
		fmt.Println(string(huruf))
	}
	jumlahHuruf := make(map[rune]int)
	for _, huruf := range kalimat {
		jumlahHuruf[huruf]++
	}
	var keys []rune
	for k := range jumlahHuruf {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	fmt.Print("map [")
	for _, k := range keys {
		fmt.Printf("%c:%d ", k, jumlahHuruf[k])
	}
	fmt.Print("]")
}

/*
slice.key untuk mengambil semua key dari map jumlahHuruf dan menyimpannya.Setelah itu,
slice keys diurutkan secara ascending dengan menggunakan fungsi sort.Slice.
*/
