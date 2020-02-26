package tdd_test

import (
	"testing"

	. "github.com/fedomn/go-knowledge/test/tdd"
)

func TestInput1(t *testing.T) {
	equals(t, "input 1 should return 1", 1, WordTrans(1))
}

func TestInput3(t *testing.T) {
	equals(t, "input 3 should return Fizz", "Fizz", WordTrans(3))
}

func TestInput5(t *testing.T) {
	equals(t, "input 5 should return Buzz", "Buzz", WordTrans(5))
}

func TestInput15(t *testing.T) {
	equals(t, "input 15 should return FizzBuzz", "FizzBuzz", WordTrans(15))
}

func TestInput7(t *testing.T) {
	equals(t, "input 7 should return Whizz", "Whizz", WordTrans(7))
}

func TestInput21(t *testing.T) {
	equals(t, "input 21 should return FizzWhizz", "FizzWhizz", WordTrans(21))
}
