package main

import (
	"log"
	"sync"
	"testing"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		// Wait()会自动释放c.L，并挂起调用者的goroutine。之后恢复执行，Wait()会在返回时对c.L加锁。
		// 除非被Signal或者Broadcast唤醒，否则Wait()不会返回。

		// 由于Wait()第一次恢复时，c.L并没有加锁，所以当Wait返回时，调用者通常并不能假设条件为真。
		// 取而代之的是, 调用者应该在循环中调用Wait。（简单来说，只要想使用condition，就必须加锁。）
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)

	c.L.Lock()
	done = true
	c.L.Unlock()

	log.Println(name, "wakes all")

	c.Broadcast()
	//c.Signal()
}

func TestNotifyOtherGoroutines(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
}

//https://ieevee.com/tech/2019/06/15/cond.html
