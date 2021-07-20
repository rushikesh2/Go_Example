package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start")
	wg.Add(1)

	go doSomething()

	n := 2
	go multiplication(n)

	fmt.Println("end")
	wg.Wait()

}

func multiplication(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d \n", n*i)
	}
}

func doSomething() {
	fmt.Println("do something")
	wg.Done() // this is done
}
