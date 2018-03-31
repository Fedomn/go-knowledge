package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

// https://www.zybuluo.com/zwh8800/note/440159
// https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185

// 一切的基础是 io包里 Reader和Writer接口

// 内存流
// 不阻塞 strings.NewReader bytes.Buffer
// 阻塞 io.Pipe
func TestMemStream(t *testing.T) {
	r := strings.NewReader("func TestMemStream(t *testing.T) {")
	p := make([]byte, 5)

	for {
		n, err := r.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("read over")
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(p[:n]))
	}
}

// io.pipe创建一个同步的内存管道 它没有内部缓存
// PipeReader 从管道中读取数据。该方法会堵塞，直到管道写入端开始写入数据 或写入端关闭了
// PipeWriter 写数据到管道中。该方法会堵塞，直到管道读取端读完所有数据 或读取端关闭了
func TestIoPipe(t *testing.T) {
	pipeReader, pipeWriter := io.Pipe()
	go PipeWrite(pipeWriter)
	go PipeRead(pipeReader)
	time.Sleep(time.Second)
}

func PipeWrite(pipeWriter *io.PipeWriter) {
	var (
		err error
		n   int
	)
	data := []byte("Go语言中文网")
	for i := 0; i < 3; i++ {
		pipeWriter.Write(data)
	}
	pipeWriter.CloseWithError(errors.New("输出3次后结束"))

	n, err = pipeWriter.Write(data)
	fmt.Println("close 后 再写入的字节数：", n, " error：", err)
}

func PipeRead(pipeReader *io.PipeReader) {
	var (
		err error
		n   int
	)
	data := make([]byte, 1024)
	for n, err = pipeReader.Read(data); err == nil; n, err = pipeReader.Read(data) {
		fmt.Printf("PipeRead %s\n", data[:n])
	}
	fmt.Println("writer 端 closewitherror 后：", err)
}

func TestFileStream(t *testing.T) {
	fr, _ := os.Open("./base_test.go")
	defer fr.Close()

	fw, _ := os.Create("./base_test.bak.go")
	defer fw.Close()

	p := make([]byte, 1000)
	for {
		rn, err := fr.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("read over")
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		// 注意 读多少写多少 p[:rn]
		wn, err := fw.Write(p[:rn])
		fmt.Printf("read %d, write %d, werr: %+v \n", rn, wn, err)
	}
}

// io.Copy内置for-loop 正确处理io.EOF 和 写入byte数
func TestIoCopy(t *testing.T) {
	fr, _ := os.Open("./base_test.go")
	fw, _ := os.Create("./base_test.bak.go")

	written, err := io.Copy(fw, fr)
	fmt.Println(written, err)
}

// bufio处理文本方便
func TestBufio(t *testing.T) {
	fr, _ := os.Open("./base_test.go")
	br := bufio.NewReader(fr)
	for i := 0; i < 5; i++ {
		line, _ := br.ReadString('\n')
		fmt.Println(line)
	}
}

func TestConcatSteam(t *testing.T) {
	file, _ := os.Open("./base_test.go")
	mr := io.MultiReader(
		strings.NewReader("Hello Multiple stream"),
		file,
	)
	io.Copy(os.Stdout, mr)
}

//  Reader 读取内容后，会自动写入到 Writer 中去
func TestDuplicateSteam(t *testing.T) {
	var buf bytes.Buffer
	r := io.TeeReader(strings.NewReader("Hello"), &buf)
	rb := make([]byte, 20)
	r.Read(rb)
	fmt.Println(buf.String())
	fmt.Println(string(rb))
}
