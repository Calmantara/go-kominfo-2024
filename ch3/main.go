package main

import (
	"fmt"

	"github.com/Calmantara/go-kominfo-2024/ch3/photo"
)

var result int

func main() {
	// fmt.Println(sum1([]int{1, 2, 3, 4}))

	// arr := []int{1, 2, 3, 4}
	// fmt.Println(
	// 	sum2(1, 2, 3, 4, 5, 6, 7, 8, 9),
	// 	sum2(),
	// 	sum2(arr...),
	// )

	// function kita (pointer)
	// tidak hanya mengeluarkan output
	// tapi juga mengubah nilai input

	// pure function
	// => dia menerima input
	// => dia memeberikan output
	// => tanpa mengubah input atau global variable

	// bisa ga function
	// input => function
	// output => function

	// function callback
	// cb := func(i int) int {
	// 	return i * 100
	// }
	// inputFunc(10, cb) // akan error
	// karena cb func(int)int
	// yang diharapkan adalah func(int)string

	// inputFunc(10, addOne)

	// // call All function
	// fn1 := func() string {
	// 	return "calman"
	// }
	// fn2 := func() string {
	// 	return "tara"
	// }
	// callAll(fn1, fn2, func() string {
	// 	return "golang 006"
	// })

	// function sebagai output dari function
	// fn := middleware("golang_006") // fn => func(int) string

	// func fn(i int) string {
	//   return fmt.Sprintf("golang_006:%d", i)
	// }
	// fmt.Println(fn(100))

	// pointer()
	// structExample()

	beautyPhoto := photo.Photo{
		ID:  1,
		Url: "https://www.google.com",
	}
	fmt.Println(beautyPhoto)
	fmt.Println(photo.GetPhotoPrefix(beautyPhoto))
}

// type StructName struct{}
type User struct {
	// tipe data
	// dengan nama User
	ID        uint64
	FirstName string
	LastName  string
	Email     string
	Address   Address
}

// suatu struct
// bisa memiliki dedicated function
// => method
func (user User) Greeting() {
	// greeting adalah suatu method punyanya user
	fmt.Println("Hello,", user.FirstName, user.LastName)
}

func (user *User) ChangeLastName() {
	// user => akan direfer ke address yang sama
	// sehingga changing the value, akan berpengaruh
	// ke struct variable yang memiliki addrs tsb
	user.LastName = "CHANGED!"
}

func (user User) ChangeLastName2() {
	// user => akan di copy value ke address lain
	// sehingga changing the value, tidak akan berpengaruh
	// ke struct variable
	user.LastName = "CHANGED!"
}

type Address struct {
	Name    string
	ZipCode uint32
}

func structExample() {
	// struct apa?
	// => tipe data bentukan
	// => tipe data dari kumpulan tipe data lainnya
	addStudent := Address{
		Name:    "Hacktiv",
		ZipCode: 2024,
	}
	student := User{
		ID:        1,
		FirstName: "Golang",
		LastName:  "006",
		Email:     "golang_006@mail.com",
		Address:   addStudent,
	}
	fmt.Printf("%+v\n", student)

	student2 := User{
		ID:        1,
		FirstName: "Golang",
		LastName:  "006",
		Email:     "golang_006@mail.com",
		Address: Address{
			Name:    "H8",
			ZipCode: 2602,
		},
	}
	fmt.Printf("%+v\n", student2)

	// struct bisa menjadi
	// array of struct
	// map of struct
	// bisa untuk input param func
	// bisa untuk out param func

	// bagaimana caranya
	// kita mengakses property variable dari struct?
	// ex: hello FirstName + LastName
	fmt.Printf("Hello, %v %v\n", student.FirstName, student.LastName)
	fmt.Printf("From, %v %v\n", student.Address.Name, student.Address.ZipCode)

	fmt.Println("=============")
	// Greeting(student)
	student.Greeting()
	student.ChangeLastName()
	student.Greeting()
	fmt.Println("=============")

	fmt.Println("before", student.FirstName)
	changeUser(&student)
	// apa yang terjadi pada variable student?
	fmt.Println("after", student.FirstName)
}

