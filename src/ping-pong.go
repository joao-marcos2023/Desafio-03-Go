package main

import (
	"fmt"
	"time"
)

func ping(pi chan string) {
	for i := 0; ; i++ {
		pi <- "ping"
	}
}

func pong(po chan string) {
	for i := 0; ; i++ {
		po <- "pong"
	}
}

func print(pi, po chan string) {
	for {
		msg1 := <-pi
		msg2 := <-po
		fmt.Println(msg1)
		fmt.Println(msg2)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var pi chan string = make(chan string)
	var po chan string = make(chan string)

	go ping(pi)
	go pong(po)
	print(pi, po)

	var input string
	fmt.Scanln(&input)

}
