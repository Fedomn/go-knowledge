#!/bin/sh
set -ex

CLANG=${CLANG:-clang}
C2GOASM=${C2GOASM:-c2goasm}

pushd ./_lib > /dev/null 2>&1

# clang flags:
# https://clang.llvm.org/docs/genindex.html
# https://www.bu.edu/tech/support/research/software-and-programming/programming/compilers/gcc-compiler-flags/

#-O3	More aggressive than -O2 with longer compile times. Recommended for codes that loops involving intensive floating point calculations.
#-mavx2	Generates code with AVX2 instructions.

$CLANG -S -O3 -mavx2 -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -c sum_float64.c

popd > /dev/null 2>&1

# c2goasm
# https://github.com/minio/c2goasm
# https://blog.minio.io/c2goasm-c-to-go-assembly-bb723d2f777f

#go get -u github.com/minio/asm2plan9s
#go get -u github.com/minio/c2goasm
#go get -u github.com/klauspost/asmfmt/cmd/asmfmt
$C2GOASM -a -f _lib/sum_float64.s sum_float64_simd.s
