package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// program golang
// bisa kita compile
// menjadi program yang
// bisa kita jalankan
// windows: .exe
// linux: bin

func main() {
	// fmt.Println("inside main", glob)
	// variable()

	operator()
}

func server() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

// aliasing / naming converter
// menghindari tabrakan variable
type calman = string

var glob = "GLOBAL" // global variable

const (
	// constanta
	// meskipun global variable,
	// dia tidak bisa di re-assign
	PI = 3.14
)

func variable() {
	// di go ada variable:
	// 1. string
	str := "hello world" // local variable
	fmt.Printf("tipe dari var str:%T\n", str)

	// 2. bool true / false
	boolean1 := true
	boolean2 := false
	fmt.Printf("tipe dari bool1:%T bool2:%T\n", boolean1, boolean2)

	// 3. int bisa bernilai - atau +
	int1 := -1
	int2 := 100
	fmt.Printf("%v, %v", int1, int2)

	// 4. uint -> unsigned integer
	uint3 := uint(100)
	fmt.Println(uint3)
	// uint3 = -100 ini akan error

	// 5. float64
	float := 6.0
	fmt.Printf("tipe data float:%T\n", float)

	// 6. rune -> alias
	runeVar := rune(64)
	fmt.Println(runeVar, string(runeVar))

	// 7. byte -> alias
	byteVar := byte(1)
	fmt.Printf("byte variable type:%T\n", byteVar)

	// 8. calman
	calman := calman("this is variable")
	fmt.Printf("calman variable type:%T\n", calman)

	// mendefine variable
	// direct => var1 := 100
	var var2 int // akan menjadi default value => 0
	fmt.Println(var2)
	var2 = 100 // assign value ke var2
	fmt.Println(var2)

	// guideline from uber
	// https://github.com/uber-go/guide/blob/master/style.md

	str1, str2, int3 := "abc", "def", 123
	_, _, _ = str1, str2, int3
	res1, _ := multiReturn()
	fmt.Println(res1)

	fmt.Println("inside variable", glob)
	// ------- advance variable
	// 1. interface
	// 2. slice / array
	// 3. struct
	// 4. map
}

func multiReturn() (int, int) {
	fmt.Println("inside multiReturn", glob)
	glob = "HAHA CHANGED IN MULTI RETURN"
	return 1, 2
}

func operator() {
	// operator aritmatika in go:
	// * kali
	mul := 100 * 100
	fmt.Println(mul) // 10000

	// / bagi
	div := 100 / 10
	fmt.Println(div) // 10

	// + tambah
	add := 100 + 100
	fmt.Println(add) // 200

	// - kurang
	sub := 100 - 100
	fmt.Println(sub) // 0

	// % modulo => sisa bagi
	// 9 / 2 => 4 sisa 1
	// 9 % 2 = 1
	mod := 9 % 2
	fmt.Println(mod) // 1

	// ++
	inc := 1
	inc++
	fmt.Println(inc) // 2

	// --
	dec := 10
	dec--
	fmt.Println(dec) // 9

	// operator relation
	// (bagaimana caranya kita mengkomparasikan
	// nilai 2 type yang sama)
	str1 := "hello"
	str2 := "world"

	// bagaimana kita check
	// apakah str1 sama dengan str2
	same := str1 == str2 // sama dengan
	fmt.Println(same)    // false
	// apakah str1 tidak sama dengan str2
	same = str1 != str2 // tidak sama dengan
	fmt.Println(same)   // true

	// number (float / integer)
	int1 := 100
	int2 := 1000

	// apakah int1 lebih dari int2
	more := int1 > int2
	// apakah int1 lebih dari sama dengan int2
	more = int1 >= int2

	// apakah int1 kurang dari int2
	less := int1 < int2
	// apakah int1 kurang dari sama dengan int2
	less = int1 <= int2

	// same , not same
	sameNum := 100 == 100
	notSameNum := 100 != 1000
	fmt.Println(more, less, sameNum, notSameNum)

	// type
	// compare := int1 == str1
	compare := 10 > 11.0
	fmt.Println(compare)

	// operator logical
	// logika matematika
	// AND OR
	// true and true => true
	// true and false => false
	// true or true => true
	// true or false => true
	float1 := 100.9
	float2 := 100.8

	// float1 > float2
	// AND float1 != 0
	logical := float1 > float2 && float1 != 0
	// float1 < float2
	// OR float1 != 0
	logical = float1 > float2 || float1 != 0
	_ = logical
}
