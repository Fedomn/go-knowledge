package main

//https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6
import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func asChan(data ...int) <-chan int {
	res := make(chan int)
	go func() {
		for _, v := range data {
			res <- v
		}
		close(res)
	}()
	return res
}

func TestMergeChans(t *testing.T) {
	a := asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c := asChan(20, 21, 22, 23, 24, 25, 26, 27, 28, 29)
	for v := range merge(a, b, c) {
		fmt.Println(v)
	}
	fmt.Println("merge over")
	return
}

func mergeReflect(chans ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface().(int)
		}
	}()
	return out
}

func TestMergeReflect(t *testing.T) {
	a := asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c := asChan(20, 21, 22, 23, 24, 25, 26, 27, 28, 29)
	for v := range mergeReflect(a, b, c) {
		fmt.Println(v)
	}
	fmt.Println("merge reflect over")
	return
}

func mergeTwo(a, b <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func mergeRec(chans ...<-chan int) <-chan int {
	switch len(chans) {
	case 0:
		c := make(chan int)
		close(c)
		return c
	case 1:
		return chans[0]
	case 2:
		return mergeTwo(chans[0], chans[1])
	default:
		m := len(chans) / 2
		return mergeTwo(
			mergeRec(chans[:m]...),
			mergeRec(chans[m:]...))
	}
}

func TestMergeRec(t *testing.T) {
	a := asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c := asChan(20, 21, 22, 23, 24, 25, 26, 27, 28, 29)
	for v := range mergeRec(a, b, c) {
		fmt.Println(v)
	}
	fmt.Println("merge recursion over")
	return
}

func BenchmarkMerge(b *testing.B) {
	var funcs = []struct {
		name string
		f    func(...<-chan int) <-chan int
	}{
		{"goroutines", merge},
		{"reflection", mergeReflect},
		{"recursion", mergeRec},
	}
	for _, f := range funcs {
		for n := 1; n <= 1024; n *= 2 {
			chans := make([]<-chan int, n)
			b.Run(fmt.Sprintf("%s/%d", f.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					for i := range chans {
						chans[i] = asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
					}
					b.StartTimer()

					c := f.f(chans...)
					for range c {
					}
				}
			})
		}
	}
}
