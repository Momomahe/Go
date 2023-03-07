package main

import "fmt"

func main() {
	i := 21
	j := true
	k := 123.456
	russia := 'Я'

	fmt.Printf("%v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")
	fmt.Printf("%t\n", j)
	fmt.Printf("%b\n", i)      // menampilkan hasil biner dari "i"
	fmt.Printf("%d\n", i)      // menampilkan hasil desimal dari "i"
	fmt.Printf("%o\n", i)      // menampilkan hasil oktal dari "i"
	fmt.Printf("%x\n", 15)     // untuk mendapatkan hex dgn kode f perlu mencari angka yang cocok
	fmt.Printf("%X\n", 15)     // untuk mendapatkan hex dgn kode F perlu mencari angka yang cocok
	fmt.Printf("%U\n", russia) // %U (menampilkan unicode)
	fmt.Printf("%.3f\n", k)
	fmt.Printf("%.6f\n", k)
	fmt.Printf("%E\n", k)

	/* beda f dan F hanya terletak pada format outpu "%x"(utk huruf kecil)
	dan "%X"(utk huruf besar)*/

}
