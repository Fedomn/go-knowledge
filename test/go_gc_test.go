package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// original idea from https://eng.uber.com/how-we-saved-70k-cores-across-30-mission-critical-services/
// 目的：
// Garbage collection in Go is concurrent and involves analyzing all objects to identify which ones are still reachable.
// We would call the reachable objects the "live dataset."
// Go offers only one knob, GOGC, expressed in percentage of live dataset, to control garbage collection
// hard_target = live_dataset + live_dataset * (GOGC / 100).

// issue:
// It is not aware of the maximum memory assigned to the container and can cause out of memory issues.
// Our microservices have a significantly diverse memory utilization portfolio. 100MB instances were having a huge GC impact.

// 解决：
// GOGCTuner dynamically computes the correct GOGC value in accordance with the container’s memory limit and sets it using Go’s runtime API.
// Go forces a garbage collection every 2 minutes.

// https://github.com/cch123/gogctuner

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
