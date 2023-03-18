package main

import (
	"fmt"
	"os"
)

type Kelas struct {
	No        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var tmn = []Kelas{
	{1, "Yoga Mahendra", "Sumatera Barat", "Programmer", "Ingin mempelajari golang lebih dalam"},
	{2, "Cahyono", "Papua", "Programmer", "Ingin meningkatkan kompetensi diri"},
	{3, "Rudi Widodo", "Jakarta", "Data Scient", "karena ikut-ikutan"},
	{4, "Viving Putri K", "Padang", "Data Scient", "Menambah ilmu"},
}

func main() {
	absen, err := getNoFromArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	isidt := getKelasByNo(absen)
	if isidt == nil {
		fmt.Println("Untuk Absen No", absen, "Tidak Tersedia")
		fmt.Println("No Absen hanya tersedia 1-4")
		return
	}
	printKelasData(isidt)
}
func getNoFromArgs() (int, error) { //error
	args := os.Args[1:] //tempt error kalau perintah kosong(krg dari 1/null)
	if len(args) < 1 {
		return 0, fmt.Errorf("Masukan no urut absen, contoh : main.go 2")
	}
	absen := 0 //tempt error kalau mskn data tipe string
	_, err := fmt.Sscanf(args[0], "%d", &absen)
	if err != nil {
		return 0, fmt.Errorf("Anda salah input, bukan '%s', Harap Masukan Angka 1 - 4", args[0])
	}

	return absen, nil
}
func getKelasByNo(absen int) *Kelas {
	for i := 0; i < len(tmn); i++ {
		if tmn[i].No == absen {
			return &tmn[i]
		}
	}
	return nil
}
func printKelasData(isidt *Kelas) {
	fmt.Println("Nama :", isidt.Nama)
	fmt.Println("Alamat :", isidt.Alamat)
	fmt.Println("Pekerjaan :", isidt.Pekerjaan)
	fmt.Println("Alan memilih kelas golang :", isidt.Alasan)
}
