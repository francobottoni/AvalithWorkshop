package main

import (
	"fmt"
	"os"
	"time"
)

func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(700 * time.Millisecond)
		ponger <- 1
	}
}

func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(400 * time.Millisecond)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	ping <- 1

	time.Sleep(10 * time.Second)
	os.Exit(0)
}
