package main

import (
	"fmt"
	"strconv"
	"time"
)

type nilai map[string]int

// kapan alias digunakan:
// 1. menghindari conflict name dengan default data type
// 2. untuk menstandardkan value (nilai menggunakan map)

func main() {
	// conditions()
	// looping()
	// ArrSliceMap()
	// StringInDepth()
	fmtInDepth()
}

func conditions() {
	// calman pergi sekolah

	isHungry := false
	clock := 8

	// jika clock kurang dari jam 7, calman akan mandi
	// jika clock lebih dari sama dengan dari jam 7 dan calman lapar, calman akan sarapan
	// jika clock lebih dari jam 7, calman langsung berangkat
	// selain itu, calman bolos

	if clock < 7 {
		// akan sarapan
		fmt.Println("calman mandi")
	} else if clock >= 7 && isHungry == true {
		fmt.Println("calman sarapan")
	} else if clock >= 7 {
		fmt.Println("calman langsung berangkat")
	} else {
		fmt.Println("calman bolos")
	}

	// switch case
	clock = 4

	// kalau ada expression
	// SAMA DENGAN operation
	switch clock {
	case 4:
		fmt.Println("calman jogging dulu")
		// dia sudah ada break by nature
	case 5:
		fmt.Println("calman mandi")
	case 7:
		fmt.Println("calman berangkat")
	default:
		fmt.Println("calman bolos")
	}

	// tidak ada expression
	// kita bisa menambahkan condition kita
	clock = 4
	switch {
	case (clock <= 5):
		fmt.Println("calman mandi")
		// fallthrough akan mengeksekusi code bawahnya
		// tanpa melakukan pengecheckan kondisi
		fallthrough
	case (clock == 3):
		fmt.Println("calman joging")
	default:
		fmt.Println("calman bolos")
	}

	// apakah bisa kita if didalam if?
	if isHungry {
		// nested condition
		if clock < 7 {

		}
	}
	// TECH DEBT: Clean Code, Code Coverage, etc
}

func looping() {
	// golang hanya mengenal FOR

	// python, js, cpp
	// while , do while, for

	// infinte loop
	// looping gak berhenti henti
	for {
		// production: job
		fmt.Println("infinite loop")
		time.Sleep(time.Second)
	}

	// range loop
	for i := 0; i < 10; i++ {
		fmt.Printf("looping: %d\n", i)
	}
	// apakah dengan hal diatas,
	// masih bisa terjadi infinite loop?
	for i := 0; i < 10; i-- {
		fmt.Printf("looping broken: %d\n", i)

		// break untuk keluar dari loop
		if i == -100 {
			break
		}
	}

	// loop with condition
	i := 0
	for i < 10 {
		// do some process
		i++
	}
	fmt.Println("ok")

	// APAKAH BISA KITA ADA
	// LOOP DI DALAM LOOP?
	// nested loop

	// outer loop
	for i := 0; i < 100; i++ {

		// inner loop
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				fmt.Println(i, j, k)
			}
			fmt.Println(i, j)
		}
		// end of process outer loop
		// go to next loop
	}

	// BIG(O) Notation
	// time complexicy
}

func ArrSliceMap() {
	arr := [10]int{}
	fmt.Println("array:", arr)
	fmt.Println(arr[5:])

	slice := []int{}
	fmt.Println("slice:", slice)

	// diff

	// arr [0 0 1 0 0 0 0 0 0 0]
	// idx [0 1 2 3 4 5 6 7 8 9]
	arr[2] = 1
	// arr[2] => kita mengakses laci dengan index no 2
	fmt.Println(arr, arr[2])
	// arr = append(arr, 100)

	slice = append(slice, 10)
	fmt.Println(slice, slice[0])
	slice[0] = 190
	fmt.Println(slice, slice[0])
	// linked list

	// use case dari nested loop
	// apakah array bisa di nested?
	arrNested := [3][3]int{}
	fmt.Println(arrNested)

	// apa hubungan array / slice dengan for loop?
	// ketika kita mau mengakses array, kita bisa menggunakan
	// for loop

	// arr.map()
	// sangat membatasin built in function

	for idx, val := range arr {
		fmt.Println(idx, val)
	}
	// kalau kita mau mengakses arrNested
	// bagaimana?
	// nested loop

	// u. mendapatkan panjang dari arr
	// len(arr)
	for i := 0; i < len(arrNested); i++ {
		innerArr := arrNested[i]
		for j := 0; j < len(innerArr); j++ {
			fmt.Println(innerArr[j])
		}
	}

	// [
	// 	[1 2 3] index 0 dari luar array
	// 	[4 5 6] index 1 dari luar array
	// 	[7 8 9] index 2 dari luar array
	// ]
	fmt.Println(arrNested[0][0]) // 1
	fmt.Println(arrNested[0][1]) // 2
	fmt.Println(arrNested[1][1]) // 5
	fmt.Println(arrNested[2][0]) // 7

	// MAP map[key]value
	nilaiCalman := nilai{
		"matematika": 10,
		"fisika":     10,
		"biologi":    5,
		"kimia":      9,
	}
	fmt.Println("nilai matematika calman", nilaiCalman["matematika"])
	fmt.Println(nilaiCalman["kimia"]) // nil? 0?

	// val akan bernilai jika ada dan akan default jika tidak ada
	// ok akan true jika ada, akan false kalau tidak ada
	val, ok := nilaiCalman["kimia"]
	if ok == false {
		// do something
		fmt.Println("calman belum ujian kimia")
	} else {
		fmt.Println("nilai ujian kimia: ", val)
	}

	// apa hubungan for loop dan map
	for key, val := range nilaiCalman {
		fmt.Println(key, val)
	}
	// hati hati dalam menggunakan map
	var nilaiTara nilai // akan ada loker tapi belum ada loker
	nilaiTara = nilai{}
	nilaiTara["matematika"] = 10
}

func StringInDepth() {
	// string di go
	// kumpulan dari bytes
	// ASCII
	str := "Calman" // 67, 97,108, 109, 97, 110
	for idx, val := range str {
		fmt.Printf("%T", val)
		fmt.Println(idx, val)
	}

	b := []byte{67, 97, 108, 109, 97, 110}
	fmt.Println(string(b))

	// package strings, strconv
	// untuk konversi string ke integer
}

func fmtInDepth() {
	// package fmt

	// mau print
	fmt.Print("hello", "world")   // tidak ada spasi
	fmt.Println("hello", "world") // spasi dan enter
	// %d => digit
	// %s => string
	// %f => float
	// %0.2f => float dengan 2 decimal
	// %T => mengetahui Type Of
	// ...
	// %v
	// untuk enter \n
	fmt.Printf("%v\n", 10.999999)
	fmt.Printf("%v\n", 10.999999)

	_ = fmt.Sprintf("%s:%s:%s:%d", "nilai", "calman", "matematika", 100)
	// use case
	// 1. membuat string pattern
	// 2. konversi langsung ke string

	// untuk konversi integer ke string
	str1 := strconv.Itoa(10)
	str2 := fmt.Sprintf("%v", 10)
	fmt.Println(str1 == str2)
}
