package main

import (
	"fmt"
	"sync"
	"testing"
)

// https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c

// And now we’ve arrived as to why the sync.Map was created.
// The Go team identified situations in the standard lib where performance wasn’t great.
// There were cases where items were fetched from data structures wrapped in a sync.RWMutex,
// under high read scenarios while deployed on very high multi-core setups and performance suffered considerably.

// fatal error: concurrent map read and map write
func ExampleConcurrencyMap() {
	a := make(map[int]int)
	go func() {
		for {
			_ = a[1]
		}
	}()

	go func() {
		for {
			a[2] = 2
		}
	}()

	select {}
}

type MutexMap struct {
	sync.RWMutex
	Map map[int]int
}

func (m *MutexMap) load(key int) (int, bool) {
	m.RLock()
	v, ok := m.Map[key]
	m.RUnlock()
	return v, ok
}

func (m *MutexMap) store(key, value int) {
	m.Lock()
	m.Map[key] = value
	m.Unlock()
}

func (m *MutexMap) delete(key int) {
	m.Lock()
	delete(m.Map, key)
	m.Unlock()
}

func TestMutexSyncMap(t *testing.T) {
	mutexMap := MutexMap{Map: make(map[int]int)}
	syncMap := sync.Map{}

	g := sync.WaitGroup{}
	g.Add(3)

	forSize := 100000

	go func() {
		for i := 0; i < forSize; i++ {
			mutexMap.store(i, i)
			syncMap.Store(i, i)
		}
		g.Done()
	}()

	go func() {
		for i := 0; i < forSize; i++ {
			mutexMap.load(i)
			syncMap.Load(i)
		}
		g.Done()
	}()

	go func() {
		for i := 0; i < forSize; i++ {
			mutexMap.delete(i)
			syncMap.Delete(i)
		}
		g.Done()
	}()

	g.Wait()
	fmt.Println("over")
	return
}

func BenchmarkRWMutexMap(b *testing.B) {
	sm := MutexMap{Map: make(map[int]int)}

	b.Run("load", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sm.load(i)
		}
		b.StopTimer()
	})

	b.Run("store", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sm.store(i, i)
		}
		b.StopTimer()
	})

	b.Run("delete", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sm.delete(i)
		}
		b.StopTimer()
	})
}

func BenchmarkSyncMap(b *testing.B) {
	syncMap := sync.Map{}

	b.Run("load", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			syncMap.Load(i)
		}
		b.StopTimer()
	})

	b.Run("store", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			syncMap.Store(i, i)
		}
		b.StopTimer()
	})

	b.Run("delete", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			syncMap.Delete(i)
		}
		b.StopTimer()
	})
}

// go test concurrency_map_test.go -bench=. -v -cpu=1,2,4,8

// sync.Map的特性
// 以空间换效率，通过read和dirty两个map来提高读取效率
// 优先从read map中读取(无锁)，否则再从dirty map中读取(加锁)
// 动态调整，当misses次数过多时，将dirty map提升为read map
// 延迟删除，删除只是为value打一个标记，在dirty map提升时才执行真正的删除

// oos: darwin
// goarch: amd64
// BenchmarkRWMutexMap/load-4         	100000000	        25.1 ns/op
// BenchmarkRWMutexMap/store-4        	 5000000	       322 ns/op
// BenchmarkRWMutexMap/delete-4       	10000000	       126 ns/op
// BenchmarkSyncMap/load-4            	200000000	         9.30 ns/op
// BenchmarkSyncMap/store-4           	 1000000	      1565 ns/op
// BenchmarkSyncMap/delete-4          	30000000	        38.7 ns/op

// 从bench结果看出，sync.Map的load/delete操作 都快要mutexMap，而store较慢。
