Golang SIMD的坎坷之路

由于目前go1.17 compiler目前还未支持SIMD指令，如果要想使用**现代CPU向量化计算的优势**，必须走些不太优雅的workaround

---

首先参考的是：[Optimizing Go programs by AVX2 using Auto-Vectorization in LLVM.](https://c-bata.medium.com/optimizing-go-by-avx2-using-auto-vectorization-in-llvm-118f7b366969)

参考代码在 ./_assembly_notwork

注意这篇文章是2019年的，作者为了避免cgo调用产生的overhead，使用了以下步骤：

1. 编写c版本代码，并通过 clang llvm 优化成 SIMD 的 assembly
2. 使用 c2goasm 将 c assembly to Go assembly，并在go代码中通过 go:noescape 标记，来调用go assembly

然而，当前版本 go 1.17，c2goasm似乎无法与当前的compiler正常工作，导致 *asm2plan9s: exit status 255*， 考虑到本身项目也不再维护。
因此只能使用cgo测试，承受调用C的overhead

---

使用cgo来尝试，参考的是 [I beat TiDB with 20 LOC](https://internals.tidb.io/t/topic/174)

参考代码 ./sum_float64.go

文章中的情况，也是想要让golang支持SIMD计算，注意文章中测试机器是linux amd cpu(考虑到clang编译是否会自动开启 vectorization)

而我们本机是darwin intel cpu，因此需要手动添加[Pragma loop hint directives](https://llvm.org/docs/Vectorizers.html#pragma-loop-hint-directives)

可以正常work，10x的性能提升

```markdown
BenchmarkSumFloat64_256
BenchmarkSumFloat64_256-12          	 4870705	       233.2 ns/op
BenchmarkSumFloat64_1024
BenchmarkSumFloat64_1024-12         	 1000000	      1021 ns/op
BenchmarkSumFloat64_8192
BenchmarkSumFloat64_8192-12         	  144931	      8250 ns/op
BenchmarkSumFloat64_AVX2_256
BenchmarkSumFloat64_AVX2_256-12     	13365306	        87.64 ns/op
BenchmarkSumFloat64_AVX2_1024
BenchmarkSumFloat64_AVX2_1024-12    	 8570164	       136.6 ns/op
BenchmarkSumFloat64_AVX2_8192
BenchmarkSumFloat64_AVX2_8192-12    	 1851094	       667.9 ns/op
```


---

基础知识补充

SIMD: Single Instruction Multiple Data. 
SSE and AVX are SIMD operation instructions on Intel CPU

A 128-bits register is available for SSE instructions.
A float type consumes 32 bits, so that 4 elements can be calculated at the same time.

AVX instruction has been embed 256 bits registers.
A float type consumes 32 bits, so that 8 elements can be calculated at the same time.

clang flags:
- -O3: More aggressive than -O2 with longer compile times. Recommended for codes that loops involving intensive 
  floating point calculations.
- -mavx2: Generates code with AVX2 instructions.
- [clang flags 索引](https://clang.llvm.org/docs/genindex.html)
- [gcc-compiler-flags](https://www.bu.edu/tech/support/research/software-and-programming/programming/compilers/gcc-compiler-flags)
