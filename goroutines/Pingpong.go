// I would like to see program writing "ping" and "pong" continously to
// console till the program exited by user.
// "Ping" should NOT be followed by "Ping" again. Same goes for "Pong"

package main

import (

)

func ping(c chan string) {
	str := "ping"
	c <- str

}

func pong(c chan string) {
	str := "pong"
	c <- str
}

func main() {

	c := make(chan string)

	for {
		go ping(c)
		pi := <-c
		fmt.Println(pi)

		go pong(c)
		po := <-c
		fmt.Println(po)

	}
}
