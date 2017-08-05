package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Count(ch chan int, count int) {
	ch <- count
	fmt.Println("Send ", count)
}

func TestSharedVariableByChannel(t *testing.T) {
	chas := make(chan int)
	for i := 0; i < 10; i++ {
		go Count(chas, i)
	}

	for j := 0; j < 10; j++ {
		data := <-chas
		runtime.Gosched()
		fmt.Println("Receive ", data)
	}
}

func TestChannelTimeout(t *testing.T) {
	ch := make(chan int)

	select {
	case <-ch:
		fmt.Println("get ch")
	case <-time.After(time.Second):
		fmt.Println("get ch timeout 1s")
	}
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestFibonacci(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
