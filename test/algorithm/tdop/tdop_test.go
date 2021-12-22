package tdop

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"unicode"
)

// nud: null denotation
// led: left denotation
// lbp: left binding power

// ---------------------------- prefix handler -------------------------------------------------
type prefixHandler interface {
	// nud - this is the prefix handler we talked about. In this simple parser it exists only for the literals (the numbers)
	nud() int
}

type literalToken struct {
	val int
}

func (l literalToken) nud() int {
	return l.val
}

// ---------------------------- infix handler -------------------------------------------------

type infixHandler interface {
	// led - the infix handler.
	led(left int) int
}

type infixOperator interface {
	// lbp - the left binding power of the operator. For an infix operator, it tells us how strongly the operator binds to the argument at its left.
	lbp() int
}

type endToken struct{}

func (e endToken) lbp() int {
	return 0
}

type addOperator struct{}

func (a addOperator) led(left int) int {
	return left + expression(10)
}

func (a addOperator) lbp() int {
	return 10
}

type mulOperator struct{}

func (m mulOperator) led(left int) int {
	return left * expression(20)
}

func (m mulOperator) lbp() int {
	return 20
}

var cursor = 0
var tokens []interface{}
var token interface{}

// rbp: right binding power
func expression(rbp int) int {
	t := tokens[cursor]
	cursor++
	token = tokens[cursor]
	left := t.(prefixHandler).nud()
	fmt.Printf("prefix : %+v\n", left)
	for rbp < token.(infixOperator).lbp() {
		fmt.Printf("infix : %T\n", token)
		t = token
		cursor++
		token = tokens[cursor]
		l := t.(infixHandler).led(left)
		fmt.Printf("next infix : %T got: %+v\n", t, l)
		left = l
	}
	return left
}

func buildTokens(expr string) []interface{} {
	res := make([]interface{}, 0)
	runes := []rune(expr)
	for _, r := range runes {
		if unicode.IsDigit(r) {
			num, _ := strconv.Atoi(string(r))
			res = append(res, literalToken{num})
		} else {
			switch string(r) {
			case "+":
				res = append(res, addOperator{})
			case "*":
				res = append(res, mulOperator{})
			}
		}
	}
	res = append(res, endToken{})
	return res
}

func TestParse(t *testing.T) {
	expr := "3 + 1 * 2 * 4 + 5"
	tokens = buildTokens(expr)
	require.Equal(t, 16, expression(0))
}
