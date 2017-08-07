package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

//基础类型:
//布尔(bool) 整型(int8...) 浮点(float32...) 复数(complex64...) 字符串(string) 字符(rune) 错误(error)
//复合类型:
//pointer array slice map chan struct interface

func TestPointer(t *testing.T) {
	p := 123
	fmt.Println(&p)

	i := &p
	fmt.Println(i)
	fmt.Println(*i)

	*i = 0
	fmt.Println(p)
}

func TestSlice(t *testing.T) {
	a := make([]int, 4, 5)
	fmt.Println(a, len(a), cap(a))

	b := a[:cap(a)]
	fmt.Println(b, len(b), cap(b))
}

func TestSliceAppend(t *testing.T) {
	var res []int
	fmt.Println("Capacity was:", cap(res))
	for i := 0; i < 5; i++ {
		res = append(res, i)
		fmt.Println("Capacity is now:", cap(res))
	}
	fmt.Println(res, len(res), cap(res))
}

func TestForRange(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	for index, value := range a {
		fmt.Println(index, value)
	}
}

func TestMap(t *testing.T) {
	type TestStruct struct {
		x float32
		y string
	}

	m := make(map[string]TestStruct)
	m["1"] = TestStruct{1, "123"}
	fmt.Println(m)

	v, isPresent1 := m["1"]
	if isPresent1 {
		fmt.Println(v, isPresent1)
	}

	v, isPresent2 := m["2"]
	if isPresent2 {
		fmt.Println(v, isPresent2)
	}
}

func TestSelect(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Recevce", msg1)
		case msg2 := <-c2:
			fmt.Println("Recevce", msg2)
		}
	}
}

func TestTimeouts(t *testing.T) {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 2")
	}
}

func TestCloseChannel(t *testing.T) {
	//接收方会一直阻塞直到有数据到来
	//如果channel是无缓冲的，发送方会一直阻塞直到接收方将数据取出
	//如果channel带有缓冲区，发送方会一直阻塞直到数据被拷贝到缓冲区
	//如果缓冲区已满，则发送方只能在接收方取走数据后才能从阻塞状态恢复
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive job", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 1; i < 4; i++ {
		jobs <- i
		fmt.Println("send job", i)
	}
	close(jobs)
	fmt.Println("send all jobs")
	fmt.Println(<-done)
}

func TestTimers(t *testing.T) {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("This 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("This 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("This 2 stoped")
	}
}

func TestTickers(t *testing.T) {
	ticker := time.NewTicker(time.Millisecond * 300)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Second * 1)
	ticker.Stop()
	fmt.Println("Ticker stop")
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func TestWorkerPoll(t *testing.T) {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j < 9; j++ {
		jobs <- j
	}

	close(jobs)
	for res := 1; res < 9; res++ {
		<-results
	}
}

func TestAtomicCounters(t *testing.T) {
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops final: ", opsFinal)
}

func TestEnum(t *testing.T) {
	const (
		one = iota
		two
		three
	)
	fmt.Println(one, two, three)
}

func throwPanic(f func()) (b bool) {
	defer func() {
		//recover()可以捕获到panic的输入值,并且恢复正常执行
		if _panic := recover(); _panic != nil {
			fmt.Println(_panic)
			b = true
		}
	}()
	f()
	return
}

func TestPanic(t *testing.T) {
	hasPanic := throwPanic(func() {
		fmt.Println("start throw panic...")
		panic("throw panic")
	})
	fmt.Println(hasPanic)
}
