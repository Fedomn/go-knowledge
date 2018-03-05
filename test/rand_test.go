package main

import (
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	mrand "math/rand"
	"testing"
	"time"
)

// http://lihaoquan.me/2016/10/15/rand-in-go.html

// math/rand	是计算机利用设计好的算法，结合提供的seed产生的随机数列，即 伪随机数
// 伪随机数是周期性的，提供相同的seed和算法，就会得出完全一样的随机数列

// crypto/rand 	是真正意义上的随机数生成方式

func TestMathRandInt(t *testing.T) {
	// 使用math/rand包里的globalRand
	mrand.Seed(time.Now().UnixNano())

	fmt.Println(mrand.Int())
	fmt.Println(mrand.Intn(10))
	fmt.Println(mrand.Float32())

	p := make([]byte, 10)
	mrand.Read(p)
	fmt.Println(base64.URLEncoding.EncodeToString(p))

	// 创建新的rand
	mr := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	fmt.Println(mr.Int31n(10))
}

func TestCryptoRandSessionId(t *testing.T) {
	b := make([]byte, 32)
	crand.Read(b)
	s := base64.URLEncoding.EncodeToString(b)
	fmt.Println(s)
}
