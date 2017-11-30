package main

import (
	"testing"
	"fmt"
	"strings"
)

// Code From
// https://blog.learngoprogramming.com/go-functions-overview-anonymous-closures-higher-order-deferred-concurrent-6799008dde7b

type Cruncher func(int) int

func mul(n int) int {
	return n * 2
}
func add(n int) int {
	return n + 100
}
func sub(n int) int {
	return n - 1
}

func crunch(nums []int, a ...Cruncher) (rnums []int) {
	rnums = append(rnums, nums...)
	for _, f := range a {
		for i, n := range rnums {
			rnums[i] = f(n)
		}
	}
	return
}

// First-class means that funcs are value objects just like any other values
// which can be stored and passed around.
func TestFirstClassFunc(t *testing.T) {
	nums := []int{1, 2, 3}
	res := crunch(nums, sub, add, mul)
	fmt.Println(res)
}

// A noname func is an anonymous func and it’s declared inline using a function literal.
// It becomes more useful when it’s used as a closure, higher-order func, deferred func, etc.
func TestAnonymousFunc(t *testing.T) {
	mul := func(n int) int {
		return n * 2
	}
	add := func(n int) int {
		return n + 100
	}
	sub := func(n int) int {
		return n - 1
	}
	nums := []int{1, 2, 3}
	res := crunch(nums, mul, add, sub)
	fmt.Println(res)
}

type tokenizer func() (token string, ok bool)

func split(s, sep string) tokenizer {
	tokens, last := strings.Split(s, sep), 0
	return func() (string, bool) {
		if len(tokens) == last {
			return "", false
		}
		last = last + 1
		return tokens[last-1], true
	}
}

// A closure can remember all the surrounding values where it’s defined.
// One of the benefits of a closure is that it can operate on the
// captured environment as long as you want — beware the leaks!
func TestClosures(t *testing.T) {
	const sentence = "The quick brown fox jumps over the lazy dog"
	tokenizer := split(sentence, " ")
	for {
		token, ok := tokenizer()
		if !ok {
			break
		}
		fmt.Println(token)
	}
}
