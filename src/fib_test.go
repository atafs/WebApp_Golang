package webapp

import "testing"

// simple code
func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

// optimized code
var result int

func benchmarkFibOptimized(i int, b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = Fib(i)
	}
	result = r
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func BenchmarkFibOptimized1(b *testing.B)  { benchmarkFibOptimized(1, b) }
func BenchmarkFibOptimized2(b *testing.B)  { benchmarkFibOptimized(2, b) }
func BenchmarkFibOptimized3(b *testing.B)  { benchmarkFibOptimized(3, b) }
func BenchmarkFibOptimized10(b *testing.B) { benchmarkFibOptimized(10, b) }
func BenchmarkFibOptimized20(b *testing.B) { benchmarkFibOptimized(20, b) }
func BenchmarkFibOptimized40(b *testing.B) { benchmarkFibOptimized(40, b) }
