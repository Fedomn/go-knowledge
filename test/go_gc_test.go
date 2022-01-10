package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

type Wrapper int

func findWrapper(r *Wrapper) {
	fmt.Printf("wrapper got: %v\n", *r)
}

// https://github.com/patrickmn/go-cache/blob/46f4078530/cache.go#L1113

func TestFinalizer(t *testing.T) {
	w := Wrapper(100)
	runtime.SetFinalizer(&w, findWrapper)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("GC...")
		runtime.GC()
	}
}
