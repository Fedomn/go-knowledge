package main

import (
	"bufio"
	"fmt"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

// 每10000条写一次磁盘
func writeLines(fd *os.File, count *uint64) {
	num := 10000
	w := bufio.NewWriter(fd)
	for i := 0; i < num; i++ {
		unixNano := time.Now().UnixNano() / 1000000
		testJson := fmt.Sprintf(`{1231231231: fsafsfsafs %d, fasfsdflf,fasfalkoqdmf}`, unixNano)
		data := []byte(testJson)
		data = append(data, []byte("\n")...)
		w.Write(data)
		atomic.AddUint64(count, 1)
	}
	w.Flush()
	fmt.Printf("Flush %d over\n", num)
}

// 开10个协程同时跑
func run(fd *os.File) {
	var count uint64 = 0
	times := 10
	for i := 0; i < times; i++ {
		go writeLines(fd, &count)
	}
	var timeOut time.Duration = 1
	time.Sleep(time.Second * timeOut)
	countFinal := atomic.LoadUint64(&count)
	fmt.Printf("mock data over, write lines: %d, in %d seconds\n", countFinal, timeOut)
}

func TestWriteFileByHighSpeed(t *testing.T) {
	now := time.Now()
	fileName := fmt.Sprintf("file_write_test.log.%s.%02d", now.Format("20060102"), now.Hour())
	fd, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file err: %+v", err)
		os.Exit(1)
	}
	defer fd.Close()

	// 跑10次 总行数 = 10000 * 10 * 10
	totalSecond := 10
	for i := 0; i < totalSecond; i++ {
		run(fd)
	}
}
