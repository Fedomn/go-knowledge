package simd

//#cgo CFLAGS: -O3 -mavx2
//
//void sum_float64(double buf[], int len, double *res) {
//    double acc = 0.0;
//    #pragma clang loop vectorize(enable)
//    for (int i = 0; i < len; i++) {
//        acc += buf[i];
//    }
//    *res = acc;
//}
import "C"

func SumFloat64Avx2(a []float64) float64 {
	var res float64
	C.sum_float64(
		(*C.double)(&a[0]),
		C.int(len(a)),
		(*C.double)(&res),
	)
	return res
}
