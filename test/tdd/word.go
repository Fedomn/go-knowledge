package tdd

type Translator interface {
	Trans(input int) interface{}
}

type CommonTrans struct{}

func (CommonTrans) Trans(input int) interface{} {
	return input
}

type FizzTrans struct {
	Successor Translator
}

func (t FizzTrans) Trans(input int) interface{} {
	if input%3 == 0 {
		return "Fizz"
	}
	return t.Successor.Trans(input)
}

type BuzzTrans struct {
	Successor Translator
}

func (t BuzzTrans) Trans(input int) interface{} {
	if input%5 == 0 {
		return "Buzz"
	}
	return t.Successor.Trans(input)
}

type WhizzTrans struct {
	Successor Translator
}

func (t WhizzTrans) Trans(input int) interface{} {
	if input%7 == 0 {
		return "Whizz"
	}
	return t.Successor.Trans(input)
}

type FizzWhizzTrans struct {
	Successor Translator
}

func (t FizzWhizzTrans) Trans(input int) interface{} {
	if input%21 == 0 {
		return "FizzWhizz"
	}
	return t.Successor.Trans(input)
}

type FizzBuzzTrans struct {
	Successor Translator
}

func (t FizzBuzzTrans) Trans(input int) interface{} {
	if input%15 == 0 {
		return "FizzBuzz"
	}
	return t.Successor.Trans(input)
}

func WordTrans(input int) interface{} {
	commonTrans := CommonTrans{}
	fizzTrans := FizzTrans{commonTrans}
	buzzTrans := BuzzTrans{fizzTrans}
	whizzTrans := WhizzTrans{buzzTrans}
	fizzWhizzTrans := FizzWhizzTrans{whizzTrans}
	fizzBuzzTrans := FizzBuzzTrans{fizzWhizzTrans}

	return fizzBuzzTrans.Trans(input)
}