func changeUser(usr *User) {
	usr.FirstName += "changed in function"
}

func pointer() {
	var abc *int
	var def int
	// => abc adalah variable pointer of integer
	// => def adalah variable integer

	abc = &def
	// kita refer address of ABC menjadi address of DEF
	def = 100

	// mengubah nilai dari value of integer
	// kita mengubah value dari ABC
	*abc = 1000

	// * untuk mengakses value
	// & untuk mengakses address

	// arr := []*int{}
	// var ptMap *map[int]int

	// var pt **int
	// var pt2 ***int

	_, _ = abc, def

	arrIn1 := &[]int{1, 2, 3, 4} // addrs 123
	fmt.Println("before call change1:", *arrIn1)
	change1(arrIn1)
	fmt.Println("after call change1:", *arrIn1)

	arrIn2 := []*int{}
	for i := 0; i < 5; i++ {
		arrIn2 = append(arrIn2, &i)
	}
	// arrIn2 => addrs 234
	// rumah => a1, a2, a3, a4, a5
	for _, val := range arrIn2 {
		fmt.Println("before call change2:", *val)
	}
	change2(arrIn2)
	for _, val := range arrIn2 {
		fmt.Println("after call change2:", *val)
	}

	mp1 := map[string]int{
		"val1": 1,
		"val2": 2,
	}

	// mutable vs immutable variable
	fmt.Println("before call changeMap1", mp1)
	changeMap1(mp1)
	fmt.Println("after call changeMap1", mp1)

	fmt.Println("before call changeMap2", mp1)
	changeMap2(&mp1)
	fmt.Println("after call changeMap2", mp1)
}

func changeMap1(mp map[string]int) {
	// copy
	for key, val := range mp {
		mp[key] = val + 1
	}
}

func changeMap2(mp *map[string]int) {
	for key, val := range *mp {
		(*mp)[key] = val + 1
	}
}

func change1(arr *[]int) {
	// arr => addrs 123
	for idx, val := range *arr {
		(*arr)[idx] = val + idx
	}
}

func change2(arr []*int) {
	// arr => addrs 456
	// arr rumah rumah => a1, a2, a3, a4, a5
	for idx, val := range arr {
		a := (*val) + idx
		arr[idx] = &a
	}
}

func middleware(name string) func(int) string {
	return func(i int) string {
		return fmt.Sprintf("%s:%d", name, i)
	}
}

type shout func() string

func callAll(shouts ...shout) {
	for _, fn := range shouts {
		fmt.Println(fn())
	}
}

func addOne(in int) string {
	// termasuk golongan callback
	// return in + 1
	return ""
}

// bisa ga kita buat alias untuk function as parameter?
type callback func(int) string

func inputFunc(in int, fn callback) {
	// functionName => inputFunc
	// INPUT =>
	// 	1. in integer
	// 	2. fn func(int) int
	// 		-> func(int) int = data type
	fmt.Println(fn(in))
}

// array / slice
// map
// pre-assignment:
// hitunglah jumlah angka yang muncul dari suatu array
// array = [1,2,3,4,1,5,3,7,8,9,0]
// output = map {1: 2, 2: 1, 3: 2, 4: 1, 5: 1, 7: 1, 8: 1, 9: 1, 0: 1}

func solution() {
	// leet code
	// codewar

	var data = []int{1, 2, 3, 4, 5, 1, 5, 3, 7, 8, 9, 0}
	var output = map[int]int{}
	fmt.Println("data: ", data)
	for _, v := range data {
		output[v] += 1
	}
	fmt.Println(output)
}

// GOLANG
// func functionName(input) output
// functionName => solution
// intput => kosong
// output => kosong

func add(a int, b int) (int, int) {
	// functionName => add
	// input => a int, b int
	// output => int
	return a + b, a + a
}

// kita belum tau, inputnya ada berapa?
func sum1(nums []int) int {
	res := 0
	for _, val := range nums {
		res += val
	}
	return res
}

func sum2(nums ...int) int {
	res := 0
	for _, val := range nums {
		res += val
	}
	return res
}
