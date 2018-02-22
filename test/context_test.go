package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"testing"
	"time"
)

// 主要作用控制 多个相关的goroutine退出
// 场景: 一个请求会开启 多个goroutine 其中一个goroutine超时或者取消 其它相关的goroutine都应该立即退出 才能释放资源
// 列子:
// 一个请求要获取用户的许多维度信息 比如身份信息、最后登录时间、购买信息等。
// 这时我们会派生出其它goroutine来查询这些信息，当其中获取身份信息的goroutine超时或者取消时，
// 应该立即停止其它的查询goroutine 来释放资源。

// 对于简单的情况 我们可以通过chan + select处理 如下
func TestChanSelect(t *testing.T) {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控1退出，停止了...")
				return
			default:
				fmt.Println("goroutine1监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控2退出，停止了...")
				return
			default:
				fmt.Println("goroutine2监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

// chan+select 可以处理些简单的并行的goroutine 但是要多次 stop <- true 好傻。
// 但是对于那种 派生的goroutine 和 链式调用的goroutine 处理起来就很复杂了

// Context 的调用应该是链式的，通过WithCancel，WithDeadline，WithTimeout或WithValue派生出新的 Context。当父 Context 被取消时，其派生的所有 Context 都将取消。
// 它的所有方法
// func Background() Context
// func TODO() Context
//
// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
// func WithValue(parent Context, key, val interface{}) Context

func TestContextWithCancel(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func TestContextWithDeadline(t *testing.T) {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestContextWithTimeout(t *testing.T) {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}

func TestContextWithValue(t *testing.T) {
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}

// 链式调用示例 https://juejin.im/entry/5a151c1b6fb9a045055dc577#shi-pin-xin-xi
func TestContextErrGroup(t *testing.T) {
	eg, egCtx := errgroup.WithContext(context.Background())

	f := func(loc string) func() error {
		return func() error {
			if loc == "http://www.facebook.com" {
				fmt.Println("bingo facebook")
				time.Sleep(time.Second * 5)
			}
			reqCtx, cancel := context.WithTimeout(egCtx, time.Second)
			defer cancel()
			req, _ := http.NewRequest("GET", loc, nil)
			var err error
			_, err = http.DefaultClient.Do(req.WithContext(reqCtx))
			if err != nil {
				fmt.Printf("http %s err: %+v\n", loc, err)
			} else {
				fmt.Printf("http %s success\n", loc)
			}
			return err
		}
	}

	eg.Go(f("http://www.google.com"))
	eg.Go(f("http://www.baidu.com"))
	eg.Go(f("http://www.facebook.com"))
	if err := eg.Wait(); err != nil {
		fmt.Printf("err: %+v", err)
	}
}

// 总结
// 所有的长的、阻塞的操作都需要 Context
// errgroup 是构架于 Context 之上很好的抽象
// 当 Request 的结束的时候，Cancel Context
// Context.Value 应该被用于告知性质的事物，而不是控制性质的事物
// 约束 Context.Value 的键空间
// Context 以及 Context.Value 应该是不可变的（immutable），并且应该是线程安全
// Context 应该随 Request 消亡而消亡
