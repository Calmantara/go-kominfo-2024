package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
)

func main() {
	defer recoverExample()
	// channelUnbuffered() // panic -> all go routines are sleep
	// channelBuffered()
	// channelUnbufferedWithGoRoutine() // ga panic
	// dataSource()
	// dataSourceBounded()
	// dataSourceWG()

	// deferExample(100)
	// deferExample(3)
	// deferExample(8)
	// exitExample()
	// var int1 *int
	// errorExample(int1)

	chanSelect()
	usecase()
}

func usecase() { // 4s
	// fetch data
	ch1 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 1
		close(ch1)
	}()

	// do some logic
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		time.Sleep(10 * time.Millisecond)
	}

	// use data from fetched data
	var1 := <-ch1
	fmt.Println(var1)
}

func chanSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Millisecond)
		ch2 <- 2
	}()

	// fmt.Println(<-ch1)
	// fmt.Println(<-ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}
}

func recoverExample() {
	if val := recover(); val != nil {
		fmt.Println(val)
	}
	// recover untuk auto restart server ketika crash
}

func errorExample(int1 *int) {
	res, err := div(*int1, 1)
	if err != nil {
		// do something
		panic(err) // program kita crash
		return
	}
	fmt.Println("success", res)
}

func div(in1, in2 int) (int, error) {
	// requirement:
	// tidak boleh dibagi sama 0
	// in2 != 0
	if in2 == 0 {
		// err := fmt.Errorf("cannot divided by 0")
		// err := errors.New("cannot divided by 0")
		// err := errors.New("cannot divided by 0")
		err := errors.WithStack(fmt.Errorf("cannot divided by 0")) // mendeteksi di mana error ini terjadi
		fmt.Println(err.Error())
		return 0, err
	}

	return in1 / in2, nil
}

func exitExample() {
	i := 0
	for {
		if i%100 == 0 {
			os.Exit(0)
		}
	}

	// kita ada process connect database
	// ketika gagal connect
	// kita akan exit program
	// os.Exit(0)

	// kalau misal
	// program kita harus exit
	// karena ada pembaharuan program
	// os.Exit(1)
}

func channelUnbuffered() {
	// declaration channel:
	//  - buffered channel ~ seperti array
	//  - unbuffered channel ~ seperti slice
	// make(chan type, buffer)
	ch1 := make(chan int)

	// cara assign data
	// ke channel
	// assign = memasukkan data ke buffer
	ch1 <- 1

	// mengakses data dari channel
	fmt.Println(<-ch1)
}

func channelBuffered() {
	// declaration channel:
	//  - buffered channel ~ seperti array
	//  - unbuffered channel ~ seperti slice
	// make(chan type, buffer)
	ch1 := make(chan int, 1)

	// cara assign data
	// ke channel
	// assign = memasukkan data ke buffer
	ch1 <- 1

	// mengakses data dari channel
	fmt.Println(<-ch1)
}

func channelUnbufferedWithGoRoutine() {
	// declaration channel:
	//  - buffered channel ~ seperti array
	//  - unbuffered channel ~ seperti slice
	// make(chan type, )
	ch1 := make(chan int, 2)

	// cara assign data
	// ke channel
	// assign = memasukkan data ke buffer
	go func() {
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Millisecond)
		ch1 <- 2
		close(ch1)
	}()

	// mengakses data dari channel
	// fmt.Println(<-ch1) // tamu 1
	// fmt.Println(<-ch1) // tamu 2

	for ch := range ch1 {
		fmt.Println(ch)
	}
}

func dataSource() {
	ch := make(chan int)

	go processData(1, ch)

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func dataSourceBounded() {
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		go processData(i, ch)
	}

	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)
}

func dataSourceWG() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go processDataWG(&wg, i, i)
		// go func(id, data int) {
		// 	fmt.Println("got message from:", id, data)
		// 	wg.Done()
		// }(i, i)
	}
	wg.Wait()
}

func processData(id int, ch chan int) {
	for c := range ch {
		fmt.Println("got message from:", id, c)
	}
}

func processDataWG(wg *sync.WaitGroup, id int, data int) {
	fmt.Println("got message from:", id, data)
	wg.Done()
}

func deferExample(in int) {
	// defer cara kita memastikan
	// suatu process akan ter-eksekusi
	// sebelum function keluar
	// dan defer function harus
	// sudah dipanggil terlebih dahulu

	if in == 100 {
		fmt.Println("in == 100")
		return
	}

	defer fmt.Println("HEY END")
	if in%3 == 0 {
		fmt.Println("in % 3")
		return
	}
	fmt.Println("invalid")
}
