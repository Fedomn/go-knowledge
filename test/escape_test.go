package main

// golang escape analysis
// https://zhuanlan.zhihu.com/p/91559562
// https://deepu.tech/memory-management-in-golang/
// https://medium.com/@mayurwadekar2/escape-analysis-in-golang-ee40a1c064c1

// 在计算机语言编译器优化原理中，逃逸分析是指分析指针动态范围的方法，它同编译器优化原理的指针分析和外形分析相关联。
// 当变量（或者对象）在方法中分配后，其指针有可能被返回或者被全局引用，这样就会被其他过程或者线程所引用，
// 这种现象称作指针（或者引用）的逃逸(Escape)。

// golang逃逸分析在编译期完成 go run -gcflags "-m -l" main.go  (-m打印逃逸分析信息 -l禁止内联编译)
// 当大量struct都escape到堆上，就会给GC带来压力
