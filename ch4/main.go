package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"sync"
	"time"
)

func main() {
	// interfaceExample()
	// interfaceExample2()
	// reflectExample()
	// concurrentExample()
	concurrentWaitGroup()
}

func concurrentWaitGroup() {

	// wait group
	// bisa membuat concurrent program
	// 1. tanpa delay
	// 2. bisa memastikan semua goroutine sudah selesai tereksekusi

	var wg sync.WaitGroup

	// mengindikasikan, kita ngespawn berapa banyak go routine
	// dan berapa banyak goroutine yang perlu kita tunggu sampai
	// process selesai
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("from goroutine 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("from goroutine 2")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("from goroutine 3")
	}()

	// method untuk menunggu semua goroutine selesai di eksekusi
	wg.Wait()

	var arr []int
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Int())
	}
	now := time.Now()
	var wg2 sync.WaitGroup
	// wg2.Add(1)
	// go func() {
	// 	olahArray(arr)
	// 	wg2.Done()
	// }()
	for _, val := range arr {
		wg2.Add(1)
		go func(i int) {
			// avoid data race
			// by passing val into i
			olahArray([]int{i})
			wg2.Done()
		}(val)
	}
	wg2.Wait()
	fmt.Println(time.Since(now))

	// apakah ada batasan kita spawn banyaknya
	// go routine?
	// 1. tipe CPU
	// 2. memory => 1x spawn goroutine 4KB

	// common problem goroutine:
	// 1. panic
	// 2. race conditino
	// 3. deadlock
	// 4. goroutine leak
	// https://blog.devgenius.io/5-useful-concurrency-patterns-in-golang-8dc90ad1ea61
}

func concurrentExample() {
	go fmt.Println("hello from async goroutine")
	fmt.Println("hello from main goroutine")
	time.Sleep(100 * time.Millisecond)

	// requirement
	// mengolah 3 array
	// secepat mungkin

	// generate array
	var arr1, arr2, arr3 []int
	for i := 0; i < 2; i++ {
		arr1 = append(arr1, rand.Int())
		arr2 = append(arr2, rand.Int())
		arr3 = append(arr3, rand.Int())
	}
	// concurrent
	// untuk melakukan process / mengolah data
	// dengan karakter data:
	// 1. urutan tidak penting
	// 2. satu data dengan data lainnya tidak ada hubungan
	now := time.Now()
	go olahArray(arr1)
	go olahArray(arr2)
	olahArray(arr3)
	fmt.Println(time.Since(now))

	// sequential
	fmt.Println("========")
	now = time.Now()
	olahArray(arr1) // 2s
	olahArray(arr2) // 2s
	olahArray(arr3) // 2s
	// program kita akan selesai dalam
	// waktu berapa lama?
	// 6s
	fmt.Println(time.Since(now))
}

func olahArray(arr []int) {
	for _, val := range arr {
		fmt.Println(val)
		time.Sleep(time.Second)
	}
}

func reflectExample() {
	// reflect bisa kita gunakan
	// untuk memvaildasi suatu interface type
	var interface1 interface{}
	interface1 = 10

	t := reflect.ValueOf(interface1)
	fmt.Println(t)

	var int1 int
	if t.Kind() == reflect.Int {
		int1 = interface1.(int)
	}
	fmt.Println(int1)

	// kegunaan reflect
	// adalah membuat suatu function
	// yang benar benar generic
}

// mini quiz
// design code
// memiliki mobil honda, mobil tesla
// setiap mobil akan bisa refill fuel
// memiliki motor yamaha
// motor juga bisa refill fuel

func interfaceExample2() {
	sq := Square{side: 10}

	ci := Circle{radius: 4}

	// yang bisa dimasukkan ke showarea
	// hanya yang satu golongan dengan Interface Shape
	ShowArea(ci)
	ShowArea(sq)

	// SOLID Progamming
	// sangat menganjurkan menggunakan interface
}

func interfaceExample() {
	// apa itu interface di go?
	// primitive / astract data type

	var interface1 interface{}
	// int => 64 byte
	interface1 = 10
	fmt.Println(interface1)
	// bisa tidak kita assign jadi string
	// string => 64 byte
	interface1 = "string?"
	fmt.Println(interface1)
	interface1 = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(interface1)

	var interface2 interface{}
	interface2 = "string"
	fmt.Println(interface2)

	// kalau misal dari interface
	// terus kita mau membuka topeng interface
	// sehingga kita balik ke jalan yang benar
	// sesuai dengan data type
	// "type casting"
	str1 := interface2.(string)
	fmt.Println(str1)

	var interface3 interface{}
	interface3 = 10

	// defence programming
	// terhindar dari PANIC!!!
	str2, ok := interface3.(string)
	if !ok {
		fmt.Println("failed to convert interface to string")
	}
	fmt.Println(str2)
	// menggunakan conditional statement => reflect
	return
}

func cliArgs() {
	fmt.Println(os.Args[1])
	// fungsi:
	// - menjalankan multi server:
	//		- http
	//		- grpc
	// - input argument untuk config
	// - multi command program
	// package yang mendukung CLI
	// - flag
	// - cobra: https://github.com/spf13/cobra
}
