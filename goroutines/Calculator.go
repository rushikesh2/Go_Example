package main

import (
	"time"
)

func sum(a, b int, c chan int) {
	time.Sleep(time.Second * 5)
	z := a+b
	
	c <- z	
}

func sub(a, b int, c chan int) {
	time.Sleep(time.Second * 8)
	z := a-b
	c <- z	
}

// func main() {

// 	chan1 := make(chan int)
// 	chan2 := make(chan int)
	
	
	
// 	start := time.Now()
	
	
// 	go sum(2,3, chan1)
// 	go sub(3,4, chan2)
	
// 	add := <- chan1
// 	fmt.Println("addition is", add)
// 	s := <- chan2
// 	fmt.Println("substractino is", s)
	
// 	elapsed := time.Since(start)
//         fmt.Printf("page took %s", elapsed)
// }
