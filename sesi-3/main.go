// cara1
package main

import "fmt"

func main() {
	kal := "selamat malam"
	jmlhuruf := make(map[rune]int)
	jmlhuruf2 := make(map[string]int)
	/*variabel ini bertujuan utk mengubah data variabel jmlhuruf*
	yang bertipe 'rune' yg menjadi 'string', mslahnya jk tidak
	dikonversikan maka otput nya seperti ini map[32:1 97:4 101:1 108:2 109:3 115:1 116:1]*/
	for _, huruf := range kal {
		fmt.Println(string(huruf))
	}
	for _, huruf := range kal {
		jmlhuruf[huruf]++
	}
	for k, v := range jmlhuruf { //bagian fungsi utk konversi jmlhuruf
		jmlhuruf2[string(k)] = v
	}
	fmt.Println(jmlhuruf2)
}

//cara2
// import (
// 	"fmt"
// 	"sort"
// ) //memakasi fungsi sort

// func main() {
// 	kal := "selamat malam"
// 	for _, huruf := range kal {
// 		fmt.Println(string(huruf))
// 	}
// 	jmlhuruf := make(map[rune]int)
// 	for _, huruf := range kal {
// 		jmlhuruf[huruf]++
// 	}
// 	var keys []rune
// 	for k := range jmlhuruf {
// 		keys = append(keys, k)
// 	}
// 	sort.Slice(keys, func(i, j int) bool {
// 		return keys[i] < keys[j]
// 	})
// 	fmt.Print("map[")
// 	//opsi1
// 	for _, k := range keys {
// 		fmt.Printf("%c:%d ", k, jmlhuruf[k])
// 		//endopsi
// 	}
// 	fmt.Print("]")
// }

/*
slice.key untuk mengambil semua key dari map jmlhuruf dan menyimpannya.Setelah itu,
slice keys diurutkan secara ascending dengan menggunakan fungsi sort.Slice, jdi jika opsi1
diganti dengan yang di bawah, maka posisi map [ :1 a:4 e:1 l:2 m:3 s:1 t:1 ] akan disesuaikan
berdasarkan pengaruh compile masing2 bisa menjadi seperti ini map [ a:4 e:1 m:3 l:2 s:1:1  t:1 ]
	for huruf, jumlah := range jmlhuruf
		fmt.Printf("%c:%d ", huruf, jumlah)

*/

