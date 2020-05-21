package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 20)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(30 * time.Millisecond)
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(30 * time.Millisecond)
			c <- i // 添加一个i后必须等待读出去后才会走
		}
		done <- true
	}()

	// we block here until done <- true
	<-done
	<-done
	close(c)

	// to unblock above
	// we need to take values off of chan c here
	// but we never get here, because we're blocked above
	for n := range c {
		fmt.Println(n)
	}
}
