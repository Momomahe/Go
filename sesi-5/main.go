package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

////Coding(1) Coding dibawah ini untuk Gourutine versi acak

var mutex = &sync.Mutex{}

func inter1(data interface{}, num int) {
	mutex.Lock()
	defer mutex.Unlock()

	for i := 1; i <= 3; i++ {
		fmt.Printf("[%v%v]", data, i)
		time.Sleep(time.Duration(rand.Intn(60)) * time.Millisecond) // kasih ksemptan muncul tiap 0,060 dtk
	}
	fmt.Printf(" %v\n", num)
}

func main() {
	data1 := "bisa"
	data2 := "coba"
	for i := 1; i <= 4; i++ {
		go inter1(data1, i)
	}
	for i := 1; i <= 4; i++ {
		go inter1(data2, i)
	}
	time.Sleep(5 * time.Second)
	/*kasih 5 dtk, kalau gak error, kalau nggk nnti eksekusinya kalau kcptn nanti error, mending kasih jeda*/
	/*gk tau apkh algoritmanya yg rumit atau spek laptop ini rendah*/
}

////Batas Coding(1)

////Coding(2) Coding dibawah ini untuk Gourutine versi rapi

// var mutex = &sync.Mutex{}
// var dataturu = true

// func inter2(data interface{}, num int) {
// 	mutex.Lock()
// 	defer mutex.Unlock()
// 	//fungsi var boel dataturu gunanya utk mengatur turn data1 dan data2 supaya muncul bergiliran
// 	//soalnya udah pake mutex malah gak rapi, gk berubah2 jdi bingung
// 	for !dataturu && data == "bisa" || dataturu && data == "coba" {
// 		mutex.Unlock()
// 		time.Sleep(time.Millisecond)
// 		mutex.Lock()
// 	}

// 	for i := 1; i <= 3; i++ {
// 		fmt.Printf("[%v%v]", data, i)
// 		time.Sleep(time.Duration(rand.Intn(60)) * time.Millisecond)
// 	}
// 	fmt.Printf(" %v\n", num)

// 	if data == "bisa" {
// 		dataturu = false
// 	} else {
// 		dataturu = true
// 	}
// }

// func main() {
// 	data1 := "bisa"
// 	data2 := "coba"
// 	for i := 1; i <= 4; i++ {
// 		go inter2(data1, i)
// 		go inter2(data2, i)
// 	}
// 	time.Sleep(5 * time.Second)
// }
