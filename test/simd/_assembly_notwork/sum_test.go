package _assembly_notwork

import "testing"

// important: due to c2goasm not work with current go version 1.17, so the following test is not work
// when c2goasm error: 'asm2plan9s: exit status 255'
func TestSumFloat64Avx2(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var expected float64 = 55

	actual := SumFloat64Avx2(input)
	if actual != expected {
		t.Errorf("expect %f, but got %f", expected, actual)
	}
}
